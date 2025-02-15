// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/multierr"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/engine/tracer"
	"github.com/cerbos/cerbos/internal/util"
	"github.com/cerbos/cerbos/internal/validator"
)

type TestFixture struct {
	Principals map[string]*enginev1.Principal
	Resources  map[string]*enginev1.Resource
	AuxData    map[string]*enginev1.AuxData
}

const (
	principalsFileName = "principals"
	resourcesFileName  = "resources"
)

var auxDataFileNames = []string{"auxdata", "auxData", "aux_data"}

func LoadTestFixture(fsys fs.FS, path string) (tf *TestFixture, err error) {
	tf = new(TestFixture)
	tf.Principals, err = loadPrincipals(fsys, path)
	if err != nil {
		return nil, err
	}

	tf.Resources, err = loadResources(fsys, path)
	if err != nil {
		return nil, err
	}

	tf.AuxData, err = loadAuxData(fsys, path)
	if err != nil {
		return nil, err
	}

	return tf, nil
}

func loadResources(fsys fs.FS, path string) (map[string]*enginev1.Resource, error) {
	pb := &policyv1.TestFixture_Resources{}
	fp := filepath.Join(path, resourcesFileName)
	if err := loadFixtureElement(fsys, fp, pb); err != nil {
		if errors.Is(err, util.ErrNoMatchingFiles) {
			return nil, nil
		}
		return nil, err
	}

	return pb.Resources, nil
}

func loadPrincipals(fsys fs.FS, path string) (map[string]*enginev1.Principal, error) {
	pb := &policyv1.TestFixture_Principals{}
	fp := filepath.Join(path, principalsFileName)
	if err := loadFixtureElement(fsys, fp, pb); err != nil {
		if errors.Is(err, util.ErrNoMatchingFiles) {
			return nil, nil
		}
		return nil, err
	}

	return pb.Principals, nil
}

func loadAuxData(fsys fs.FS, path string) (map[string]*enginev1.AuxData, error) {
	pb := &policyv1.TestFixture_AuxData{}
	for _, fn := range auxDataFileNames {
		fp := filepath.Join(path, fn)
		if err := loadFixtureElement(fsys, fp, pb); err != nil {
			if errors.Is(err, util.ErrNoMatchingFiles) {
				continue
			}
			return nil, err
		}

		return pb.AuxData, nil
	}

	return nil, nil
}

func loadFixtureElement(fsys fs.FS, path string, pb proto.Message) error {
	file, err := util.OpenOneOfSupportedFiles(fsys, path)
	if err != nil || file == nil {
		return err
	}

	defer file.Close()
	err = util.ReadJSONOrYAML(file, pb)
	if err != nil {
		return err
	}

	return validator.Validate(pb)
}

func (tf *TestFixture) checkDupes(suite *policyv1.TestSuite) error {
	dupes := make(map[string]struct{})
	var errs error
	for _, t := range suite.Tests {
		if _, ok := dupes[t.Name]; ok {
			errs = multierr.Append(errs, fmt.Errorf("another test named %s already exists", t.Name))
		}
		dupes[t.Name] = struct{}{}
	}

	return errs
}

func (tf *TestFixture) runTestSuite(ctx context.Context, eng Checker, shouldRun func(string) bool, file string, suite *policyv1.TestSuite, trace bool) *policyv1.TestResults_Suite {
	suiteResult := &policyv1.TestResults_Suite{
		File:        file,
		Name:        suite.Name,
		Description: suite.Description,
		Summary:     &policyv1.TestResults_Summary{},
	}

	if suite.Skip {
		suiteResult.Summary.OverallResult = policyv1.TestResults_RESULT_SKIPPED
		return suiteResult
	}

	if err := tf.checkDupes(suite); err != nil {
		suiteResult.Summary.OverallResult = policyv1.TestResults_RESULT_ERRORED
		suiteResult.Error = fmt.Sprintf("Invalid test suite: %v", err)
		return suiteResult
	}

	tests, err := tf.getTests(suite)
	if err != nil {
		suiteResult.Summary.OverallResult = policyv1.TestResults_RESULT_ERRORED
		suiteResult.Error = fmt.Sprintf("Failed to load the test suite: %s", err.Error())
		return suiteResult
	}

	for _, test := range tests {
		if err := ctx.Err(); err != nil {
			return suiteResult
		}

		for _, action := range test.Input.Actions {
			testResult := runTest(ctx, eng, test, action, shouldRun, suite, trace)
			addTestResult(suiteResult, test.Name.PrincipalKey, test.Name.ResourceKey, action, test.Name.TestTableName, testResult)
		}
	}

	return suiteResult
}

