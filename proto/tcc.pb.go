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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type RegistryResData struct {
	TxID                 int64    `protobuf:"varint,1,opt,name=TxID,proto3" json:"TxID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryResData) Reset()         { *m = RegistryResData{} }
func (m *RegistryResData) String() string { return proto.CompactTextString(m) }
func (*RegistryResData) ProtoMessage()    {}
func (*RegistryResData) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{1}
}

func (m *RegistryResData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryResData.Unmarshal(m, b)
}
func (m *RegistryResData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryResData.Marshal(b, m, deterministic)
}
func (m *RegistryResData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryResData.Merge(m, src)
}
func (m *RegistryResData) XXX_Size() int {
	return xxx_messageInfo_RegistryResData.Size(m)
}
func (m *RegistryResData) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryResData.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryResData proto.InternalMessageInfo

func (m *RegistryResData) GetTxID() int64 {
	if m != nil {
		return m.TxID
	}
	return 0
}

type RegistryRes struct {
	Code                 int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *RegistryResData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RegistryRes) Reset()         { *m = RegistryRes{} }
func (m *RegistryRes) String() string { return proto.CompactTextString(m) }
func (*RegistryRes) ProtoMessage()    {}
func (*RegistryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{2}
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

func (m *RegistryRes) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegistryRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegistryRes) GetData() *RegistryResData {
	if m != nil {
		return m.Data
	}
	return nil
}

type LogActionData struct {
	Pk                   int64    `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	TxID                 int64    `protobuf:"varint,2,opt,name=txID,proto3" json:"txID,omitempty"`
	EndpointTxID         int64    `protobuf:"varint,3,opt,name=endpointTxID,proto3" json:"endpointTxID,omitempty"`
	EndpointName         string   `protobuf:"bytes,4,opt,name=endpointName,proto3" json:"endpointName,omitempty"`
	Level                int64    `protobuf:"varint,5,opt,name=level,proto3" json:"level,omitempty"`
	Status               int64    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogActionData) Reset()         { *m = LogActionData{} }
func (m *LogActionData) String() string { return proto.CompactTextString(m) }
func (*LogActionData) ProtoMessage()    {}
func (*LogActionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{3}
}

func (m *LogActionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogActionData.Unmarshal(m, b)
}
func (m *LogActionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogActionData.Marshal(b, m, deterministic)
}
func (m *LogActionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogActionData.Merge(m, src)
}
func (m *LogActionData) XXX_Size() int {
	return xxx_messageInfo_LogActionData.Size(m)
}
func (m *LogActionData) XXX_DiscardUnknown() {
	xxx_messageInfo_LogActionData.DiscardUnknown(m)
}

var xxx_messageInfo_LogActionData proto.InternalMessageInfo

func (m *LogActionData) GetPk() int64 {
	if m != nil {
		return m.Pk
	}
	return 0
}

func (m *LogActionData) GetTxID() int64 {
	if m != nil {
		return m.TxID
	}
	return 0
}

func (m *LogActionData) GetEndpointTxID() int64 {
	if m != nil {
		return m.EndpointTxID
	}
	return 0
}

func (m *LogActionData) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

func (m *LogActionData) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LogActionData) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type LogActionReq struct {
	Data                 []*LogActionData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LogActionReq) Reset()         { *m = LogActionReq{} }
func (m *LogActionReq) String() string { return proto.CompactTextString(m) }
func (*LogActionReq) ProtoMessage()    {}
func (*LogActionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc7e9ebe8c6643, []int{4}
}

