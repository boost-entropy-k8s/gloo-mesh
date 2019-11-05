// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/external/smi/traffictarget/v1alpha1/traffic-target.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TrafficTarget associates a set of traffic definitions (rules) with a service identity which is allocated to a group of pods.
// Access is controlled via referenced TrafficSpecs and by a list of source service identities.
// * If a pod which holds the referenced service identity makes a call to the destination on one of the defined routes then access
//   will be allowed
// * Any pod which attempts to connect and is not in the defined list of sources will be denied
// * Any pod which is in the defined list, but attempts to connect on a route which is not in the list of the
//   TrafficSpecs will be denied
type TrafficTarget struct {
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Destination is the pod or group of pods to allow ingress traffic
	Destination *IdentityBindingSubject `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// Sources are the pod or group of pods to allow ingress traffic
	Sources []*IdentityBindingSubject `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	// Rules are the traffic rules to allow (HTTPRoutes | TCPRoute),
	Specs                []*TrafficTargetSpec `protobuf:"bytes,3,rep,name=specs,proto3" json:"specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TrafficTarget) Reset()         { *m = TrafficTarget{} }
func (m *TrafficTarget) String() string { return proto.CompactTextString(m) }
func (*TrafficTarget) ProtoMessage()    {}
func (*TrafficTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_008857e4c9138ea8, []int{0}
}
func (m *TrafficTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTarget.Unmarshal(m, b)
}
func (m *TrafficTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTarget.Marshal(b, m, deterministic)
}
func (m *TrafficTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTarget.Merge(m, src)
}
func (m *TrafficTarget) XXX_Size() int {
	return xxx_messageInfo_TrafficTarget.Size(m)
}
func (m *TrafficTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTarget.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTarget proto.InternalMessageInfo

func (m *TrafficTarget) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *TrafficTarget) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *TrafficTarget) GetDestination() *IdentityBindingSubject {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *TrafficTarget) GetSources() []*IdentityBindingSubject {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *TrafficTarget) GetSpecs() []*TrafficTargetSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

// TrafficTargetSpec is the TrafficSpec to allow for a TrafficTarget
type TrafficTargetSpec struct {
	// Kind is the kind of TrafficSpec to allow
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of the TrafficSpec to use
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Matches is a list of TrafficSpec routes to allow traffic for
	Matches              []string `protobuf:"bytes,3,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficTargetSpec) Reset()         { *m = TrafficTargetSpec{} }
func (m *TrafficTargetSpec) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetSpec) ProtoMessage()    {}
func (*TrafficTargetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_008857e4c9138ea8, []int{1}
}
func (m *TrafficTargetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetSpec.Unmarshal(m, b)
}
func (m *TrafficTargetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetSpec.Marshal(b, m, deterministic)
}
func (m *TrafficTargetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetSpec.Merge(m, src)
}
func (m *TrafficTargetSpec) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetSpec.Size(m)
}
func (m *TrafficTargetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetSpec proto.InternalMessageInfo

func (m *TrafficTargetSpec) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *TrafficTargetSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TrafficTargetSpec) GetMatches() []string {
	if m != nil {
		return m.Matches
	}
	return nil
}

