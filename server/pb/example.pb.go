// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/example.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateExampleRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Example              *Example `protobuf:"bytes,2,opt,name=example,proto3" json:"example,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateExampleRequest) Reset()         { *m = CreateExampleRequest{} }
func (m *CreateExampleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateExampleRequest) ProtoMessage()    {}
func (*CreateExampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6418901c248c1961, []int{0}
}

func (m *CreateExampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExampleRequest.Unmarshal(m, b)
}
func (m *CreateExampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExampleRequest.Marshal(b, m, deterministic)
}
func (m *CreateExampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExampleRequest.Merge(m, src)
}
func (m *CreateExampleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateExampleRequest.Size(m)
}
func (m *CreateExampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExampleRequest proto.InternalMessageInfo

func (m *CreateExampleRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateExampleRequest) GetExample() *Example {
	if m != nil {
		return m.Example
	}
	return nil
}

type Example struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_6418901c248c1961, []int{1}
}

func (m *Example) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example.Unmarshal(m, b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example.Marshal(b, m, deterministic)
}
func (m *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(m, src)
}
func (m *Example) XXX_Size() int {
	return xxx_messageInfo_Example.Size(m)
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Example) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ListExamplesRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListExamplesRequest) Reset()         { *m = ListExamplesRequest{} }
func (m *ListExamplesRequest) String() string { return proto.CompactTextString(m) }
func (*ListExamplesRequest) ProtoMessage()    {}
func (*ListExamplesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6418901c248c1961, []int{2}
}