func runTest(ctx context.Context, eng Checker, test *policyv1.Test, action string, shouldRun func(string) bool, suite *policyv1.TestSuite, trace bool) *policyv1.TestResults_Details {
	details := &policyv1.TestResults_Details{}

	if test.Skip || !shouldRun(fmt.Sprintf("%s/%s", suite.Name, test.Name.String())) {
		details.Result = policyv1.TestResults_RESULT_SKIPPED
		return details
	}

	inputs := []*enginev1.CheckInput{{
		RequestId: test.Input.RequestId,
		Resource:  test.Input.Resource,
		Principal: test.Input.Principal,
		Actions:   []string{action},
		AuxData:   test.Input.AuxData,
	}}

	nowFunc := time.Now
	lenientScopeSearch := false
	if test.Options != nil {
		if test.Options.Now != nil {
			ts := test.Options.Now.AsTime()
			nowFunc = func() time.Time { return ts }
		}

		lenientScopeSearch = test.Options.LenientScopeSearch
	}

	actual, traces, err := performCheck(ctx, eng, inputs, trace, nowFunc, lenientScopeSearch)
	details.EngineTrace = traces

	if err != nil {
		details.Result = policyv1.TestResults_RESULT_ERRORED
		details.Outcome = &policyv1.TestResults_Details_Error{Error: err.Error()}
		return details
	}

	if len(actual) == 0 {
		details.Result = policyv1.TestResults_RESULT_ERRORED
		details.Outcome = &policyv1.TestResults_Details_Error{Error: "Empty response from server"}
		return details
	}

	if test.Expected[action] != actual[0].Actions[action].Effect {
		details.Result = policyv1.TestResults_RESULT_FAILED
		details.Outcome = &policyv1.TestResults_Details_Failure{
			Failure: &policyv1.TestResults_Failure{
				Expected: test.Expected[action],
				Actual:   actual[0].Actions[action].Effect,
			},
		}
		return details
	}

	if expectedOutputs, ok := test.ExpectedOutputs[action]; ok {
		actualOutputs := make(map[string]*structpb.Value, len(actual[0].Outputs))
		for _, output := range actual[0].Outputs {
			actualOutputs[output.Src] = output.Val
		}

		var failures []*policyv1.TestResults_OutputFailure
		for wantKey, wantValue := range expectedOutputs.Entries {
			haveValue, ok := actualOutputs[wantKey]
			if !ok {
				failures = append(failures, &policyv1.TestResults_OutputFailure{
					Src: wantKey,
					Outcome: &policyv1.TestResults_OutputFailure_Missing{
						Missing: &policyv1.TestResults_OutputFailure_MissingValue{
							Expected: wantValue,
						},
					},
				})
				continue
			}

			if !cmp.Equal(wantValue, haveValue, protocmp.Transform()) {
				failures = append(failures, &policyv1.TestResults_OutputFailure{
					Src: wantKey,
					Outcome: &policyv1.TestResults_OutputFailure_Mismatched{
						Mismatched: &policyv1.TestResults_OutputFailure_MismatchedValue{
							Actual:   haveValue,
							Expected: wantValue,
						},
					},
				})
			}
		}

		if len(failures) > 0 {
			details.Result = policyv1.TestResults_RESULT_FAILED
			details.Outcome = &policyv1.TestResults_Details_Failure{
				Failure: &policyv1.TestResults_Failure{
					Expected: test.Expected[action],
					Actual:   actual[0].Actions[action].Effect,
					Outputs:  failures,
				},
			}
			return details
		}
	}

	details.Result = policyv1.TestResults_RESULT_PASSED
	return details
}

func performCheck(ctx context.Context, eng Checker, inputs []*enginev1.CheckInput, trace bool, nowFunc func() time.Time, lenientScopeSearch bool) ([]*enginev1.CheckOutput, []*enginev1.Trace, error) {
	checkOpts := []engine.CheckOpt{engine.WithNowFunc(nowFunc)}
	if lenientScopeSearch {
		checkOpts = append(checkOpts, engine.WithLenientScopeSearch())
	}

	if !trace {
		output, err := eng.Check(ctx, inputs, checkOpts...)
		return output, nil, err
	}

	traceCollector := tracer.NewCollector()
	checkOpts = append(checkOpts, engine.WithTraceSink(traceCollector))
	output, err := eng.Check(ctx, inputs, checkOpts...)
	return output, traceCollector.Traces(), err
}

func addTestResult(suite *policyv1.TestResults_Suite, principal, resource, action, testName string, details *policyv1.TestResults_Details) {
	addAction(addResource(addPrincipal(addTestCase(suite, testName), principal), resource), action).Details = details
	suite.Summary.TestsCount++
	incrementTally(suite.Summary, details.Result, 1)

	if details.Result > suite.Summary.OverallResult {
		suite.Summary.OverallResult = details.Result
	}
}

