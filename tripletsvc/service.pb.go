// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: service.proto

package tripletsvc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComputeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []byte `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ComputeRequest) Reset() {
	*x = ComputeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRequest) ProtoMessage() {}

func (x *ComputeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRequest.ProtoReflect.Descriptor instead.
func (*ComputeRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRequest) GetText() []byte {
	if x != nil {
		return x.Text
	}
	return nil
}

type Triplet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words []string `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *Triplet) Reset() {
	*x = Triplet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Triplet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Triplet) ProtoMessage() {}

func (x *Triplet) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Triplet.ProtoReflect.Descriptor instead.
func (*Triplet) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Triplet) GetWords() []string {
	if x != nil {
		return x.Words
	}
	return nil
}

type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Triplet   *Triplet `protobuf:"bytes,1,opt,name=triplet,proto3" json:"triplet,omitempty"`
	Occurence int32    `protobuf:"varint,2,opt,name=occurence,proto3" json:"occurence,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Stat) GetTriplet() *Triplet {
	if x != nil {
		return x.Triplet
	}
	return nil
}

func (x *Stat) GetOccurence() int32 {
	if x != nil {
		return x.Occurence
	}
	return 0
}

type ComputeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *ComputeResponse) Reset() {
	*x = ComputeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResponse) ProtoMessage() {}

func (x *ComputeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResponse.ProtoReflect.Descriptor instead.
func (*ComputeResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ComputeResponse) GetStats() []*Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x48, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x22,
	0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x52, 0x07, 0x74, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x2e, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x32, 0x40, 0x0a, 0x0e, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6b, 0x68, 0x65, 0x6e, 0x61, 0x6b, 0x68, 0x2f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x74, 0x65, 0x73, 0x74, 0x6e, 0x72, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x73, 0x76,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_proto_goTypes = []interface{}{
	(*ComputeRequest)(nil),  // 0: ComputeRequest
	(*Triplet)(nil),         // 1: Triplet
	(*Stat)(nil),            // 2: Stat
	(*ComputeResponse)(nil), // 3: ComputeResponse
}
var file_service_proto_depIdxs = []int32{
	1, // 0: Stat.triplet:type_name -> Triplet
	2, // 1: ComputeResponse.stats:type_name -> Stat
	0, // 2: TripletService.Compute:input_type -> ComputeRequest
	3, // 3: TripletService.Compute:output_type -> ComputeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Triplet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
