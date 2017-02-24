// Code generated by protoc-gen-go.
// source: persist/options.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	persist/options.proto

It has these top-level messages:
	QLImpl
	TypeMapping
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PersistenceOptions int32

const (
	PersistenceOptions_SQL PersistenceOptions = 0
)

var PersistenceOptions_name = map[int32]string{
	0: "SQL",
}
var PersistenceOptions_value = map[string]int32{
	"SQL": 0,
}

func (x PersistenceOptions) Enum() *PersistenceOptions {
	p := new(PersistenceOptions)
	*p = x
	return p
}
func (x PersistenceOptions) String() string {
	return proto.EnumName(PersistenceOptions_name, int32(x))
}
func (x *PersistenceOptions) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PersistenceOptions_value, data, "PersistenceOptions")
	if err != nil {
		return err
	}
	*x = PersistenceOptions(value)
	return nil
}
func (PersistenceOptions) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type QLImpl struct {
	Query            *string             `protobuf:"bytes,1,req,name=query" json:"query,omitempty"`
	Arguments        []string            `protobuf:"bytes,2,rep,name=arguments" json:"arguments,omitempty"`
	Persist          *PersistenceOptions `protobuf:"varint,3,opt,name=persist,enum=persist.PersistenceOptions" json:"persist,omitempty"`
	Mapping          *TypeMapping        `protobuf:"bytes,4,opt,name=mapping" json:"mapping,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *QLImpl) Reset()                    { *m = QLImpl{} }
func (m *QLImpl) String() string            { return proto.CompactTextString(m) }
func (*QLImpl) ProtoMessage()               {}
func (*QLImpl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QLImpl) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *QLImpl) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *QLImpl) GetPersist() PersistenceOptions {
	if m != nil && m.Persist != nil {
		return *m.Persist
	}
	return PersistenceOptions_SQL
}

func (m *QLImpl) GetMapping() *TypeMapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

type TypeMapping struct {
	Types            []*TypeMapping_TypeDescriptor `protobuf:"bytes,1,rep,name=types" json:"types,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *TypeMapping) Reset()                    { *m = TypeMapping{} }
func (m *TypeMapping) String() string            { return proto.CompactTextString(m) }
func (*TypeMapping) ProtoMessage()               {}
func (*TypeMapping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TypeMapping) GetTypes() []*TypeMapping_TypeDescriptor {
	if m != nil {
		return m.Types
	}
	return nil
}

type TypeMapping_TypeDescriptor struct {
	// if this is not setup the proto_type must be one of the built-in types
	ProtoTypeName *string `protobuf:"bytes,1,opt,name=proto_type_name,json=protoTypeName" json:"proto_type_name,omitempty"`
	// If proto_type_name is set, this need not be set.  If both this and proto_type_name
	// are set, this must be one of TYPE_ENUM, TYPE_MESSAGE
	// TYPE_GROUP is not supported
	ProtoType *google_protobuf.FieldDescriptorProto_Type `protobuf:"varint,2,opt,name=proto_type,json=protoType,enum=google.protobuf.FieldDescriptorProto_Type" json:"proto_type,omitempty"`
	// if proto_label is not setup we consider any option except LABAEL_REPEATED
	ProtoLabel       *google_protobuf.FieldDescriptorProto_Label `protobuf:"varint,3,opt,name=proto_label,json=protoLabel,enum=google.protobuf.FieldDescriptorProto_Label" json:"proto_label,omitempty"`
	GoType           *string                                     `protobuf:"bytes,4,req,name=go_type,json=goType" json:"go_type,omitempty"`
	GoPackage        *string                                     `protobuf:"bytes,5,req,name=go_package,json=goPackage" json:"go_package,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *TypeMapping_TypeDescriptor) Reset()                    { *m = TypeMapping_TypeDescriptor{} }
func (m *TypeMapping_TypeDescriptor) String() string            { return proto.CompactTextString(m) }
func (*TypeMapping_TypeDescriptor) ProtoMessage()               {}
func (*TypeMapping_TypeDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *TypeMapping_TypeDescriptor) GetProtoTypeName() string {
	if m != nil && m.ProtoTypeName != nil {
		return *m.ProtoTypeName
	}
	return ""
}

func (m *TypeMapping_TypeDescriptor) GetProtoType() google_protobuf.FieldDescriptorProto_Type {
	if m != nil && m.ProtoType != nil {
		return *m.ProtoType
	}
	return google_protobuf.FieldDescriptorProto_TYPE_DOUBLE
}

func (m *TypeMapping_TypeDescriptor) GetProtoLabel() google_protobuf.FieldDescriptorProto_Label {
	if m != nil && m.ProtoLabel != nil {
		return *m.ProtoLabel
	}
	return google_protobuf.FieldDescriptorProto_LABEL_OPTIONAL
}

func (m *TypeMapping_TypeDescriptor) GetGoType() string {
	if m != nil && m.GoType != nil {
		return *m.GoType
	}
	return ""
}

func (m *TypeMapping_TypeDescriptor) GetGoPackage() string {
	if m != nil && m.GoPackage != nil {
		return *m.GoPackage
	}
	return ""
}

var E_Ql = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*QLImpl)(nil),
	Field:         560000,
	Name:          "persist.ql",
	Tag:           "bytes,560000,opt,name=ql",
	Filename:      "persist/options.proto",
}

var E_Mapping = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*TypeMapping)(nil),
	Field:         560001,
	Name:          "persist.mapping",
	Tag:           "bytes,560001,opt,name=mapping",
	Filename:      "persist/options.proto",
}

func init() {
	proto.RegisterType((*QLImpl)(nil), "persist.QLImpl")
	proto.RegisterType((*TypeMapping)(nil), "persist.TypeMapping")
	proto.RegisterType((*TypeMapping_TypeDescriptor)(nil), "persist.TypeMapping.TypeDescriptor")
	proto.RegisterEnum("persist.PersistenceOptions", PersistenceOptions_name, PersistenceOptions_value)
	proto.RegisterExtension(E_Ql)
	proto.RegisterExtension(E_Mapping)
}

func init() { proto.RegisterFile("persist/options.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0x36, 0xe9, 0xf5, 0x4a, 0xa6, 0x78, 0x27, 0xcb, 0x89, 0xcb, 0xe9, 0x69, 0xa8, 0x20, 0xe1,
	0xe4, 0x52, 0x28, 0x28, 0x7a, 0x3e, 0x29, 0x22, 0x1c, 0xf4, 0xb4, 0xdd, 0xf3, 0xc9, 0x97, 0x92,
	0xa6, 0xe3, 0x5e, 0x70, 0x93, 0xdd, 0x26, 0x1b, 0xa1, 0x6f, 0xfa, 0xe8, 0x2f, 0xf0, 0x07, 0xf8,
	0x97, 0xfc, 0x41, 0x92, 0xdd, 0x4d, 0xaa, 0x54, 0xc4, 0xa7, 0xd9, 0xfd, 0xe6, 0x9b, 0x6f, 0x66,
	0x67, 0x66, 0xe1, 0xb6, 0xc2, 0xb2, 0xca, 0x2a, 0x3d, 0x96, 0x4a, 0x67, 0xb2, 0xa8, 0x62, 0x55,
	0x4a, 0x2d, 0xc9, 0xc0, 0xc1, 0xc7, 0x21, 0x97, 0x92, 0x0b, 0x1c, 0x1b, 0x78, 0x59, 0x7f, 0x1c,
	0xaf, 0xb0, 0x4a, 0xcb, 0x4c, 0x69, 0x59, 0x5a, 0xea, 0xe8, 0x87, 0x07, 0xfb, 0xf3, 0xe9, 0x45,
	0xae, 0x04, 0x39, 0x82, 0xfe, 0xba, 0xc6, 0x72, 0x43, 0xbd, 0xd0, 0x8f, 0x02, 0x66, 0x2f, 0xe4,
	0x1e, 0x04, 0x49, 0xc9, 0xeb, 0x1c, 0x0b, 0x5d, 0x51, 0x3f, 0xec, 0x45, 0x01, 0xdb, 0x02, 0xe4,
	0x09, 0xb4, 0xb9, 0x68, 0x2f, 0xf4, 0xa2, 0x83, 0xc9, 0xdd, 0xd8, 0xdd, 0xe3, 0x99, 0xb5, 0x58,
	0xa4, 0xf8, 0xce, 0x56, 0xc7, 0x5a, 0x2e, 0x89, 0x61, 0x90, 0x27, 0x4a, 0x65, 0x05, 0xa7, 0x7b,
	0xa1, 0x17, 0x0d, 0x27, 0x47, 0x5d, 0xd8, 0xfb, 0x8d, 0xc2, 0x4b, 0xeb, 0x63, 0x2d, 0x69, 0xf4,
	0xd3, 0x87, 0xe1, 0x6f, 0x0e, 0xf2, 0x1c, 0xfa, 0x7a, 0xa3, 0xb0, 0xa2, 0x5e, 0xd8, 0x8b, 0x86,
	0x93, 0x87, 0x7f, 0x8b, 0x36, 0xe7, 0xd7, 0xdd, 0x7b, 0x99, 0x8d, 0x38, 0xfe, 0xe6, 0xc3, 0xc1,
	0x9f, 0x1e, 0xf2, 0x08, 0x0e, 0x4d, 0x33, 0x16, 0x0d, 0x63, 0x51, 0x24, 0x39, 0x52, 0x2f, 0xf4,
	0xa2, 0x80, 0xdd, 0x34, 0x70, 0xc3, 0x7e, 0x9b, 0xe4, 0x48, 0x2e, 0x00, 0xb6, 0x3c, 0xea, 0x9b,
	0xf7, 0x9e, 0xc6, 0xb6, 0xc5, 0x71, 0xdb, 0xe2, 0xf8, 0x4d, 0x86, 0x62, 0xb5, 0x55, 0x9f, 0x35,
	0xb8, 0xa9, 0x85, 0x05, 0x9d, 0x1c, 0x99, 0xc2, 0xd0, 0x4a, 0x89, 0x64, 0x89, 0xc2, 0xf5, 0xee,
	0xf1, 0xff, 0x69, 0x4d, 0x9b, 0x10, 0x66, 0x4b, 0x31, 0x67, 0x72, 0x07, 0x06, 0xdc, 0x55, 0xb5,
	0x67, 0x66, 0xb7, 0xcf, 0x6d, 0x9a, 0x13, 0x00, 0x2e, 0x17, 0x2a, 0x49, 0x3f, 0x25, 0x1c, 0x69,
	0xdf, 0xf8, 0x02, 0x2e, 0x67, 0x16, 0x38, 0x3d, 0x01, 0xb2, 0x3b, 0x25, 0x32, 0x80, 0xde, 0xd5,
	0x7c, 0x7a, 0xeb, 0xc6, 0xf9, 0x4b, 0xf0, 0xd7, 0x82, 0xdc, 0xdf, 0xa9, 0xea, 0x12, 0xf5, 0xb5,
	0x5c, 0x39, 0x3a, 0xfd, 0xf2, 0x7d, 0x64, 0x46, 0x78, 0xd8, 0x0d, 0xc1, 0xee, 0x13, 0xf3, 0xd7,
	0xe2, 0x7c, 0xde, 0x0d, 0x9a, 0x3c, 0xd8, 0xd1, 0xb9, 0xc2, 0xf2, 0x73, 0xd6, 0xe5, 0xa5, 0x5f,
	0x9d, 0xd0, 0xbf, 0x77, 0xe1, 0xd5, 0xb3, 0x0f, 0x4f, 0x79, 0xa6, 0xaf, 0xeb, 0x65, 0x9c, 0xca,
	0x7c, 0xac, 0xd3, 0x22, 0x15, 0xb2, 0x5e, 0xd9, 0x15, 0x4f, 0xcf, 0x38, 0x16, 0x67, 0xed, 0xa7,
	0x70, 0xf6, 0x85, 0xb3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x6b, 0x12, 0xa4, 0x2e, 0x03,
	0x00, 0x00,
}
