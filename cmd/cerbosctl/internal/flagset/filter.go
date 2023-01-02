// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"errors"
	"time"

	"github.com/cerbos/cerbos/client"
)

var errMoreThanOneFilter = errors.New("more than one filter specified: choose from either `tail`, `between`, `since` or `lookup`")

type AuditFilters struct {
	Lookup  string        `help:"View a specific record using the Cerbos Call ID"`
	Between timerange     `help:"View records captured between two timestamps. The timestamps must be formatted as ISO-8601"`
	Since   time.Duration `help:"View records from X hours/minutes/seconds ago to now. Unit suffixes are: h=hours, m=minutes s=seconds"`
	Tail    uint16        `help:"View the last N records"`
}

func (af *AuditFilters) Validate() error {
	filterCount := 0
	if af.Tail > 0 {
		filterCount++
	}

	if af.Between.IsSet() {
		filterCount++
	}

	if af.Since > 0 {
		filterCount++
	}

	if af.Lookup != "" {
		filterCount++
	}

	if filterCount > 1 {
		return errMoreThanOneFilter
	}

	if filterCount == 0 {
		af.Tail = 30
	}

	return nil
}

func (af *AuditFilters) GenOptions() client.AuditLogOptions {
	switch {
	case af.Tail > 0:
		return client.AuditLogOptions{
			Tail: uint32(af.Tail),
		}
	case af.Between.IsSet():
		return client.AuditLogOptions{
			StartTime: af.Between.Values[0].AsTime(),
			EndTime:   af.Between.Values[1].AsTime(),
		}
	case af.Since > 0:
		return client.AuditLogOptions{
			StartTime: time.Now().Add(time.Duration(-1) * af.Since),
			EndTime:   time.Now(),
		}
	case af.Lookup != "":
		return client.AuditLogOptions{
			Lookup: af.Lookup,
		}
	default:
		return client.AuditLogOptions{}
	}
}