func (m *LogActionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogActionReq.Unmarshal(m, b)
}
func (m *LogActionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogActionReq.Marshal(b, m, deterministic)
}
func (m *LogActionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogActionReq.Merge(m, src)
}
func (m *LogActionReq) XXX_Size() int {
	return xxx_messageInfo_LogActionReq.Size(m)
}
func (m *LogActionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogActionReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogActionReq proto.InternalMessageInfo

func (m *LogActionReq) GetData() []*LogActionData {
	if m != nil {
		return m.Data
	}
	return nil
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
	return fileDescriptor_35dc7e9ebe8c6643, []int{5}
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
	return fileDescriptor_35dc7e9ebe8c6643, []int{6}
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
	proto.RegisterType((*Empty)(nil), "tcc.empty")
	proto.RegisterType((*RegistryResData)(nil), "tcc.registryResData")
	proto.RegisterType((*RegistryRes)(nil), "tcc.registryRes")
	proto.RegisterType((*LogActionData)(nil), "tcc.logActionData")
	proto.RegisterType((*LogActionReq)(nil), "tcc.logActionReq")
	proto.RegisterType((*TccReq)(nil), "tcc.tccReq")
	proto.RegisterMapType((map[string]string)(nil), "tcc.tccReq.DataEntry")
	proto.RegisterType((*TccRes)(nil), "tcc.tccRes")
	proto.RegisterMapType((map[string]string)(nil), "tcc.tccRes.DataEntry")
}

func init() { proto.RegisterFile("tcc.proto", fileDescriptor_35dc7e9ebe8c6643) }

var fileDescriptor_35dc7e9ebe8c6643 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6e, 0x9b, 0x40,
	0x10, 0xc7, 0xb5, 0x60, 0x68, 0x19, 0xdc, 0xd6, 0x5d, 0xb9, 0x15, 0xf2, 0x09, 0x21, 0xd5, 0xc2,
	0xaa, 0x84, 0x25, 0x5b, 0xea, 0xe7, 0x29, 0x91, 0x73, 0x88, 0x14, 0xe5, 0xb0, 0xf1, 0x29, 0x52,
	0x0e, 0x78, 0xbd, 0x22, 0xc4, 0x7c, 0x19, 0x16, 0x2b, 0x7e, 0x8d, 0xbc, 0x45, 0xde, 0x32, 0xda,
	0x81, 0x58, 0xc6, 0xa7, 0x1c, 0x72, 0x81, 0xd9, 0xd9, 0xff, 0x32, 0xbf, 0xf9, 0xef, 0x00, 0x96,
	0xe4, 0x3c, 0x28, 0xca, 0x5c, 0xe6, 0x54, 0x97, 0x9c, 0x7b, 0x1f, 0xc0, 0x10, 0x69, 0x21, 0xf7,
	0xde, 0x0f, 0xf8, 0x52, 0x8a, 0x28, 0xae, 0x64, 0xb9, 0x67, 0xa2, 0x5a, 0x84, 0x32, 0xa4, 0x14,
	0x7a, 0xcb, 0xc7, 0xcb, 0x85, 0x43, 0x5c, 0xe2, 0xeb, 0x0c, 0x63, 0xef, 0x0e, 0xec, 0x23, 0x99,
	0x92, 0xf0, 0x7c, 0x2d, 0x50, 0x62, 0x30, 0x8c, 0xe9, 0x00, 0xf4, 0xb4, 0x8a, 0x1c, 0xcd, 0x25,
	0xbe, 0xc5, 0x54, 0x48, 0x7d, 0xe8, 0xad, 0x43, 0x19, 0x3a, 0xba, 0x4b, 0x7c, 0x7b, 0x36, 0x0c,
	0x14, 0xc3, 0x49, 0x31, 0x86, 0x0a, 0xef, 0x99, 0xc0, 0xa7, 0x24, 0x8f, 0xce, 0xb8, 0x8c, 0xf3,
	0x0c, 0x21, 0x3e, 0x83, 0x56, 0x6c, 0x5a, 0x04, 0xad, 0xd8, 0xa8, 0x8a, 0x52, 0x41, 0x69, 0x0d,
	0x94, 0x8a, 0xa9, 0x07, 0x7d, 0x91, 0xad, 0x8b, 0x3c, 0xce, 0x24, 0x02, 0xeb, 0xb8, 0xd7, 0xc9,
	0x1d, 0x6b, 0xae, 0xc3, 0x54, 0x38, 0x3d, 0xc4, 0xeb, 0xe4, 0xe8, 0x10, 0x8c, 0x44, 0xec, 0x44,
	0xe2, 0x18, 0xf8, 0x81, 0x66, 0x41, 0xbf, 0x83, 0x59, 0xc9, 0x50, 0xd6, 0x95, 0x63, 0x62, 0xba,
	0x5d, 0x79, 0xbf, 0xa0, 0x7f, 0x40, 0x65, 0x62, 0x4b, 0xc7, 0x6d, 0x97, 0xc4, 0xd5, 0x7d, 0x7b,
	0x46, 0xb1, 0xcb, 0x4e, 0x2f, 0x6d, 0x8f, 0x09, 0x98, 0x92, 0x73, 0x75, 0x62, 0xd2, 0x39, 0xf1,
	0x0d, 0x4f, 0x34, 0x5b, 0x81, 0xd2, 0x5f, 0x64, 0xca, 0x1f, 0x94, 0x8c, 0x7e, 0x83, 0x75, 0x48,
	0x29, 0x87, 0x37, 0x62, 0x8f, 0xa6, 0x58, 0x4c, 0x85, 0x8a, 0x7c, 0x17, 0x26, 0xb5, 0x68, 0x5d,
	0x6f, 0x16, 0xff, 0xb4, 0x3f, 0xc4, 0x7b, 0x22, 0x6d, 0xb9, 0xb7, 0x5e, 0xd6, 0xe4, 0x70, 0x59,
	0x27, 0x50, 0xd5, 0xbb, 0x41, 0xcd, 0x56, 0x60, 0x2d, 0x39, 0xbf, 0x11, 0xe5, 0x4e, 0x94, 0xd4,
	0x87, 0x8f, 0xac, 0x1d, 0x06, 0x0a, 0x58, 0x0e, 0x27, 0x72, 0x34, 0x38, 0x9d, 0x13, 0xfa, 0x13,
	0xac, 0xab, 0x57, 0x43, 0xe9, 0xd7, 0xae, 0xc1, 0x4c, 0x6c, 0x47, 0xf6, 0x11, 0xec, 0xb9, 0x7f,
	0x3b, 0x8e, 0x62, 0x79, 0x5f, 0xaf, 0x02, 0x9e, 0xa7, 0xd3, 0x24, 0xae, 0x1f, 0xea, 0xac, 0x14,
	0xd9, 0xdf, 0xf9, 0x54, 0x72, 0x3e, 0xc5, 0x5f, 0xe0, 0x3f, 0x3e, 0x57, 0x26, 0xbe, 0xe6, 0x2f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xa6, 0xce, 0x6b, 0x1c, 0x03, 0x00, 0x00,
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
	Registry(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegistryRes, error)
	LogAction(ctx context.Context, in *LogActionReq, opts ...grpc.CallOption) (*TccRes, error)
}

