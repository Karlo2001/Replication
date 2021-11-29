// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: Replication.proto

package Replication

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

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp int32  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Bid       int32  `protobuf:"varint,3,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_Replication_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BidRequest) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BidRequest) GetBid() int32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ended      bool   `protobuf:"varint,1,opt,name=ended,proto3" json:"ended,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp  int32  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	HighestBid int32  `protobuf:"varint,4,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_Replication_proto_rawDescGZIP(), []int{1}
}

func (x *Outcome) GetEnded() bool {
	if x != nil {
		return x.Ended
	}
	return false
}

func (x *Outcome) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Outcome) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Outcome) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack       int32 `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Timestamp int32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_Replication_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetAck() int32 {
	if x != nil {
		return x.Ack
	}
	return 0
}

func (x *Ack) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_Replication_proto_rawDescGZIP(), []int{3}
}

var File_Replication_proto protoreflect.FileDescriptor

var file_Replication_proto_rawDesc = []byte{
	0x0a, 0x11, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x50, 0x0a, 0x0a, 0x42, 0x69, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x07, 0x4f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x57,
	0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x12, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12,
	0x26, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x61, 0x72, 0x6c, 0x6f, 0x32, 0x30, 0x30, 0x31, 0x2f,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Replication_proto_rawDescOnce sync.Once
	file_Replication_proto_rawDescData = file_Replication_proto_rawDesc
)

func file_Replication_proto_rawDescGZIP() []byte {
	file_Replication_proto_rawDescOnce.Do(func() {
		file_Replication_proto_rawDescData = protoimpl.X.CompressGZIP(file_Replication_proto_rawDescData)
	})
	return file_Replication_proto_rawDescData
}

var file_Replication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Replication_proto_goTypes = []interface{}{
	(*BidRequest)(nil), // 0: main.BidRequest
	(*Outcome)(nil),    // 1: main.Outcome
	(*Ack)(nil),        // 2: main.Ack
	(*Empty)(nil),      // 3: main.Empty
}
var file_Replication_proto_depIdxs = []int32{
	0, // 0: main.Auction.Bid:input_type -> main.BidRequest
	3, // 1: main.Auction.Result:input_type -> main.Empty
	2, // 2: main.Auction.Bid:output_type -> main.Ack
	1, // 3: main.Auction.Result:output_type -> main.Outcome
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Replication_proto_init() }
func file_Replication_proto_init() {
	if File_Replication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Replication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_Replication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_Replication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_Replication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_Replication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Replication_proto_goTypes,
		DependencyIndexes: file_Replication_proto_depIdxs,
		MessageInfos:      file_Replication_proto_msgTypes,
	}.Build()
	File_Replication_proto = out.File
	file_Replication_proto_rawDesc = nil
	file_Replication_proto_goTypes = nil
	file_Replication_proto_depIdxs = nil
}