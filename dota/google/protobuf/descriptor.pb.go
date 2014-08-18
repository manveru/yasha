// Code generated by protoc-gen-gogo.
// source: google/protobuf/descriptor.proto
// DO NOT EDIT!

/*
Package google_protobuf is a generated protocol buffer package.

It is generated from these files:
	google/protobuf/descriptor.proto

It has these top-level messages:
	FileDescriptorSet
	FileDescriptorProto
	DescriptorProto
	FieldDescriptorProto
	EnumDescriptorProto
	EnumValueDescriptorProto
	ServiceDescriptorProto
	MethodDescriptorProto
	FileOptions
	MessageOptions
	FieldOptions
	EnumOptions
	EnumValueOptions
	ServiceOptions
	MethodOptions
	UninterpretedOption
	SourceCodeInfo
*/
package google_protobuf

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type FieldDescriptorProto_Type int32

const (
	FieldDescriptorProto_TYPE_DOUBLE   FieldDescriptorProto_Type = 1
	FieldDescriptorProto_TYPE_FLOAT    FieldDescriptorProto_Type = 2
	FieldDescriptorProto_TYPE_INT64    FieldDescriptorProto_Type = 3
	FieldDescriptorProto_TYPE_UINT64   FieldDescriptorProto_Type = 4
	FieldDescriptorProto_TYPE_INT32    FieldDescriptorProto_Type = 5
	FieldDescriptorProto_TYPE_FIXED64  FieldDescriptorProto_Type = 6
	FieldDescriptorProto_TYPE_FIXED32  FieldDescriptorProto_Type = 7
	FieldDescriptorProto_TYPE_BOOL     FieldDescriptorProto_Type = 8
	FieldDescriptorProto_TYPE_STRING   FieldDescriptorProto_Type = 9
	FieldDescriptorProto_TYPE_GROUP    FieldDescriptorProto_Type = 10
	FieldDescriptorProto_TYPE_MESSAGE  FieldDescriptorProto_Type = 11
	FieldDescriptorProto_TYPE_BYTES    FieldDescriptorProto_Type = 12
	FieldDescriptorProto_TYPE_UINT32   FieldDescriptorProto_Type = 13
	FieldDescriptorProto_TYPE_ENUM     FieldDescriptorProto_Type = 14
	FieldDescriptorProto_TYPE_SFIXED32 FieldDescriptorProto_Type = 15
	FieldDescriptorProto_TYPE_SFIXED64 FieldDescriptorProto_Type = 16
	FieldDescriptorProto_TYPE_SINT32   FieldDescriptorProto_Type = 17
	FieldDescriptorProto_TYPE_SINT64   FieldDescriptorProto_Type = 18
)

var FieldDescriptorProto_Type_name = map[int32]string{
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var FieldDescriptorProto_Type_value = map[string]int32{
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x FieldDescriptorProto_Type) Enum() *FieldDescriptorProto_Type {
	p := new(FieldDescriptorProto_Type)
	*p = x
	return p
}
func (x FieldDescriptorProto_Type) String() string {
	return proto.EnumName(FieldDescriptorProto_Type_name, int32(x))
}
func (x *FieldDescriptorProto_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldDescriptorProto_Type_value, data, "FieldDescriptorProto_Type")
	if err != nil {
		return err
	}
	*x = FieldDescriptorProto_Type(value)
	return nil
}

type FieldDescriptorProto_Label int32

const (
	FieldDescriptorProto_LABEL_OPTIONAL FieldDescriptorProto_Label = 1
	FieldDescriptorProto_LABEL_REQUIRED FieldDescriptorProto_Label = 2
	FieldDescriptorProto_LABEL_REPEATED FieldDescriptorProto_Label = 3
)

var FieldDescriptorProto_Label_name = map[int32]string{
	1: "LABEL_OPTIONAL",
	2: "LABEL_REQUIRED",
	3: "LABEL_REPEATED",
}
var FieldDescriptorProto_Label_value = map[string]int32{
	"LABEL_OPTIONAL": 1,
	"LABEL_REQUIRED": 2,
	"LABEL_REPEATED": 3,
}