// IdentityBindingSubject is a Kubernetes objects which should be allowed access to the TrafficTarget
type IdentityBindingSubject struct {
	// Kind is the type of Subject to allow ingress (ServiceAccount | Group)
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of the Subject, i.e. ServiceAccountName
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Namespace where the Subject is deployed
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Port defines a TCP port to apply the TrafficTarget to
	Port                 string   `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityBindingSubject) Reset()         { *m = IdentityBindingSubject{} }
func (m *IdentityBindingSubject) String() string { return proto.CompactTextString(m) }
func (*IdentityBindingSubject) ProtoMessage()    {}
func (*IdentityBindingSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_008857e4c9138ea8, []int{2}
}
func (m *IdentityBindingSubject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityBindingSubject.Unmarshal(m, b)
}
func (m *IdentityBindingSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityBindingSubject.Marshal(b, m, deterministic)
}
func (m *IdentityBindingSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityBindingSubject.Merge(m, src)
}
func (m *IdentityBindingSubject) XXX_Size() int {
	return xxx_messageInfo_IdentityBindingSubject.Size(m)
}
func (m *IdentityBindingSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityBindingSubject.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityBindingSubject proto.InternalMessageInfo

func (m *IdentityBindingSubject) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *IdentityBindingSubject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdentityBindingSubject) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *IdentityBindingSubject) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func init() {
	proto.RegisterType((*TrafficTarget)(nil), "smi.traffictarget.v1alpha1.TrafficTarget")
	proto.RegisterType((*TrafficTargetSpec)(nil), "smi.traffictarget.v1alpha1.TrafficTargetSpec")
	proto.RegisterType((*IdentityBindingSubject)(nil), "smi.traffictarget.v1alpha1.IdentityBindingSubject")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/external/smi/traffictarget/v1alpha1/traffic-target.proto", fileDescriptor_008857e4c9138ea8)
}

var fileDescriptor_008857e4c9138ea8 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0x49, 0x13, 0x5a, 0xe2, 0xaa, 0x40, 0xad, 0xaa, 0x5a, 0x22, 0x44, 0xab, 0xbd, 0xc0,
	0x81, 0xd8, 0x4a, 0xb9, 0xa0, 0x1e, 0xc3, 0x01, 0x21, 0x95, 0xcb, 0x26, 0x5c, 0x90, 0x38, 0x38,
	0xbb, 0xb3, 0x1b, 0x93, 0xac, 0x6d, 0xd9, 0x93, 0x0a, 0xae, 0x7d, 0x16, 0x0e, 0x3c, 0x0a, 0x4f,
	0xd1, 0x03, 0x6f, 0x50, 0x9e, 0x00, 0xd9, 0xce, 0x02, 0x2b, 0x0a, 0x84, 0x9e, 0x3c, 0xfe, 0xc7,
	0xdf, 0xbf, 0x33, 0x5e, 0x0f, 0x79, 0x57, 0x49, 0x9c, 0xaf, 0x66, 0x2c, 0xd7, 0x35, 0x77, 0x7a,
	0xa9, 0x87, 0x52, 0xf3, 0x1a, 0xdc, 0x7c, 0x68, 0xac, 0x7e, 0x0f, 0x39, 0x3a, 0x2e, 0x8c, 0xe4,
	0xf0, 0x01, 0xc1, 0x2a, 0xb1, 0xe4, 0xae, 0x96, 0x1c, 0xad, 0x28, 0x4b, 0x99, 0xa3, 0xb0, 0x15,
	0x20, 0x3f, 0x1f, 0x89, 0xa5, 0x99, 0x8b, 0x51, 0x23, 0x0f, 0xa3, 0xce, 0x8c, 0xd5, 0xa8, 0xe9,
	0xc0, 0xd5, 0x92, 0xb5, 0x00, 0xd6, 0x00, 0x83, 0x83, 0x4a, 0x57, 0x3a, 0x1c, 0xe3, 0x3e, 0x8a,
	0xc4, 0x60, 0x74, 0x4d, 0x41, 0x61, 0x5d, 0x48, 0x0c, 0xb5, 0x9c, 0x8f, 0x78, 0x0d, 0x28, 0x0a,
	0x81, 0x62, 0x8d, 0x3c, 0xdd, 0x00, 0xb1, 0x50, 0xfe, 0xc7, 0x07, 0x9a, 0xfd, 0x1a, 0xe1, 0x9b,
	0x20, 0x28, 0x70, 0xe5, 0x22, 0x90, 0x7e, 0xea, 0x92, 0xbd, 0x69, 0xec, 0x7a, 0x1a, 0xba, 0xa6,
	0xcf, 0xc9, 0x9d, 0xa6, 0xea, 0x04, 0x8e, 0x3b, 0x4f, 0x76, 0x4f, 0x0e, 0x59, 0xae, 0x2d, 0x30,
	0xef, 0xc3, 0xa4, 0x66, 0xaf, 0xd7, 0xd9, 0x71, 0xef, 0xcb, 0xe5, 0xd1, 0xad, 0xec, 0xc7, 0x69,
	0xfa, 0x92, 0x6c, 0x47, 0xef, 0xa4, 0x08, 0xdc, 0x41, 0x9b, 0x9b, 0x84, 0xdc, 0xf8, 0x81, 0xa7,
	0xbe, 0x5d, 0x1e, 0xed, 0x23, 0x38, 0x2c, 0x64, 0x59, 0x9e, 0xa6, 0xb2, 0x52, 0xda, 0x42, 0x9a,
	0xad, 0x71, 0x3a, 0x25, 0xbb, 0x05, 0x38, 0x94, 0x4a, 0xa0, 0xd4, 0x2a, 0xe9, 0x04, 0xb7, 0x13,
	0xf6, 0xe7, 0x3f, 0xc4, 0x5e, 0x15, 0xa0, 0x50, 0xe2, 0xc7, 0xb1, 0x54, 0x85, 0x54, 0xd5, 0x64,
	0x35, 0xf3, 0x4f, 0x22, 0xfb, 0xd5, 0x86, 0x9e, 0x91, 0x1d, 0xa7, 0x57, 0x36, 0x07, 0x97, 0x6c,
	0x1d, 0x77, 0x6f, 0xe8, 0xd8, 0x58, 0xd0, 0x17, 0xe4, 0xb6, 0x33, 0x90, 0xbb, 0xa4, 0x1b, 0xbc,
	0x86, 0x7f, 0xf3, 0x6a, 0x5d, 0xf0, 0xc4, 0x40, 0x9e, 0x45, 0xf6, 0xf4, 0xf1, 0xc5, 0x55, 0xef,
	0x1e, 0xd9, 0x6b, 0x61, 0x17, 0x57, 0xbd, 0xfb, 0xf4, 0x6e, 0x4b, 0x72, 0xe9, 0x1b, 0xb2, 0xff,
	0x9b, 0x09, 0xa5, 0xa4, 0xb7, 0x90, 0xaa, 0x08, 0xf7, 0xd3, 0xcf, 0x42, 0xec, 0x35, 0x25, 0x6a,
	0x48, 0xb6, 0xa2, 0xe6, 0x63, 0x9a, 0x90, 0x9d, 0x5a, 0x60, 0x3e, 0x87, 0x58, 0x6c, 0x3f, 0x6b,
	0xb6, 0xa9, 0x25, 0x87, 0xd7, 0xf7, 0xb9, 0xb1, 0xf7, 0x43, 0xd2, 0xf7, 0xab, 0x33, 0x22, 0x87,
	0xa4, 0x1b, 0x12, 0x3f, 0x05, 0x4f, 0x18, 0x6d, 0x31, 0xe9, 0x45, 0xc2, 0xc7, 0xe3, 0xec, 0xf3,
	0xd7, 0x47, 0x9d, 0xb7, 0x67, 0xff, 0x9c, 0x66, 0xb3, 0xa8, 0x36, 0x9d, 0xe8, 0xd9, 0x76, 0x78,
	0xcc, 0xcf, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x48, 0x14, 0x78, 0xf4, 0x24, 0x04, 0x00, 0x00,
}

func (this *TrafficTarget) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficTarget)
	if !ok {
		that2, ok := that.(TrafficTarget)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	if len(this.Sources) != len(that1.Sources) {
		return false
	}
	for i := range this.Sources {
		if !this.Sources[i].Equal(that1.Sources[i]) {
			return false
		}
	}
	if len(this.Specs) != len(that1.Specs) {
		return false
	}
	for i := range this.Specs {
		if !this.Specs[i].Equal(that1.Specs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficTargetSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficTargetSpec)
	if !ok {
		that2, ok := that.(TrafficTargetSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Matches) != len(that1.Matches) {
		return false
	}
	for i := range this.Matches {
		if this.Matches[i] != that1.Matches[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IdentityBindingSubject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentityBindingSubject)
	if !ok {
		that2, ok := that.(IdentityBindingSubject)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
