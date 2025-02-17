// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: traffic.proto

package traffic

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

type PbTraffic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FD      uint64 `protobuf:"varint,1,opt,name=FD,proto3" json:"FD,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Network string `protobuf:"bytes,3,opt,name=Network,proto3" json:"Network,omitempty"`
}

func (x *PbTraffic) Reset() {
	*x = PbTraffic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbTraffic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbTraffic) ProtoMessage() {}

func (x *PbTraffic) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbTraffic.ProtoReflect.Descriptor instead.
func (*PbTraffic) Descriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{0}
}

func (x *PbTraffic) GetFD() uint64 {
	if x != nil {
		return x.FD
	}
	return 0
}

func (x *PbTraffic) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *PbTraffic) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type PbTraffics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Traffic []*PbTraffic `protobuf:"bytes,1,rep,name=traffic,proto3" json:"traffic,omitempty"`
}

func (x *PbTraffics) Reset() {
	*x = PbTraffics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbTraffics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbTraffics) ProtoMessage() {}

func (x *PbTraffics) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbTraffics.ProtoReflect.Descriptor instead.
func (*PbTraffics) Descriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{1}
}

func (x *PbTraffics) GetTraffic() []*PbTraffic {
	if x != nil {
		return x.Traffic
	}
	return nil
}

var File_traffic_proto protoreflect.FileDescriptor

var file_traffic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x09, 0x50, 0x62, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x46, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x46, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x22, 0x3a, 0x0a, 0x0a, 0x50, 0x62, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x73, 0x12, 0x2c, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x62, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traffic_proto_rawDescOnce sync.Once
	file_traffic_proto_rawDescData = file_traffic_proto_rawDesc
)

func file_traffic_proto_rawDescGZIP() []byte {
	file_traffic_proto_rawDescOnce.Do(func() {
		file_traffic_proto_rawDescData = protoimpl.X.CompressGZIP(file_traffic_proto_rawDescData)
	})
	return file_traffic_proto_rawDescData
}

var file_traffic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_traffic_proto_goTypes = []interface{}{
	(*PbTraffic)(nil),  // 0: service.PbTraffic
	(*PbTraffics)(nil), // 1: service.PbTraffics
}
var file_traffic_proto_depIdxs = []int32{
	0, // 0: service.PbTraffics.traffic:type_name -> service.PbTraffic
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_traffic_proto_init() }
func file_traffic_proto_init() {
	if File_traffic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traffic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbTraffic); i {
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
		file_traffic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbTraffics); i {
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
			RawDescriptor: file_traffic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_traffic_proto_goTypes,
		DependencyIndexes: file_traffic_proto_depIdxs,
		MessageInfos:      file_traffic_proto_msgTypes,
	}.Build()
	File_traffic_proto = out.File
	file_traffic_proto_rawDesc = nil
	file_traffic_proto_goTypes = nil
	file_traffic_proto_depIdxs = nil
}
