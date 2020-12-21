// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tcc.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegistryReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryReq) Reset()         { *m = RegistryReq{} }
func (m *RegistryReq) String() string { return proto.CompactTextString(m) }
func (*RegistryReq) ProtoMessage()    {}
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{0}
}

func (m *RegistryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryReq.Unmarshal(m, b)
}
func (m *RegistryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryReq.Marshal(b, m, deterministic)
}
func (m *RegistryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryReq.Merge(m, src)
}
func (m *RegistryReq) XXX_Size() int {
	return xxx_messageInfo_RegistryReq.Size(m)
}
func (m *RegistryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryReq proto.InternalMessageInfo

type RegistryRes struct {
	TxID                 int64    `protobuf:"varint,1,opt,name=TxID,proto3" json:"TxID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryRes) Reset()         { *m = RegistryRes{} }
func (m *RegistryRes) String() string { return proto.CompactTextString(m) }
func (*RegistryRes) ProtoMessage()    {}
func (*RegistryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{1}
}

func (m *RegistryRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryRes.Unmarshal(m, b)
}
func (m *RegistryRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryRes.Marshal(b, m, deterministic)
}
func (m *RegistryRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryRes.Merge(m, src)
}
func (m *RegistryRes) XXX_Size() int {
	return xxx_messageInfo_RegistryRes.Size(m)
}
func (m *RegistryRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryRes.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryRes proto.InternalMessageInfo

func (m *RegistryRes) GetTxID() int64 {
	if m != nil {
		return m.TxID
	}
	return 0
}

type AddEndpointReq struct {
	TxID                 int64    `protobuf:"varint,1,opt,name=TxID,proto3" json:"TxID,omitempty"`
	EndpointTxID         int64    `protobuf:"varint,2,opt,name=EndpointTxID,proto3" json:"EndpointTxID,omitempty"`
	ParamData            string   `protobuf:"bytes,3,opt,name=ParamData,proto3" json:"ParamData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEndpointReq) Reset()         { *m = AddEndpointReq{} }
func (m *AddEndpointReq) String() string { return proto.CompactTextString(m) }
func (*AddEndpointReq) ProtoMessage()    {}
func (*AddEndpointReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{2}
}

func (m *AddEndpointReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEndpointReq.Unmarshal(m, b)
}
func (m *AddEndpointReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEndpointReq.Marshal(b, m, deterministic)
}
func (m *AddEndpointReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEndpointReq.Merge(m, src)
}
func (m *AddEndpointReq) XXX_Size() int {
	return xxx_messageInfo_AddEndpointReq.Size(m)
}
func (m *AddEndpointReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEndpointReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddEndpointReq proto.InternalMessageInfo

func (m *AddEndpointReq) GetTxID() int64 {
	if m != nil {
		return m.TxID
	}
	return 0
}

func (m *AddEndpointReq) GetEndpointTxID() int64 {
	if m != nil {
		return m.EndpointTxID
	}
	return 0
}

func (m *AddEndpointReq) GetParamData() string {
	if m != nil {
		return m.ParamData
	}
	return ""
}

type TccReq struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TccReq) Reset()         { *m = TccReq{} }
func (m *TccReq) String() string { return proto.CompactTextString(m) }
func (*TccReq) ProtoMessage()    {}
func (*TccReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{3}
}

func (m *TccReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TccReq.Unmarshal(m, b)
}
func (m *TccReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TccReq.Marshal(b, m, deterministic)
}
func (m *TccReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TccReq.Merge(m, src)
}
func (m *TccReq) XXX_Size() int {
	return xxx_messageInfo_TccReq.Size(m)
}
func (m *TccReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TccReq.DiscardUnknown(m)
}

var xxx_messageInfo_TccReq proto.InternalMessageInfo

func (m *TccReq) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type TccRes struct {
	Code                 int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TccRes) Reset()         { *m = TccRes{} }
func (m *TccRes) String() string { return proto.CompactTextString(m) }
func (*TccRes) ProtoMessage()    {}
func (*TccRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{4}
}

func (m *TccRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TccRes.Unmarshal(m, b)
}
func (m *TccRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TccRes.Marshal(b, m, deterministic)
}
func (m *TccRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TccRes.Merge(m, src)
}
func (m *TccRes) XXX_Size() int {
	return xxx_messageInfo_TccRes.Size(m)
}
func (m *TccRes) XXX_DiscardUnknown() {
	xxx_messageInfo_TccRes.DiscardUnknown(m)
}

var xxx_messageInfo_TccRes proto.InternalMessageInfo

func (m *TccRes) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TccRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TccRes) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RegistryReq)(nil), "tcc.RegistryReq")
	proto.RegisterType((*RegistryRes)(nil), "tcc.RegistryRes")
	proto.RegisterType((*AddEndpointReq)(nil), "tcc.AddEndpointReq")
	proto.RegisterType((*TccReq)(nil), "tcc.tccReq")
	proto.RegisterMapType((map[string]string)(nil), "tcc.tccReq.DataEntry")
	proto.RegisterType((*TccRes)(nil), "tcc.tccRes")
	proto.RegisterMapType((map[string]string)(nil), "tcc.tccRes.DataEntry")
}

