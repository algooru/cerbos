// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0
// Source: cerbos/audit/v1/audit.proto

package auditv1

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AccessLogEntry) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_audit_v1_AccessLogEntry_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *DecisionLogEntry) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_audit_v1_DecisionLogEntry_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *MetaValues) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_audit_v1_MetaValues_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Peer) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_audit_v1_Peer_hashpb_sum(m, hasher, ignore)
	}
}