func (x FieldDescriptorProto_Label) Enum() *FieldDescriptorProto_Label {
	p := new(FieldDescriptorProto_Label)
	*p = x
	return p
}
func (x FieldDescriptorProto_Label) String() string {
	return proto.EnumName(FieldDescriptorProto_Label_name, int32(x))
}
func (x *FieldDescriptorProto_Label) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldDescriptorProto_Label_value, data, "FieldDescriptorProto_Label")
	if err != nil {
		return err
	}
	*x = FieldDescriptorProto_Label(value)
	return nil
}

type FileOptions_OptimizeMode int32

const (
	FileOptions_SPEED        FileOptions_OptimizeMode = 1
	FileOptions_CODE_SIZE    FileOptions_OptimizeMode = 2
	FileOptions_LITE_RUNTIME FileOptions_OptimizeMode = 3
)

var FileOptions_OptimizeMode_name = map[int32]string{
	1: "SPEED",
	2: "CODE_SIZE",
	3: "LITE_RUNTIME",
}
var FileOptions_OptimizeMode_value = map[string]int32{
	"SPEED":        1,
	"CODE_SIZE":    2,
	"LITE_RUNTIME": 3,
}

func (x FileOptions_OptimizeMode) Enum() *FileOptions_OptimizeMode {
	p := new(FileOptions_OptimizeMode)
	*p = x
	return p
}
func (x FileOptions_OptimizeMode) String() string {
	return proto.EnumName(FileOptions_OptimizeMode_name, int32(x))
}
func (x *FileOptions_OptimizeMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FileOptions_OptimizeMode_value, data, "FileOptions_OptimizeMode")
	if err != nil {
		return err
	}
	*x = FileOptions_OptimizeMode(value)
	return nil
}

type FieldOptions_CType int32

const (
	FieldOptions_STRING       FieldOptions_CType = 0
	FieldOptions_CORD         FieldOptions_CType = 1
	FieldOptions_STRING_PIECE FieldOptions_CType = 2
)

var FieldOptions_CType_name = map[int32]string{
	0: "STRING",
	1: "CORD",
	2: "STRING_PIECE",
}
var FieldOptions_CType_value = map[string]int32{
	"STRING":       0,
	"CORD":         1,
	"STRING_PIECE": 2,
}

func (x FieldOptions_CType) Enum() *FieldOptions_CType {
	p := new(FieldOptions_CType)
	*p = x
	return p
}
func (x FieldOptions_CType) String() string {
	return proto.EnumName(FieldOptions_CType_name, int32(x))
}
func (x *FieldOptions_CType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldOptions_CType_value, data, "FieldOptions_CType")
	if err != nil {
		return err
	}
	*x = FieldOptions_CType(value)
	return nil
}