func init() { proto.RegisterFile("tcc.proto", fileDescriptor_35dc7e9ebe8c6643) }

var fileDescriptor_35dc7e9ebe8c6643 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x53, 0x0a, 0xdc, 0xdb, 0xd3, 0x7b, 0x6f, 0xc8, 0x5c, 0x4d, 0x08, 0x31, 0x11, 0xc7,
	0xc4, 0xd4, 0x4d, 0x9b, 0xc0, 0xc2, 0x7f, 0x2b, 0x05, 0x16, 0xee, 0xcc, 0xc8, 0xca, 0x5d, 0x99,
	0x0e, 0x58, 0x6d, 0xa7, 0x30, 0x9d, 0x12, 0x79, 0x0d, 0x1f, 0xce, 0xe7, 0x31, 0x73, 0xf8, 0x63,
	0x21, 0x86, 0x95, 0x9b, 0xf6, 0x9c, 0xf3, 0xfd, 0x66, 0xe6, 0x3b, 0xc9, 0x07, 0x8e, 0xe6, 0xdc,
	0x9f, 0xaa, 0x4c, 0x67, 0xc4, 0xd6, 0x9c, 0xd3, 0xbf, 0xe0, 0x32, 0x31, 0x89, 0x73, 0xad, 0x16,
	0x4c, 0xcc, 0xe8, 0x49, 0xb9, 0xcd, 0x09, 0x81, 0xea, 0xf0, 0xed, 0xbe, 0xdf, 0xb4, 0xda, 0x96,
	0x67, 0x33, 0xac, 0xe9, 0x18, 0xfe, 0xdd, 0x46, 0xd1, 0x40, 0x46, 0xd3, 0x2c, 0x96, 0x9a, 0x89,
	0xd9, 0x77, 0x14, 0xa1, 0xf0, 0x67, 0x8d, 0xa0, 0x56, 0x41, 0x6d, 0x6b, 0x46, 0x8e, 0xc0, 0x79,
	0x08, 0x55, 0x98, 0xf6, 0x43, 0x1d, 0x36, 0xed, 0xb6, 0xe5, 0x39, 0xec, 0x6b, 0x40, 0x13, 0xa8,
	0x6b, 0xce, 0xcd, 0xfd, 0xe7, 0x50, 0x8d, 0x0c, 0x62, 0xb5, 0x6d, 0xcf, 0xed, 0x1c, 0xfa, 0x66,
	0x85, 0xa5, 0xe4, 0x1b, 0x72, 0x20, 0x8d, 0x5b, 0x44, 0x5a, 0x17, 0xe0, 0x6c, 0x46, 0xa4, 0x01,
	0xf6, 0xab, 0x58, 0xa0, 0x2d, 0x87, 0x99, 0x92, 0x1c, 0x40, 0x6d, 0x1e, 0x26, 0x85, 0x40, 0x3b,
	0x0e, 0x5b, 0x36, 0xd7, 0x95, 0x4b, 0x8b, 0xbe, 0x5b, 0xab, 0xe7, 0x70, 0x69, 0x9e, 0x45, 0x02,
	0xcf, 0xd5, 0x18, 0xd6, 0xe6, 0xaa, 0x34, 0x9f, 0xac, 0x8e, 0x99, 0x72, 0x63, 0xca, 0xde, 0x35,
	0x95, 0xff, 0x98, 0xa9, 0xce, 0x87, 0x05, 0xce, 0x90, 0xf3, 0x47, 0xa1, 0xe6, 0x42, 0x91, 0x63,
	0xb0, 0x87, 0x6a, 0x41, 0xdc, 0xd2, 0xfe, 0xad, 0x52, 0x93, 0x13, 0x0a, 0xf5, 0x5e, 0x28, 0xb9,
	0x48, 0xf6, 0x30, 0xa7, 0xf0, 0xab, 0x97, 0xc9, 0x71, 0xac, 0xd2, 0x3d, 0x90, 0x0f, 0xbf, 0xd7,
	0x29, 0x20, 0x0d, 0x14, 0x4a, 0x19, 0x69, 0xed, 0x4e, 0x72, 0x12, 0x80, 0x5b, 0x8a, 0x04, 0xf9,
	0x8f, 0xc0, 0x76, 0x48, 0xb6, 0x1e, 0xb8, 0xf3, 0x9e, 0xce, 0x26, 0xb1, 0x7e, 0x2e, 0x46, 0x3e,
	0xcf, 0xd2, 0x20, 0x89, 0x8b, 0x97, 0x42, 0x2a, 0x21, 0xaf, 0xba, 0x81, 0xe6, 0x3c, 0xc0, 0x78,
	0xde, 0xe0, 0x77, 0x54, 0xc7, 0x5f, 0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x5e, 0x03, 0x08,
	0xb8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TccServerClient is the client API for TccServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TccServerClient interface {
	Try(ctx context.Context, in *TccReq, opts ...grpc.CallOption) (*TccRes, error)
	Cancel(ctx context.Context, in *TccReq, opts ...grpc.CallOption) (*TccRes, error)
	Confirm(ctx context.Context, in *TccReq, opts ...grpc.CallOption) (*TccRes, error)
	Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryRes, error)
	AddEndpoint(ctx context.Context, in *AddEndpointReq, opts ...grpc.CallOption) (*TccRes, error)
}

