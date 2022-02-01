// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package enginev1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_engine_v1_AuxData_hashpb_sum(m *AuxData, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.AuxData.jwt"]; !ok {
		if len(m.Jwt) > 0 {
			keys := make([]string, len(m.Jwt))
			i := 0
			for k := range m.Jwt {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Jwt[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Jwt[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_CheckInput_hashpb_sum(m *CheckInput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckInput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_engine_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_engine_v1_CheckOutput_ActionEffect_hashpb_sum(m *CheckOutput_ActionEffect, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.ActionEffect.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.ActionEffect.policy"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Policy))

	}
}

func cerbos_engine_v1_CheckOutput_hashpb_sum(m *CheckOutput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.resource_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ResourceId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.actions"]; !ok {
		if len(m.Actions) > 0 {
			keys := make([]string, len(m.Actions))
			i := 0
			for k := range m.Actions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Actions[k] != nil {
					cerbos_engine_v1_CheckOutput_ActionEffect_hashpb_sum(m.Actions[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.effective_derived_roles"]; !ok {
		if len(m.EffectiveDerivedRoles) > 0 {
			for _, v := range m.EffectiveDerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.validation_errors"]; !ok {
		if len(m.ValidationErrors) > 0 {
			for _, v := range m.ValidationErrors {
				if v != nil {
					cerbos_schema_v1_ValidationError_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_Principal_hashpb_sum(m *Principal, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Principal.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_Resource_hashpb_sum(m *Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_ResourcesQueryPlanOutput_LogicalOperation_hashpb_sum(m *ResourcesQueryPlanOutput_LogicalOperation, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.LogicalOperation.operator"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Operator)))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.LogicalOperation.nodes"]; !ok {
		if len(m.Nodes) > 0 {
			for _, v := range m.Nodes {
				if v != nil {
					cerbos_engine_v1_ResourcesQueryPlanOutput_Node_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_ResourcesQueryPlanOutput_Node_hashpb_sum(m *ResourcesQueryPlanOutput_Node, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Node != nil {
		if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.Node.node"]; !ok {
			switch t := m.Node.(type) {
			case *ResourcesQueryPlanOutput_Node_LogicalOperation:
				if t.LogicalOperation != nil {
					cerbos_engine_v1_ResourcesQueryPlanOutput_LogicalOperation_hashpb_sum(t.LogicalOperation, hasher, ignore)
				}

			case *ResourcesQueryPlanOutput_Node_Expression:
				if t.Expression != nil {
					google_api_expr_v1alpha1_CheckedExpr_hashpb_sum(t.Expression, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_ResourcesQueryPlanOutput_hashpb_sum(m *ResourcesQueryPlanOutput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanOutput.filter"]; !ok {
		if m.Filter != nil {
			cerbos_engine_v1_ResourcesQueryPlanOutput_Node_hashpb_sum(m.Filter, hasher, ignore)
		}

	}
}

func cerbos_engine_v1_ResourcesQueryPlanRequest_Resource_hashpb_sum(m *ResourcesQueryPlanRequest_Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
}

func cerbos_engine_v1_ResourcesQueryPlanRequest_hashpb_sum(m *ResourcesQueryPlanRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_ResourcesQueryPlanRequest_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_engine_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.include_meta"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.IncludeMeta)))

	}
}

func cerbos_schema_v1_ValidationError_hashpb_sum(m *v1.ValidationError, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.schema.v1.ValidationError.path"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Path))

	}
	if _, ok := ignore["cerbos.schema.v1.ValidationError.message"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Message))

	}
	if _, ok := ignore["cerbos.schema.v1.ValidationError.source"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Source)))

	}
}

func google_api_expr_v1alpha1_CheckedExpr_hashpb_sum(m *v1alpha1.CheckedExpr, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.reference_map"]; !ok {
		if len(m.ReferenceMap) > 0 {
			keys := make([]int64, len(m.ReferenceMap))
			i := 0
			for k := range m.ReferenceMap {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.ReferenceMap[k] != nil {
					google_api_expr_v1alpha1_Reference_hashpb_sum(m.ReferenceMap[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.type_map"]; !ok {
		if len(m.TypeMap) > 0 {
			keys := make([]int64, len(m.TypeMap))
			i := 0
			for k := range m.TypeMap {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.TypeMap[k] != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(m.TypeMap[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.expr"]; !ok {
		if m.Expr != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Expr, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.source_info"]; !ok {
		if m.SourceInfo != nil {
			google_api_expr_v1alpha1_SourceInfo_hashpb_sum(m.SourceInfo, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Constant_hashpb_sum(m *v1alpha1.Constant, hasher hash.Hash, ignore map[string]struct{}) {
	if m.ConstantKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Constant.constant_kind"]; !ok {
			switch t := m.ConstantKind.(type) {
			case *v1alpha1.Constant_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *v1alpha1.Constant_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *v1alpha1.Constant_Int64Value:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Int64Value)))

			case *v1alpha1.Constant_Uint64Value:
				_, _ = hasher.Write(protowire.AppendVarint(nil, t.Uint64Value))

			case *v1alpha1.Constant_DoubleValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.DoubleValue)))

			case *v1alpha1.Constant_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *v1alpha1.Constant_BytesValue:
				_, _ = hasher.Write(protowire.AppendBytes(nil, t.BytesValue))

			case *v1alpha1.Constant_DurationValue:
				if t.DurationValue != nil {
					google_protobuf_Duration_hashpb_sum(t.DurationValue, hasher, ignore)
				}

			case *v1alpha1.Constant_TimestampValue:
				if t.TimestampValue != nil {
					google_protobuf_Timestamp_hashpb_sum(t.TimestampValue, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_Call_hashpb_sum(m *v1alpha1.Expr_Call, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Call.target"]; !ok {
		if m.Target != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Target, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Call.function"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Function))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Call.args"]; !ok {
		if len(m.Args) > 0 {
			for _, v := range m.Args {
				if v != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_Comprehension_hashpb_sum(m *v1alpha1.Expr_Comprehension, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.iter_var"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.IterVar))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.iter_range"]; !ok {
		if m.IterRange != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.IterRange, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.accu_var"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.AccuVar))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.accu_init"]; !ok {
		if m.AccuInit != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.AccuInit, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.loop_condition"]; !ok {
		if m.LoopCondition != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.LoopCondition, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.loop_step"]; !ok {
		if m.LoopStep != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.LoopStep, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.result"]; !ok {
		if m.Result != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Result, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Expr_CreateList_hashpb_sum(m *v1alpha1.Expr_CreateList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateList.elements"]; !ok {
		if len(m.Elements) > 0 {
			for _, v := range m.Elements {
				if v != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_CreateStruct_Entry_hashpb_sum(m *v1alpha1.Expr_CreateStruct_Entry, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.Entry.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Id)))

	}
	if m.KeyKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.Entry.key_kind"]; !ok {
			switch t := m.KeyKind.(type) {
			case *v1alpha1.Expr_CreateStruct_Entry_FieldKey:
				_, _ = hasher.Write(protowire.AppendString(nil, t.FieldKey))

			case *v1alpha1.Expr_CreateStruct_Entry_MapKey:
				if t.MapKey != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(t.MapKey, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.Entry.value"]; !ok {
		if m.Value != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Value, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Expr_CreateStruct_hashpb_sum(m *v1alpha1.Expr_CreateStruct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.message_name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.MessageName))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.entries"]; !ok {
		if len(m.Entries) > 0 {
			for _, v := range m.Entries {
				if v != nil {
					google_api_expr_v1alpha1_Expr_CreateStruct_Entry_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_Ident_hashpb_sum(m *v1alpha1.Expr_Ident, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Ident.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
}

func google_api_expr_v1alpha1_Expr_Select_hashpb_sum(m *v1alpha1.Expr_Select, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Select.operand"]; !ok {
		if m.Operand != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Operand, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Select.field"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Field))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Select.test_only"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.TestOnly)))

	}
}

func google_api_expr_v1alpha1_Expr_hashpb_sum(m *v1alpha1.Expr, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Id)))

	}
	if m.ExprKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Expr.expr_kind"]; !ok {
			switch t := m.ExprKind.(type) {
			case *v1alpha1.Expr_ConstExpr:
				if t.ConstExpr != nil {
					google_api_expr_v1alpha1_Constant_hashpb_sum(t.ConstExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_IdentExpr:
				if t.IdentExpr != nil {
					google_api_expr_v1alpha1_Expr_Ident_hashpb_sum(t.IdentExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_SelectExpr:
				if t.SelectExpr != nil {
					google_api_expr_v1alpha1_Expr_Select_hashpb_sum(t.SelectExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_CallExpr:
				if t.CallExpr != nil {
					google_api_expr_v1alpha1_Expr_Call_hashpb_sum(t.CallExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_ListExpr:
				if t.ListExpr != nil {
					google_api_expr_v1alpha1_Expr_CreateList_hashpb_sum(t.ListExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_StructExpr:
				if t.StructExpr != nil {
					google_api_expr_v1alpha1_Expr_CreateStruct_hashpb_sum(t.StructExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_ComprehensionExpr:
				if t.ComprehensionExpr != nil {
					google_api_expr_v1alpha1_Expr_Comprehension_hashpb_sum(t.ComprehensionExpr, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Reference_hashpb_sum(m *v1alpha1.Reference, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Reference.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Reference.overload_id"]; !ok {
		if len(m.OverloadId) > 0 {
			for _, v := range m.OverloadId {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.Reference.value"]; !ok {
		if m.Value != nil {
			google_api_expr_v1alpha1_Constant_hashpb_sum(m.Value, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_SourceInfo_hashpb_sum(m *v1alpha1.SourceInfo, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.syntax_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SyntaxVersion))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.location"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Location))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.line_offsets"]; !ok {
		if len(m.LineOffsets) > 0 {
			for _, v := range m.LineOffsets {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.positions"]; !ok {
		if len(m.Positions) > 0 {
			keys := make([]int64, len(m.Positions))
			i := 0
			for k := range m.Positions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Positions[k])))

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.macro_calls"]; !ok {
		if len(m.MacroCalls) > 0 {
			keys := make([]int64, len(m.MacroCalls))
			i := 0
			for k := range m.MacroCalls {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.MacroCalls[k] != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(m.MacroCalls[k], hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Type_AbstractType_hashpb_sum(m *v1alpha1.Type_AbstractType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.AbstractType.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Type.AbstractType.parameter_types"]; !ok {
		if len(m.ParameterTypes) > 0 {
			for _, v := range m.ParameterTypes {
				if v != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Type_FunctionType_hashpb_sum(m *v1alpha1.Type_FunctionType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.FunctionType.result_type"]; !ok {
		if m.ResultType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.ResultType, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Type.FunctionType.arg_types"]; !ok {
		if len(m.ArgTypes) > 0 {
			for _, v := range m.ArgTypes {
				if v != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Type_ListType_hashpb_sum(m *v1alpha1.Type_ListType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.ListType.elem_type"]; !ok {
		if m.ElemType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.ElemType, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Type_MapType_hashpb_sum(m *v1alpha1.Type_MapType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.MapType.key_type"]; !ok {
		if m.KeyType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.KeyType, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Type.MapType.value_type"]; !ok {
		if m.ValueType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.ValueType, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Type_hashpb_sum(m *v1alpha1.Type, hasher hash.Hash, ignore map[string]struct{}) {
	if m.TypeKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Type.type_kind"]; !ok {
			switch t := m.TypeKind.(type) {
			case *v1alpha1.Type_Dyn:
				if t.Dyn != nil {
					google_protobuf_Empty_hashpb_sum(t.Dyn, hasher, ignore)
				}

			case *v1alpha1.Type_Null:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Null)))

			case *v1alpha1.Type_Primitive:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Primitive)))

			case *v1alpha1.Type_Wrapper:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Wrapper)))

			case *v1alpha1.Type_WellKnown:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.WellKnown)))

			case *v1alpha1.Type_ListType_:
				if t.ListType != nil {
					google_api_expr_v1alpha1_Type_ListType_hashpb_sum(t.ListType, hasher, ignore)
				}

			case *v1alpha1.Type_MapType_:
				if t.MapType != nil {
					google_api_expr_v1alpha1_Type_MapType_hashpb_sum(t.MapType, hasher, ignore)
				}

			case *v1alpha1.Type_Function:
				if t.Function != nil {
					google_api_expr_v1alpha1_Type_FunctionType_hashpb_sum(t.Function, hasher, ignore)
				}

			case *v1alpha1.Type_MessageType:
				_, _ = hasher.Write(protowire.AppendString(nil, t.MessageType))

			case *v1alpha1.Type_TypeParam:
				_, _ = hasher.Write(protowire.AppendString(nil, t.TypeParam))

			case *v1alpha1.Type_Type:
				if t.Type != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(t.Type, hasher, ignore)
				}

			case *v1alpha1.Type_Error:
				if t.Error != nil {
					google_protobuf_Empty_hashpb_sum(t.Error, hasher, ignore)
				}

			case *v1alpha1.Type_AbstractType_:
				if t.AbstractType != nil {
					google_api_expr_v1alpha1_Type_AbstractType_hashpb_sum(t.AbstractType, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Duration_hashpb_sum(m *durationpb.Duration, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Duration.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Duration.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_Empty_hashpb_sum(m *emptypb.Empty, hasher hash.Hash, ignore map[string]struct{}) {
}

func google_protobuf_ListValue_hashpb_sum(m *structpb.ListValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.ListValue.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				if v != nil {
					google_protobuf_Value_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Struct_hashpb_sum(m *structpb.Struct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Struct.fields"]; !ok {
		if len(m.Fields) > 0 {
			keys := make([]string, len(m.Fields))
			i := 0
			for k := range m.Fields {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Fields[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Fields[k], hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_Value_hashpb_sum(m *structpb.Value, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["google.protobuf.Value.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *structpb.Value_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *structpb.Value_NumberValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.NumberValue)))

			case *structpb.Value_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *structpb.Value_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *structpb.Value_StructValue:
				if t.StructValue != nil {
					google_protobuf_Struct_hashpb_sum(t.StructValue, hasher, ignore)
				}

			case *structpb.Value_ListValue:
				if t.ListValue != nil {
					google_protobuf_ListValue_hashpb_sum(t.ListValue, hasher, ignore)
				}

			}
		}
	}
}