type tccServerClient struct {
	cc *grpc.ClientConn
}

func NewTccServerClient(cc *grpc.ClientConn) TccServerClient {
	return &tccServerClient{cc}
}

func (c *tccServerClient) Registry(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegistryRes, error) {
	out := new(RegistryRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tccServerClient) LogAction(ctx context.Context, in *LogActionReq, opts ...grpc.CallOption) (*TccRes, error) {
	out := new(TccRes)
	err := c.cc.Invoke(ctx, "/tcc.TccServer/LogAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TccServerServer is the server API for TccServer service.
type TccServerServer interface {
	Registry(context.Context, *Empty) (*RegistryRes, error)
	LogAction(context.Context, *LogActionReq) (*TccRes, error)
}

// UnimplementedTccServerServer can be embedded to have forward compatible implementations.
type UnimplementedTccServerServer struct {
}

func (*UnimplementedTccServerServer) Registry(ctx context.Context, req *Empty) (*RegistryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registry not implemented")
}
func (*UnimplementedTccServerServer) LogAction(ctx context.Context, req *LogActionReq) (*TccRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogAction not implemented")
}

func RegisterTccServerServer(s *grpc.Server, srv TccServerServer) {
	s.RegisterService(&_TccServer_serviceDesc, srv)
}

func _TccServer_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(TccServerServer).Registry(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TccServer_LogAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TccServerServer).LogAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcc.TccServer/LogAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TccServerServer).LogAction(ctx, req.(*LogActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TccServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tcc.TccServer",
	HandlerType: (*TccServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registry",
			Handler:    _TccServer_Registry_Handler,
		},
		{
			MethodName: "LogAction",
			Handler:    _TccServer_LogAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tcc.proto",
}