func addTestCase(suite *policyv1.TestResults_Suite, name string) *policyv1.TestResults_TestCase {
	for _, tc := range suite.TestCases {
		if tc.Name == name {
			return tc
		}
	}

	tc := &policyv1.TestResults_TestCase{Name: name}
	suite.TestCases = append(suite.TestCases, tc)
	return tc
}

func addPrincipal(testCaseResult *policyv1.TestResults_TestCase, name string) *policyv1.TestResults_Principal {
	for _, principal := range testCaseResult.Principals {
		if principal.Name == name {
			return principal
		}
	}

	principal := &policyv1.TestResults_Principal{Name: name}
	testCaseResult.Principals = append(testCaseResult.Principals, principal)
	return principal
}

func addResource(principal *policyv1.TestResults_Principal, name string) *policyv1.TestResults_Resource {
	for _, resource := range principal.Resources {
		if resource.Name == name {
			return resource
		}
	}

	resource := &policyv1.TestResults_Resource{Name: name}
	principal.Resources = append(principal.Resources, resource)
	return resource
}

func addAction(resource *policyv1.TestResults_Resource, name string) *policyv1.TestResults_Action {
	for _, action := range resource.Actions {
		if action.Name == name {
			return action
		}
	}

	action := &policyv1.TestResults_Action{Name: name, Details: &policyv1.TestResults_Details{}}
	resource.Actions = append(resource.Actions, action)
	return action
}

func (tf *TestFixture) getTests(suite *policyv1.TestSuite) ([]*policyv1.Test, error) {
	var allTests []*policyv1.Test

	for _, table := range suite.Tests {
		tests, err := tf.buildTests(suite, table)
		if err != nil {
			return nil, fmt.Errorf("invalid test %q: %w", table.Name, err)
		}

		allTests = append(allTests, tests...)
	}

	return allTests, nil
}

func (tf *TestFixture) buildTests(suite *policyv1.TestSuite, table *policyv1.TestTable) ([]*policyv1.Test, error) {
	matrix, err := buildTestMatrix(table)
	if err != nil {
		return nil, err
	}

	tests := make([]*policyv1.Test, len(matrix))

	for i, element := range matrix {
		tests[i], err = tf.buildTest(suite, table, element)
		if err != nil {
			return nil, err
		}
	}

	return tests, nil
}

func (tf *TestFixture) buildTest(suite *policyv1.TestSuite, table *policyv1.TestTable, matrixElement testMatrixElement) (*policyv1.Test, error) {
	name := &policyv1.Test_TestName{
		TestTableName: table.Name,
		PrincipalKey:  matrixElement.Principal,
		ResourceKey:   matrixElement.Resource,
	}

	principal, err := tf.lookupPrincipal(suite, matrixElement.Principal)
	if err != nil {
		return nil, err
	}

	resource, err := tf.lookupResource(suite, matrixElement.Resource)
	if err != nil {
		return nil, err
	}

	auxData, err := tf.lookupAuxData(suite, table.Input.AuxData)
	if err != nil {
		return nil, err
	}

	options := table.Options
	if options == nil {
		options = suite.Options
	}

	return &policyv1.Test{
		Name:        name,
		Description: table.Description,
		Skip:        table.Skip,
		SkipReason:  table.SkipReason,
		Input: &enginev1.CheckInput{
			Principal: principal,
			Resource:  resource,
			Actions:   table.Input.Actions,
			AuxData:   auxData,
		},
		Expected:        matrixElement.Expected.actions,
		ExpectedOutputs: matrixElement.Expected.outputs,
		Options:         options,
	}, nil
}

func (tf *TestFixture) lookupPrincipal(ts *policyv1.TestSuite, k string) (*enginev1.Principal, error) {
	if v, ok := ts.Principals[k]; ok {
		return v, nil
	}

	if tf != nil {
		if v, ok := tf.Principals[k]; ok {
			return v, nil
		}
	}

	return nil, fmt.Errorf("principal %q not found", k)
}

func (tf *TestFixture) lookupResource(ts *policyv1.TestSuite, k string) (*enginev1.Resource, error) {
	if v, ok := ts.Resources[k]; ok {
		return v, nil
	}

	if tf != nil {
		if v, ok := tf.Resources[k]; ok {
			return v, nil
		}
	}

	return nil, fmt.Errorf("resource %q not found", k)
}

func (tf *TestFixture) lookupAuxData(ts *policyv1.TestSuite, k string) (*enginev1.AuxData, error) {
	if k == "" {
		return nil, nil
	}

	if v, ok := ts.AuxData[k]; ok {
		return v, nil
	}

	if tf != nil {
		if v, ok := tf.AuxData[k]; ok {
			return v, nil
		}
	}

	return nil, fmt.Errorf("auxData %q not found", k)
}
