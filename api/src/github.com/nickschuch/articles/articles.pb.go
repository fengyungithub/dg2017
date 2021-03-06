// Code generated by protoc-gen-go.
// source: articles.proto
// DO NOT EDIT!

/*
Package articles is a generated protocol buffer package.

It is generated from these files:
	articles.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
	Article
*/
package articles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CreateRequest struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type CreateResponse struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type GetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type UpdateRequest struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateRequest) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListResponse struct {
	Articles []*Article `protobuf:"bytes,1,rep,name=articles" json:"articles,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListResponse) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Body  string `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Article) Reset()                    { *m = Article{} }
func (m *Article) String() string            { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()               {}
func (*Article) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Article) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "articles.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "articles.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "articles.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "articles.GetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "articles.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "articles.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "articles.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "articles.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "articles.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "articles.ListResponse")
	proto.RegisterType((*Article)(nil), "articles.Article")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Articles service

type ArticlesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type articlesClient struct {
	cc *grpc.ClientConn
}

func NewArticlesClient(cc *grpc.ClientConn) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Articles service

type ArticlesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterArticlesServer(s *grpc.Server, srv ArticlesServer) {
	s.RegisterService(&_Articles_serviceDesc, srv)
}

func _Articles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Articles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "articles.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Articles_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Articles_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Articles_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Articles_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Articles_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "articles.proto",
}

func init() { proto.RegisterFile("articles.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xdb, 0xa4, 0xf6, 0xcf, 0x57, 0x13, 0xe2, 0x52, 0x35, 0x04, 0xc1, 0x92, 0x53, 0x41,
	0xec, 0xa1, 0x0a, 0x82, 0x34, 0x07, 0xa9, 0xd0, 0x8b, 0xa7, 0x80, 0x0f, 0xd0, 0x9a, 0x3d, 0x2c,
	0x04, 0x13, 0xb3, 0xeb, 0xc1, 0x17, 0xf3, 0xf9, 0xa4, 0xbb, 0xd9, 0x64, 0xb3, 0x7a, 0xb1, 0xb7,
	0xdd, 0x99, 0xf9, 0xcd, 0x37, 0xb3, 0x5f, 0x02, 0x7f, 0x57, 0x09, 0xf6, 0x96, 0x53, 0xbe, 0x2c,
	0xab, 0x42, 0x14, 0x64, 0xac, 0xef, 0xf1, 0x1a, 0xde, 0xa6, 0xa2, 0x3b, 0x41, 0x53, 0xfa, 0xf1,
	0x49, 0xb9, 0x20, 0x37, 0x18, 0xd5, 0xc9, 0xb0, 0x3f, 0xef, 0x2f, 0xa6, 0xab, 0xb3, 0x65, 0x03,
	0x3f, 0xa9, 0x43, 0xaa, 0x2b, 0xe2, 0x04, 0xbe, 0xa6, 0x79, 0x59, 0xbc, 0x73, 0xfa, 0x3f, 0xfc,
	0x0a, 0xd8, 0x52, 0xa1, 0x95, 0x7d, 0x38, 0x2c, 0x93, 0xd4, 0x24, 0x75, 0x58, 0x16, 0x3f, 0x62,
	0x2a, 0xb3, 0xc7, 0x74, 0x5e, 0xc3, 0x7b, 0x2d, 0xb3, 0x63, 0xd7, 0x0a, 0xe0, 0x6b, 0x5a, 0x89,
	0xc7, 0xd7, 0xf0, 0x9e, 0x69, 0x4e, 0xdb, 0x7e, 0xf6, 0xb0, 0x01, 0x7c, 0x5d, 0xd0, 0x20, 0xd3,
	0x17, 0xc6, 0x9b, 0xed, 0x02, 0xb8, 0x2c, 0xe3, 0x61, 0x7f, 0xee, 0x2e, 0x26, 0xe9, 0xe1, 0x18,
	0x27, 0x38, 0x55, 0x05, 0xf5, 0x82, 0xb7, 0x68, 0x6c, 0x91, 0x65, 0x7f, 0xce, 0xd8, 0x3a, 0xb7,
	0xc1, 0xa8, 0x0e, 0xda, 0xc3, 0x90, 0x19, 0x4e, 0x04, 0x13, 0x39, 0x0d, 0x1d, 0x19, 0x52, 0x17,
	0x42, 0x30, 0xd8, 0x17, 0xd9, 0x57, 0xe8, 0xca, 0xa0, 0x3c, 0xaf, 0xbe, 0x1d, 0x8c, 0xeb, 0x2e,
	0x9c, 0x24, 0x18, 0x2a, 0x37, 0xc9, 0x65, 0x2b, 0xdc, 0xf9, 0x3a, 0xa2, 0xf0, 0x77, 0xa2, 0x5e,
	0xb7, 0x47, 0xee, 0xe1, 0x6e, 0xa9, 0x20, 0xb3, 0xb6, 0xa4, 0x35, 0x37, 0x3a, 0xb7, 0xa2, 0x0d,
	0x95, 0x60, 0xa8, 0xde, 0xda, 0x14, 0xed, 0x78, 0x67, 0x8a, 0x5a, 0xb6, 0x48, 0x5c, 0xbd, 0xbb,
	0x89, 0x77, 0xac, 0x32, 0x71, 0xcb, 0xa2, 0x1e, 0x79, 0xc0, 0xe0, 0xe0, 0x01, 0x31, 0xc6, 0x33,
	0x4c, 0x8b, 0x2e, 0xec, 0xb0, 0x06, 0xf7, 0x43, 0xf9, 0x23, 0xdd, 0xfd, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x1f, 0xf5, 0x16, 0x5a, 0x03, 0x00, 0x00,
}
