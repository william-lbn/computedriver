// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: rlet.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NodeType int32

const (
	NodeType_None    NodeType = 0
	NodeType_Compute NodeType = 1
	NodeType_Storage NodeType = 2
	NodeType_Network NodeType = 3
	NodeType_Stack   NodeType = 4
	NodeType_Monitor NodeType = 5
	NodeType_All     NodeType = 6
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "None",
		1: "Compute",
		2: "Storage",
		3: "Network",
		4: "Stack",
		5: "Monitor",
		6: "All",
	}
	NodeType_value = map[string]int32{
		"None":    0,
		"Compute": 1,
		"Storage": 2,
		"Network": 3,
		"Stack":   4,
		"Monitor": 5,
		"All":     6,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_rlet_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_rlet_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_rlet_proto_rawDescGZIP(), []int{0}
}

type ChessmapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string   `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	NodeType NodeType `protobuf:"varint,2,opt,name=NodeType,proto3,enum=rlet.NodeType" json:"NodeType,omitempty"`
}

func (x *ChessmapRequest) Reset() {
	*x = ChessmapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rlet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessmapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessmapRequest) ProtoMessage() {}

func (x *ChessmapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rlet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessmapRequest.ProtoReflect.Descriptor instead.
func (*ChessmapRequest) Descriptor() ([]byte, []int) {
	return file_rlet_proto_rawDescGZIP(), []int{0}
}

func (x *ChessmapRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ChessmapRequest) GetNodeType() NodeType {
	if x != nil {
		return x.NodeType
	}
	return NodeType_None
}

type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeType       NodeType `protobuf:"varint,1,opt,name=NodeType,proto3,enum=rlet.NodeType" json:"NodeType,omitempty"`
	Version        string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	SupervisorList []string `protobuf:"bytes,3,rep,name=SupervisorList,proto3" json:"SupervisorList,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rlet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_rlet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_rlet_proto_rawDescGZIP(), []int{1}
}

func (x *Component) GetNodeType() NodeType {
	if x != nil {
		return x.NodeType
	}
	return NodeType_None
}

func (x *Component) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Component) GetSupervisorList() []string {
	if x != nil {
		return x.SupervisorList
	}
	return nil
}

type ChessmapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId        string       `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	ComponentList []*Component `protobuf:"bytes,2,rep,name=ComponentList,proto3" json:"ComponentList,omitempty"`
}

func (x *ChessmapResponse) Reset() {
	*x = ChessmapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rlet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessmapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessmapResponse) ProtoMessage() {}

func (x *ChessmapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rlet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessmapResponse.ProtoReflect.Descriptor instead.
func (*ChessmapResponse) Descriptor() ([]byte, []int) {
	return file_rlet_proto_rawDescGZIP(), []int{2}
}

func (x *ChessmapResponse) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ChessmapResponse) GetComponentList() []*Component {
	if x != nil {
		return x.ComponentList
	}
	return nil
}

var File_rlet_proto protoreflect.FileDescriptor

var file_rlet_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6c,
	0x65, 0x74, 0x22, 0x55, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x72, 0x6c, 0x65, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x79, 0x0a, 0x09, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x72, 0x6c, 0x65, 0x74, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6c, 0x65, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x5c, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x6c, 0x6c, 0x10, 0x06, 0x32, 0x4b, 0x0a, 0x0b, 0x52, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x6d, 0x61, 0x70, 0x12, 0x15, 0x2e, 0x72, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x6d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x6c, 0x65,
	0x74, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x6d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rlet_proto_rawDescOnce sync.Once
	file_rlet_proto_rawDescData = file_rlet_proto_rawDesc
)

func file_rlet_proto_rawDescGZIP() []byte {
	file_rlet_proto_rawDescOnce.Do(func() {
		file_rlet_proto_rawDescData = protoimpl.X.CompressGZIP(file_rlet_proto_rawDescData)
	})
	return file_rlet_proto_rawDescData
}

var file_rlet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rlet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rlet_proto_goTypes = []interface{}{
	(NodeType)(0),            // 0: rlet.NodeType
	(*ChessmapRequest)(nil),  // 1: rlet.ChessmapRequest
	(*Component)(nil),        // 2: rlet.Component
	(*ChessmapResponse)(nil), // 3: rlet.ChessmapResponse
}
var file_rlet_proto_depIdxs = []int32{
	0, // 0: rlet.ChessmapRequest.NodeType:type_name -> rlet.NodeType
	0, // 1: rlet.Component.NodeType:type_name -> rlet.NodeType
	2, // 2: rlet.ChessmapResponse.ComponentList:type_name -> rlet.Component
	1, // 3: rlet.RletService.GetChessmap:input_type -> rlet.ChessmapRequest
	3, // 4: rlet.RletService.GetChessmap:output_type -> rlet.ChessmapResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rlet_proto_init() }
func file_rlet_proto_init() {
	if File_rlet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rlet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessmapRequest); i {
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
		file_rlet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Component); i {
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
		file_rlet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessmapResponse); i {
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
			RawDescriptor: file_rlet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rlet_proto_goTypes,
		DependencyIndexes: file_rlet_proto_depIdxs,
		EnumInfos:         file_rlet_proto_enumTypes,
		MessageInfos:      file_rlet_proto_msgTypes,
	}.Build()
	File_rlet_proto = out.File
	file_rlet_proto_rawDesc = nil
	file_rlet_proto_goTypes = nil
	file_rlet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RletServiceClient is the client API for RletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RletServiceClient interface {
	GetChessmap(ctx context.Context, in *ChessmapRequest, opts ...grpc.CallOption) (*ChessmapResponse, error)
}

type rletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRletServiceClient(cc grpc.ClientConnInterface) RletServiceClient {
	return &rletServiceClient{cc}
}

func (c *rletServiceClient) GetChessmap(ctx context.Context, in *ChessmapRequest, opts ...grpc.CallOption) (*ChessmapResponse, error) {
	out := new(ChessmapResponse)
	err := c.cc.Invoke(ctx, "/rlet.RletService/GetChessmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RletServiceServer is the server API for RletService service.
type RletServiceServer interface {
	GetChessmap(context.Context, *ChessmapRequest) (*ChessmapResponse, error)
}

// UnimplementedRletServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRletServiceServer struct {
}

func (*UnimplementedRletServiceServer) GetChessmap(context.Context, *ChessmapRequest) (*ChessmapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChessmap not implemented")
}

func RegisterRletServiceServer(s *grpc.Server, srv RletServiceServer) {
	s.RegisterService(&_RletService_serviceDesc, srv)
}

func _RletService_GetChessmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChessmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RletServiceServer).GetChessmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rlet.RletService/GetChessmap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RletServiceServer).GetChessmap(ctx, req.(*ChessmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RletService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rlet.RletService",
	HandlerType: (*RletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChessmap",
			Handler:    _RletService_GetChessmap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rlet.proto",
}