func (m *ListExamplesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExamplesRequest.Unmarshal(m, b)
}
func (m *ListExamplesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExamplesRequest.Marshal(b, m, deterministic)
}
func (m *ListExamplesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExamplesRequest.Merge(m, src)
}
func (m *ListExamplesRequest) XXX_Size() int {
	return xxx_messageInfo_ListExamplesRequest.Size(m)
}
func (m *ListExamplesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExamplesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListExamplesRequest proto.InternalMessageInfo

func (m *ListExamplesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListExamplesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListExamplesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListExamplesResponse struct {
	Examples             []*Example `protobuf:"bytes,1,rep,name=examples,proto3" json:"examples,omitempty"`
	NextPageToken        string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListExamplesResponse) Reset()         { *m = ListExamplesResponse{} }
func (m *ListExamplesResponse) String() string { return proto.CompactTextString(m) }
func (*ListExamplesResponse) ProtoMessage()    {}
func (*ListExamplesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6418901c248c1961, []int{3}
}

func (m *ListExamplesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExamplesResponse.Unmarshal(m, b)
}
func (m *ListExamplesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExamplesResponse.Marshal(b, m, deterministic)
}
func (m *ListExamplesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExamplesResponse.Merge(m, src)
}
func (m *ListExamplesResponse) XXX_Size() int {
	return xxx_messageInfo_ListExamplesResponse.Size(m)
}
func (m *ListExamplesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExamplesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListExamplesResponse proto.InternalMessageInfo

func (m *ListExamplesResponse) GetExamples() []*Example {
	if m != nil {
		return m.Examples
	}
	return nil
}

func (m *ListExamplesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type DeleteExampleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteExampleRequest) Reset()         { *m = DeleteExampleRequest{} }
func (m *DeleteExampleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteExampleRequest) ProtoMessage()    {}
func (*DeleteExampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6418901c248c1961, []int{4}
}

func (m *DeleteExampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteExampleRequest.Unmarshal(m, b)
}
func (m *DeleteExampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteExampleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteExampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteExampleRequest.Merge(m, src)
}
func (m *DeleteExampleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteExampleRequest.Size(m)
}
func (m *DeleteExampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteExampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteExampleRequest proto.InternalMessageInfo

func (m *DeleteExampleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6418901c248c1961, []int{5}
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

func init() {
	proto.RegisterType((*CreateExampleRequest)(nil), "CreateExampleRequest")
	proto.RegisterType((*Example)(nil), "Example")
	proto.RegisterType((*ListExamplesRequest)(nil), "ListExamplesRequest")
	proto.RegisterType((*ListExamplesResponse)(nil), "ListExamplesResponse")
	proto.RegisterType((*DeleteExampleRequest)(nil), "DeleteExampleRequest")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("pb/example.proto", fileDescriptor_6418901c248c1961) }

var fileDescriptor_6418901c248c1961 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x4d, 0x45, 0xbe, 0x46, 0x51, 0xb3, 0x16, 0xd3, 0x60, 0x4c, 0x48, 0x63, 0x0c, 0x31, 0x61,
	0xab, 0x78, 0xf4, 0xa6, 0x72, 0xf3, 0x60, 0x8a, 0x27, 0x2f, 0x64, 0x0b, 0x13, 0xb2, 0xa1, 0x1f,
	0xeb, 0xee, 0x82, 0xc8, 0x9f, 0xf2, 0x2f, 0x9a, 0x2e, 0x0b, 0xa1, 0xb1, 0x89, 0xb7, 0xed, 0x9b,
	0x79, 0x33, 0xef, 0xbd, 0x29, 0x9c, 0x89, 0x28, 0xc0, 0x15, 0x4b, 0x44, 0x8c, 0x54, 0xc8, 0x4c,
	0x67, 0x7e, 0x08, 0xee, 0xb3, 0x44, 0xa6, 0x71, 0xb8, 0x81, 0x43, 0xfc, 0x5c, 0xa0, 0xd2, 0xe4,
	0x02, 0x6a, 0x82, 0x49, 0x4c, 0xb5, 0xe7, 0x74, 0x9d, 0x5e, 0x33, 0xb4, 0x5f, 0xc4, 0x87, 0xba,
	0x1d, 0xe0, 0x1d, 0x74, 0x9d, 0xde, 0xd1, 0xa0, 0x41, 0xb7, 0xcc, 0x6d, 0xc1, 0xbf, 0x87, 0xba,
	0xc5, 0x08, 0x81, 0xc3, 0x94, 0x25, 0x68, 0x87, 0x98, 0x77, 0x8e, 0x69, 0x5c, 0x69, 0xc3, 0x6f,
	0x86, 0xe6, 0xed, 0x73, 0x38, 0x7f, 0xe5, 0x4a, 0x5b, 0x9a, 0xfa, 0x4f, 0xc5, 0x25, 0x34, 0x05,
	0x9b, 0xe1, 0x58, 0xf1, 0xf5, 0x46, 0x47, 0x35, 0x6c, 0xe4, 0xc0, 0x88, 0xaf, 0x91, 0x5c, 0x01,
	0x98, 0xa2, 0xce, 0xe6, 0x98, 0x7a, 0x15, 0x43, 0x34, 0xed, 0xef, 0x39, 0xe0, 0x4f, 0xc1, 0x2d,
	0xae, 0x52, 0x22, 0x4b, 0x15, 0x92, 0x6b, 0x68, 0x58, 0x03, 0xca, 0x73, 0xba, 0x95, 0x82, 0xb5,
	0x5d, 0x85, 0xdc, 0xc0, 0x69, 0x8a, 0x2b, 0x3d, 0xde, 0xdb, 0xb0, 0xf1, 0xd1, 0xca, 0xe1, 0xb7,
	0xdd, 0x96, 0x5b, 0x70, 0x5f, 0x30, 0xc6, 0x3f, 0xb9, 0x96, 0x04, 0xe2, 0xd7, 0xa1, 0x3a, 0x4c,
	0x84, 0xfe, 0x1e, 0xfc, 0x38, 0x70, 0x62, 0xfb, 0x47, 0x28, 0x97, 0x7c, 0x82, 0xe4, 0x0e, 0x5a,
	0x85, 0xfb, 0x90, 0x36, 0x2d, 0xbb, 0x57, 0x67, 0xa7, 0x95, 0x3c, 0xc2, 0xf1, 0xbe, 0x3f, 0xe2,
	0xd2, 0x92, 0x64, 0x3b, 0x6d, 0x5a, 0x1a, 0x02, 0x85, 0x56, 0x41, 0x36, 0x69, 0xd3, 0x32, 0x1b,
	0x9d, 0x1a, 0x35, 0x8a, 0x9f, 0x82, 0x8f, 0xfe, 0x8c, 0xeb, 0x98, 0x45, 0x74, 0x92, 0x25, 0x41,
	0x22, 0xbe, 0x58, 0x3c, 0x47, 0x39, 0xe5, 0x29, 0x06, 0xcb, 0x05, 0xf6, 0xb5, 0xea, 0xcf, 0xa4,
	0x98, 0x04, 0x0a, 0xe5, 0x12, 0x65, 0x20, 0xa2, 0xa8, 0x66, 0x7e, 0xbb, 0x87, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xaa, 0x1c, 0x80, 0xb1, 0x8a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	CreateExample(ctx context.Context, in *CreateExampleRequest, opts ...grpc.CallOption) (*Example, error)
	ListExamples(ctx context.Context, in *ListExamplesRequest, opts ...grpc.CallOption) (*ListExamplesResponse, error)
	DeleteExample(ctx context.Context, in *DeleteExampleRequest, opts ...grpc.CallOption) (*Empty, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) CreateExample(ctx context.Context, in *CreateExampleRequest, opts ...grpc.CallOption) (*Example, error) {
	out := new(Example)
	err := c.cc.Invoke(ctx, "/ExampleService/CreateExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ListExamples(ctx context.Context, in *ListExamplesRequest, opts ...grpc.CallOption) (*ListExamplesResponse, error) {
	out := new(ListExamplesResponse)
	err := c.cc.Invoke(ctx, "/ExampleService/ListExamples", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) DeleteExample(ctx context.Context, in *DeleteExampleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ExampleService/DeleteExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	CreateExample(context.Context, *CreateExampleRequest) (*Example, error)
	ListExamples(context.Context, *ListExamplesRequest) (*ListExamplesResponse, error)
	DeleteExample(context.Context, *DeleteExampleRequest) (*Empty, error)
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_CreateExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).CreateExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleService/CreateExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).CreateExample(ctx, req.(*CreateExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ListExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ListExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleService/ListExamples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ListExamples(ctx, req.(*ListExamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_DeleteExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).DeleteExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleService/DeleteExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).DeleteExample(ctx, req.(*DeleteExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExample",
			Handler:    _ExampleService_CreateExample_Handler,
		},
		{
			MethodName: "ListExamples",
			Handler:    _ExampleService_ListExamples_Handler,
		},
		{
			MethodName: "DeleteExample",
			Handler:    _ExampleService_DeleteExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/example.proto",
}