type tccServerClient struct {
	cc *grpc.ClientConn
}

func NewTccServerClient(cc *grpc.ClientConn) TccServerClient {
	return &tccServerClient{cc}
}

func (c *tccServerClient) Try(ctx context.Context, in *TccReq, opts ...grpc.CallOption) (*TccRes, error) {
	out := new(TccRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/Try", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tccServerClient) Cancel(ctx context.Context, in *TccReq, opts ...grpc.CallOption) (*TccRes, error) {
	out := new(TccRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tccServerClient) Confirm(ctx context.Context, in *TccReq, opts ...grpc.CallOption) (*TccRes, error) {
	out := new(TccRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tccServerClient) Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryRes, error) {
	out := new(RegistryRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tccServerClient) AddEndpoint(ctx context.Context, in *AddEndpointReq, opts ...grpc.CallOption) (*TccRes, error) {
	out := new(TccRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/AddEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TccServerServer is the server API for TccServer service.
type TccServerServer interface {
	Try(context.Context, *TccReq) (*TccRes, error)
	Cancel(context.Context, *TccReq) (*TccRes, error)
	Confirm(context.Context, *TccReq) (*TccRes, error)
	Registry(context.Context, *RegistryReq) (*RegistryRes, error)
	AddEndpoint(context.Context, *AddEndpointReq) (*TccRes, error)
}

// UnimplementedTccServerServer can be embedded to have forward compatible implementations.
type UnimplementedTccServerServer struct {
}

func (*UnimplementedTccServerServer) Try(ctx context.Context, req *TccReq) (*TccRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Try not implemented")
}
func (*UnimplementedTccServerServer) Cancel(ctx context.Context, req *TccReq) (*TccRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedTccServerServer) Confirm(ctx context.Context, req *TccReq) (*TccRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (*UnimplementedTccServerServer) Registry(ctx context.Context, req *RegistryReq) (*RegistryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registry not implemented")
}
func (*UnimplementedTccServerServer) AddEndpoint(ctx context.Context, req *AddEndpointReq) (*TccRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEndpoint not implemented")
}

func RegisterTccServerServer(s *grpc.Server, srv TccServerServer) {
	s.RegisterService(&_TccServer_serviceDesc, srv)
}

func _TccServer_Try_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TccServerServer).Try(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcc.TccServer/Try",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TccServerServer).Try(ctx, req.(*TccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TccServer_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TccServerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcc.TccServer/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TccServerServer).Cancel(ctx, req.(*TccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TccServer_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TccServerServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcc.TccServer/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TccServerServer).Confirm(ctx, req.(*TccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TccServer_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TccServerServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcc.TccServer/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TccServerServer).Registry(ctx, req.(*RegistryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TccServer_AddEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEndpointReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TccServerServer).AddEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcc.TccServer/AddEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TccServerServer).AddEndpoint(ctx, req.(*AddEndpointReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TccServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tcc.TccServer",
	HandlerType: (*TccServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Try",
			Handler:    _TccServer_Try_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _TccServer_Cancel_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _TccServer_Confirm_Handler,
		},
		{
			MethodName: "Registry",
			Handler:    _TccServer_Registry_Handler,
		},
		{
			MethodName: "AddEndpoint",
			Handler:    _TccServer_AddEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tcc.proto",
}