type FileDescriptorSet struct {
	File             []*FileDescriptorProto `protobuf:"bytes,1,rep,name=file" json:"file,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *FileDescriptorSet) Reset()         { *m = FileDescriptorSet{} }
func (m *FileDescriptorSet) String() string { return proto.CompactTextString(m) }
func (*FileDescriptorSet) ProtoMessage()    {}

func (m *FileDescriptorSet) GetFile() []*FileDescriptorProto {
	if m != nil {
		return m.File
	}
	return nil
}

type FileDescriptorProto struct {
	Name             *string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Package          *string                   `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	Dependency       []string                  `protobuf:"bytes,3,rep,name=dependency" json:"dependency,omitempty"`
	PublicDependency []int32                   `protobuf:"varint,10,rep,name=public_dependency" json:"public_dependency,omitempty"`
	WeakDependency   []int32                   `protobuf:"varint,11,rep,name=weak_dependency" json:"weak_dependency,omitempty"`
	MessageType      []*DescriptorProto        `protobuf:"bytes,4,rep,name=message_type" json:"message_type,omitempty"`
	EnumType         []*EnumDescriptorProto    `protobuf:"bytes,5,rep,name=enum_type" json:"enum_type,omitempty"`
	Service          []*ServiceDescriptorProto `protobuf:"bytes,6,rep,name=service" json:"service,omitempty"`
	Extension        []*FieldDescriptorProto   `protobuf:"bytes,7,rep,name=extension" json:"extension,omitempty"`
	Options          *FileOptions              `protobuf:"bytes,8,opt,name=options" json:"options,omitempty"`
	SourceCodeInfo   *SourceCodeInfo           `protobuf:"bytes,9,opt,name=source_code_info" json:"source_code_info,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *FileDescriptorProto) Reset()         { *m = FileDescriptorProto{} }
func (m *FileDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*FileDescriptorProto) ProtoMessage()    {}

func (m *FileDescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FileDescriptorProto) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *FileDescriptorProto) GetDependency() []string {
	if m != nil {
		return m.Dependency
	}
	return nil
}

func (m *FileDescriptorProto) GetPublicDependency() []int32 {
	if m != nil {
		return m.PublicDependency
	}
	return nil
}

func (m *FileDescriptorProto) GetWeakDependency() []int32 {
	if m != nil {
		return m.WeakDependency
	}
	return nil
}

func (m *FileDescriptorProto) GetMessageType() []*DescriptorProto {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *FileDescriptorProto) GetEnumType() []*EnumDescriptorProto {
	if m != nil {
		return m.EnumType
	}
	return nil
}

func (m *FileDescriptorProto) GetService() []*ServiceDescriptorProto {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *FileDescriptorProto) GetExtension() []*FieldDescriptorProto {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *FileDescriptorProto) GetOptions() *FileOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *FileDescriptorProto) GetSourceCodeInfo() *SourceCodeInfo {
	if m != nil {
		return m.SourceCodeInfo
	}
	return nil
}

type DescriptorProto struct {
	Name             *string                           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Field            []*FieldDescriptorProto           `protobuf:"bytes,2,rep,name=field" json:"field,omitempty"`
	Extension        []*FieldDescriptorProto           `protobuf:"bytes,6,rep,name=extension" json:"extension,omitempty"`
	NestedType       []*DescriptorProto                `protobuf:"bytes,3,rep,name=nested_type" json:"nested_type,omitempty"`
	EnumType         []*EnumDescriptorProto            `protobuf:"bytes,4,rep,name=enum_type" json:"enum_type,omitempty"`
	ExtensionRange   []*DescriptorProto_ExtensionRange `protobuf:"bytes,5,rep,name=extension_range" json:"extension_range,omitempty"`
	Options          *MessageOptions                   `protobuf:"bytes,7,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *DescriptorProto) Reset()         { *m = DescriptorProto{} }
func (m *DescriptorProto) String() string { return proto.CompactTextString(m) }
func (*DescriptorProto) ProtoMessage()    {}

func (m *DescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DescriptorProto) GetField() []*FieldDescriptorProto {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *DescriptorProto) GetExtension() []*FieldDescriptorProto {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *DescriptorProto) GetNestedType() []*DescriptorProto {
	if m != nil {
		return m.NestedType
	}
	return nil
}

func (m *DescriptorProto) GetEnumType() []*EnumDescriptorProto {
	if m != nil {
		return m.EnumType
	}
	return nil
}

func (m *DescriptorProto) GetExtensionRange() []*DescriptorProto_ExtensionRange {
	if m != nil {
		return m.ExtensionRange
	}
	return nil
}

func (m *DescriptorProto) GetOptions() *MessageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type DescriptorProto_ExtensionRange struct {
	Start            *int32 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End              *int32 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DescriptorProto_ExtensionRange) Reset()         { *m = DescriptorProto_ExtensionRange{} }
func (m *DescriptorProto_ExtensionRange) String() string { return proto.CompactTextString(m) }
func (*DescriptorProto_ExtensionRange) ProtoMessage()    {}

func (m *DescriptorProto_ExtensionRange) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *DescriptorProto_ExtensionRange) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

