// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1

import (
    proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto for the AccessLogRecord.Spec
func (in *AccessLogRecordSpec) DeepCopyInto(out *AccessLogRecordSpec) {
    p := proto.Clone(in).(*AccessLogRecordSpec)
    *out = *p
}
// DeepCopyInto for the AccessLogRecord.Status
func (in *AccessLogRecordStatus) DeepCopyInto(out *AccessLogRecordStatus) {
    p := proto.Clone(in).(*AccessLogRecordStatus)
    *out = *p
}