type FieldDescriptorProto struct {
	Name             *string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Number           *int32                      `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	Label            *FieldDescriptorProto_Label `protobuf:"varint,4,opt,name=label,enum=google.protobuf.FieldDescriptorProto_Label,def=1" json:"label,omitempty"`
	Type             *FieldDescriptorProto_Type  `protobuf:"varint,5,opt,name=type,enum=google.protobuf.FieldDescriptorProto_Type,def=1" json:"type,omitempty"`
	TypeName         *string                     `protobuf:"bytes,6,opt,name=type_name" json:"type_name,omitempty"`
	Extendee         *string                     `protobuf:"bytes,2,opt,name=extendee" json:"extendee,omitempty"`
	DefaultValue     *string                     `protobuf:"bytes,7,opt,name=default_value" json:"default_value,omitempty"`
	Options          *FieldOptions               `protobuf:"bytes,8,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *FieldDescriptorProto) Reset()         { *m = FieldDescriptorProto{} }
func (m *FieldDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*FieldDescriptorProto) ProtoMessage()    {}

const Default_FieldDescriptorProto_Label FieldDescriptorProto_Label = FieldDescriptorProto_LABEL_OPTIONAL
const Default_FieldDescriptorProto_Type FieldDescriptorProto_Type = FieldDescriptorProto_TYPE_DOUBLE

func (m *FieldDescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FieldDescriptorProto) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *FieldDescriptorProto) GetLabel() FieldDescriptorProto_Label {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return Default_FieldDescriptorProto_Label
}

func (m *FieldDescriptorProto) GetType() FieldDescriptorProto_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_FieldDescriptorProto_Type
}

func (m *FieldDescriptorProto) GetTypeName() string {
	if m != nil && m.TypeName != nil {
		return *m.TypeName
	}
	return ""
}

func (m *FieldDescriptorProto) GetExtendee() string {
	if m != nil && m.Extendee != nil {
		return *m.Extendee
	}
	return ""
}

func (m *FieldDescriptorProto) GetDefaultValue() string {
	if m != nil && m.DefaultValue != nil {
		return *m.DefaultValue
	}
	return ""
}

func (m *FieldDescriptorProto) GetOptions() *FieldOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type EnumDescriptorProto struct {
	Name             *string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            []*EnumValueDescriptorProto `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
	Options          *EnumOptions                `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *EnumDescriptorProto) Reset()         { *m = EnumDescriptorProto{} }
func (m *EnumDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*EnumDescriptorProto) ProtoMessage()    {}

func (m *EnumDescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EnumDescriptorProto) GetValue() []*EnumValueDescriptorProto {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EnumDescriptorProto) GetOptions() *EnumOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type EnumValueDescriptorProto struct {
	Name             *string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Number           *int32            `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	Options          *EnumValueOptions `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *EnumValueDescriptorProto) Reset()         { *m = EnumValueDescriptorProto{} }
func (m *EnumValueDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*EnumValueDescriptorProto) ProtoMessage()    {}

func (m *EnumValueDescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EnumValueDescriptorProto) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *EnumValueDescriptorProto) GetOptions() *EnumValueOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ServiceDescriptorProto struct {
	Name             *string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Method           []*MethodDescriptorProto `protobuf:"bytes,2,rep,name=method" json:"method,omitempty"`
	Options          *ServiceOptions          `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ServiceDescriptorProto) Reset()         { *m = ServiceDescriptorProto{} }
func (m *ServiceDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*ServiceDescriptorProto) ProtoMessage()    {}

func (m *ServiceDescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ServiceDescriptorProto) GetMethod() []*MethodDescriptorProto {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *ServiceDescriptorProto) GetOptions() *ServiceOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type MethodDescriptorProto struct {
	Name             *string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	InputType        *string        `protobuf:"bytes,2,opt,name=input_type" json:"input_type,omitempty"`
	OutputType       *string        `protobuf:"bytes,3,opt,name=output_type" json:"output_type,omitempty"`
	Options          *MethodOptions `protobuf:"bytes,4,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MethodDescriptorProto) Reset()         { *m = MethodDescriptorProto{} }
func (m *MethodDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*MethodDescriptorProto) ProtoMessage()    {}

func (m *MethodDescriptorProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MethodDescriptorProto) GetInputType() string {
	if m != nil && m.InputType != nil {
		return *m.InputType
	}
	return ""
}

func (m *MethodDescriptorProto) GetOutputType() string {
	if m != nil && m.OutputType != nil {
		return *m.OutputType
	}
	return ""
}

func (m *MethodDescriptorProto) GetOptions() *MethodOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type FileOptions struct {
	JavaPackage               *string                   `protobuf:"bytes,1,opt,name=java_package" json:"java_package,omitempty"`
	JavaOuterClassname        *string                   `protobuf:"bytes,8,opt,name=java_outer_classname" json:"java_outer_classname,omitempty"`
	JavaMultipleFiles         *bool                     `protobuf:"varint,10,opt,name=java_multiple_files,def=0" json:"java_multiple_files,omitempty"`
	JavaGenerateEqualsAndHash *bool                     `protobuf:"varint,20,opt,name=java_generate_equals_and_hash,def=0" json:"java_generate_equals_and_hash,omitempty"`
	OptimizeFor               *FileOptions_OptimizeMode `protobuf:"varint,9,opt,name=optimize_for,enum=google.protobuf.FileOptions_OptimizeMode,def=1" json:"optimize_for,omitempty"`
	GoPackage                 *string                   `protobuf:"bytes,11,opt,name=go_package" json:"go_package,omitempty"`
	CcGenericServices         *bool                     `protobuf:"varint,16,opt,name=cc_generic_services,def=0" json:"cc_generic_services,omitempty"`
	JavaGenericServices       *bool                     `protobuf:"varint,17,opt,name=java_generic_services,def=0" json:"java_generic_services,omitempty"`
	PyGenericServices         *bool                     `protobuf:"varint,18,opt,name=py_generic_services,def=0" json:"py_generic_services,omitempty"`
	UninterpretedOption       []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions            map[int32]proto.Extension `json:"-"`
	XXX_unrecognized          []byte                    `json:"-"`
}

func (m *FileOptions) Reset()         { *m = FileOptions{} }
func (m *FileOptions) String() string { return proto.CompactTextString(m) }
func (*FileOptions) ProtoMessage()    {}

var extRange_FileOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*FileOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_FileOptions
}
func (m *FileOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_FileOptions_JavaMultipleFiles bool = false
const Default_FileOptions_JavaGenerateEqualsAndHash bool = false
const Default_FileOptions_OptimizeFor FileOptions_OptimizeMode = FileOptions_SPEED
const Default_FileOptions_CcGenericServices bool = false
const Default_FileOptions_JavaGenericServices bool = false
const Default_FileOptions_PyGenericServices bool = false

func (m *FileOptions) GetJavaPackage() string {
	if m != nil && m.JavaPackage != nil {
		return *m.JavaPackage
	}
	return ""
}

func (m *FileOptions) GetJavaOuterClassname() string {
	if m != nil && m.JavaOuterClassname != nil {
		return *m.JavaOuterClassname
	}
	return ""
}

func (m *FileOptions) GetJavaMultipleFiles() bool {
	if m != nil && m.JavaMultipleFiles != nil {
		return *m.JavaMultipleFiles
	}
	return Default_FileOptions_JavaMultipleFiles
}

func (m *FileOptions) GetJavaGenerateEqualsAndHash() bool {
	if m != nil && m.JavaGenerateEqualsAndHash != nil {
		return *m.JavaGenerateEqualsAndHash
	}
	return Default_FileOptions_JavaGenerateEqualsAndHash
}

func (m *FileOptions) GetOptimizeFor() FileOptions_OptimizeMode {
	if m != nil && m.OptimizeFor != nil {
		return *m.OptimizeFor
	}
	return Default_FileOptions_OptimizeFor
}

func (m *FileOptions) GetGoPackage() string {
	if m != nil && m.GoPackage != nil {
		return *m.GoPackage
	}
	return ""
}

func (m *FileOptions) GetCcGenericServices() bool {
	if m != nil && m.CcGenericServices != nil {
		return *m.CcGenericServices
	}
	return Default_FileOptions_CcGenericServices
}

func (m *FileOptions) GetJavaGenericServices() bool {
	if m != nil && m.JavaGenericServices != nil {
		return *m.JavaGenericServices
	}
	return Default_FileOptions_JavaGenericServices
}

func (m *FileOptions) GetPyGenericServices() bool {
	if m != nil && m.PyGenericServices != nil {
		return *m.PyGenericServices
	}
	return Default_FileOptions_PyGenericServices
}

func (m *FileOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type MessageOptions struct {
	MessageSetWireFormat         *bool                     `protobuf:"varint,1,opt,name=message_set_wire_format,def=0" json:"message_set_wire_format,omitempty"`
	NoStandardDescriptorAccessor *bool                     `protobuf:"varint,2,opt,name=no_standard_descriptor_accessor,def=0" json:"no_standard_descriptor_accessor,omitempty"`
	UninterpretedOption          []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions               map[int32]proto.Extension `json:"-"`
	XXX_unrecognized             []byte                    `json:"-"`
}

func (m *MessageOptions) Reset()         { *m = MessageOptions{} }
func (m *MessageOptions) String() string { return proto.CompactTextString(m) }
func (*MessageOptions) ProtoMessage()    {}

var extRange_MessageOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*MessageOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MessageOptions
}
func (m *MessageOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_MessageOptions_MessageSetWireFormat bool = false
const Default_MessageOptions_NoStandardDescriptorAccessor bool = false

func (m *MessageOptions) GetMessageSetWireFormat() bool {
	if m != nil && m.MessageSetWireFormat != nil {
		return *m.MessageSetWireFormat
	}
	return Default_MessageOptions_MessageSetWireFormat
}

func (m *MessageOptions) GetNoStandardDescriptorAccessor() bool {
	if m != nil && m.NoStandardDescriptorAccessor != nil {
		return *m.NoStandardDescriptorAccessor
	}
	return Default_MessageOptions_NoStandardDescriptorAccessor
}

func (m *MessageOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type FieldOptions struct {
	Ctype               *FieldOptions_CType       `protobuf:"varint,1,opt,name=ctype,enum=google.protobuf.FieldOptions_CType,def=0" json:"ctype,omitempty"`
	Packed              *bool                     `protobuf:"varint,2,opt,name=packed" json:"packed,omitempty"`
	Lazy                *bool                     `protobuf:"varint,5,opt,name=lazy,def=0" json:"lazy,omitempty"`
	Deprecated          *bool                     `protobuf:"varint,3,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
	ExperimentalMapKey  *string                   `protobuf:"bytes,9,opt,name=experimental_map_key" json:"experimental_map_key,omitempty"`
	Weak                *bool                     `protobuf:"varint,10,opt,name=weak,def=0" json:"weak,omitempty"`
	UninterpretedOption []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions      map[int32]proto.Extension `json:"-"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *FieldOptions) Reset()         { *m = FieldOptions{} }
func (m *FieldOptions) String() string { return proto.CompactTextString(m) }
func (*FieldOptions) ProtoMessage()    {}

var extRange_FieldOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*FieldOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_FieldOptions
}
func (m *FieldOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_FieldOptions_Ctype FieldOptions_CType = FieldOptions_STRING
const Default_FieldOptions_Lazy bool = false
const Default_FieldOptions_Deprecated bool = false
const Default_FieldOptions_Weak bool = false

func (m *FieldOptions) GetCtype() FieldOptions_CType {
	if m != nil && m.Ctype != nil {
		return *m.Ctype
	}
	return Default_FieldOptions_Ctype
}

func (m *FieldOptions) GetPacked() bool {
	if m != nil && m.Packed != nil {
		return *m.Packed
	}
	return false
}

func (m *FieldOptions) GetLazy() bool {
	if m != nil && m.Lazy != nil {
		return *m.Lazy
	}
	return Default_FieldOptions_Lazy
}

func (m *FieldOptions) GetDeprecated() bool {
	if m != nil && m.Deprecated != nil {
		return *m.Deprecated
	}
	return Default_FieldOptions_Deprecated
}

func (m *FieldOptions) GetExperimentalMapKey() string {
	if m != nil && m.ExperimentalMapKey != nil {
		return *m.ExperimentalMapKey
	}
	return ""
}

func (m *FieldOptions) GetWeak() bool {
	if m != nil && m.Weak != nil {
		return *m.Weak
	}
	return Default_FieldOptions_Weak
}

func (m *FieldOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type EnumOptions struct {
	AllowAlias          *bool                     `protobuf:"varint,2,opt,name=allow_alias,def=1" json:"allow_alias,omitempty"`
	UninterpretedOption []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions      map[int32]proto.Extension `json:"-"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *EnumOptions) Reset()         { *m = EnumOptions{} }
func (m *EnumOptions) String() string { return proto.CompactTextString(m) }
func (*EnumOptions) ProtoMessage()    {}

var extRange_EnumOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*EnumOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnumOptions
}
func (m *EnumOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_EnumOptions_AllowAlias bool = true

func (m *EnumOptions) GetAllowAlias() bool {
	if m != nil && m.AllowAlias != nil {
		return *m.AllowAlias
	}
	return Default_EnumOptions_AllowAlias
}

func (m *EnumOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type EnumValueOptions struct {
	UninterpretedOption []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions      map[int32]proto.Extension `json:"-"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *EnumValueOptions) Reset()         { *m = EnumValueOptions{} }
func (m *EnumValueOptions) String() string { return proto.CompactTextString(m) }
func (*EnumValueOptions) ProtoMessage()    {}

var extRange_EnumValueOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*EnumValueOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnumValueOptions
}
func (m *EnumValueOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *EnumValueOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type ServiceOptions struct {
	UninterpretedOption []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions      map[int32]proto.Extension `json:"-"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *ServiceOptions) Reset()         { *m = ServiceOptions{} }
func (m *ServiceOptions) String() string { return proto.CompactTextString(m) }
func (*ServiceOptions) ProtoMessage()    {}

var extRange_ServiceOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*ServiceOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ServiceOptions
}
func (m *ServiceOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *ServiceOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type MethodOptions struct {
	UninterpretedOption []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option" json:"uninterpreted_option,omitempty"`
	XXX_extensions      map[int32]proto.Extension `json:"-"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *MethodOptions) Reset()         { *m = MethodOptions{} }
func (m *MethodOptions) String() string { return proto.CompactTextString(m) }
func (*MethodOptions) ProtoMessage()    {}

var extRange_MethodOptions = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*MethodOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MethodOptions
}
func (m *MethodOptions) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *MethodOptions) GetUninterpretedOption() []*UninterpretedOption {
	if m != nil {
		return m.UninterpretedOption
	}
	return nil
}

type UninterpretedOption struct {
	Name             []*UninterpretedOption_NamePart `protobuf:"bytes,2,rep,name=name" json:"name,omitempty"`
	IdentifierValue  *string                         `protobuf:"bytes,3,opt,name=identifier_value" json:"identifier_value,omitempty"`
	PositiveIntValue *uint64                         `protobuf:"varint,4,opt,name=positive_int_value" json:"positive_int_value,omitempty"`
	NegativeIntValue *int64                          `protobuf:"varint,5,opt,name=negative_int_value" json:"negative_int_value,omitempty"`
	DoubleValue      *float64                        `protobuf:"fixed64,6,opt,name=double_value" json:"double_value,omitempty"`
	StringValue      []byte                          `protobuf:"bytes,7,opt,name=string_value" json:"string_value,omitempty"`
	AggregateValue   *string                         `protobuf:"bytes,8,opt,name=aggregate_value" json:"aggregate_value,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *UninterpretedOption) Reset()         { *m = UninterpretedOption{} }
func (m *UninterpretedOption) String() string { return proto.CompactTextString(m) }
func (*UninterpretedOption) ProtoMessage()    {}

func (m *UninterpretedOption) GetName() []*UninterpretedOption_NamePart {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UninterpretedOption) GetIdentifierValue() string {
	if m != nil && m.IdentifierValue != nil {
		return *m.IdentifierValue
	}
	return ""
}

func (m *UninterpretedOption) GetPositiveIntValue() uint64 {
	if m != nil && m.PositiveIntValue != nil {
		return *m.PositiveIntValue
	}
	return 0
}

func (m *UninterpretedOption) GetNegativeIntValue() int64 {
	if m != nil && m.NegativeIntValue != nil {
		return *m.NegativeIntValue
	}
	return 0
}

func (m *UninterpretedOption) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *UninterpretedOption) GetStringValue() []byte {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *UninterpretedOption) GetAggregateValue() string {
	if m != nil && m.AggregateValue != nil {
		return *m.AggregateValue
	}
	return ""
}

type UninterpretedOption_NamePart struct {
	NamePart         *string `protobuf:"bytes,1,req,name=name_part" json:"name_part,omitempty"`
	IsExtension      *bool   `protobuf:"varint,2,req,name=is_extension" json:"is_extension,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UninterpretedOption_NamePart) Reset()         { *m = UninterpretedOption_NamePart{} }
func (m *UninterpretedOption_NamePart) String() string { return proto.CompactTextString(m) }
func (*UninterpretedOption_NamePart) ProtoMessage()    {}

func (m *UninterpretedOption_NamePart) GetNamePart() string {
	if m != nil && m.NamePart != nil {
		return *m.NamePart
	}
	return ""
}

func (m *UninterpretedOption_NamePart) GetIsExtension() bool {
	if m != nil && m.IsExtension != nil {
		return *m.IsExtension
	}
	return false
}

type SourceCodeInfo struct {
	Location         []*SourceCodeInfo_Location `protobuf:"bytes,1,rep,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *SourceCodeInfo) Reset()         { *m = SourceCodeInfo{} }
func (m *SourceCodeInfo) String() string { return proto.CompactTextString(m) }
func (*SourceCodeInfo) ProtoMessage()    {}

func (m *SourceCodeInfo) GetLocation() []*SourceCodeInfo_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type SourceCodeInfo_Location struct {
	Path             []int32 `protobuf:"varint,1,rep,packed,name=path" json:"path,omitempty"`
	Span             []int32 `protobuf:"varint,2,rep,packed,name=span" json:"span,omitempty"`
	LeadingComments  *string `protobuf:"bytes,3,opt,name=leading_comments" json:"leading_comments,omitempty"`
	TrailingComments *string `protobuf:"bytes,4,opt,name=trailing_comments" json:"trailing_comments,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SourceCodeInfo_Location) Reset()         { *m = SourceCodeInfo_Location{} }
func (m *SourceCodeInfo_Location) String() string { return proto.CompactTextString(m) }
func (*SourceCodeInfo_Location) ProtoMessage()    {}

func (m *SourceCodeInfo_Location) GetPath() []int32 {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *SourceCodeInfo_Location) GetSpan() []int32 {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *SourceCodeInfo_Location) GetLeadingComments() string {
	if m != nil && m.LeadingComments != nil {
		return *m.LeadingComments
	}
	return ""
}

func (m *SourceCodeInfo_Location) GetTrailingComments() string {
	if m != nil && m.TrailingComments != nil {
		return *m.TrailingComments
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.protobuf.FieldDescriptorProto_Type", FieldDescriptorProto_Type_name, FieldDescriptorProto_Type_value)
	proto.RegisterEnum("google.protobuf.FieldDescriptorProto_Label", FieldDescriptorProto_Label_name, FieldDescriptorProto_Label_value)
	proto.RegisterEnum("google.protobuf.FileOptions_OptimizeMode", FileOptions_OptimizeMode_name, FileOptions_OptimizeMode_value)
	proto.RegisterEnum("google.protobuf.FieldOptions_CType", FieldOptions_CType_name, FieldOptions_CType_value)
}
