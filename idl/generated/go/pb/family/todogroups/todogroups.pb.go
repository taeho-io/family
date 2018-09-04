// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/family/todogroups/todogroups.proto

package todogroups // import "github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type PermitType int32

const (
	PermitType_PERMIT_TYPE_OWNER  PermitType = 0
	PermitType_PERMIT_TYPE_EDITOR PermitType = 1
	PermitType_PERMIT_TYPE_VIEWER PermitType = 2
)

var PermitType_name = map[int32]string{
	0: "PERMIT_TYPE_OWNER",
	1: "PERMIT_TYPE_EDITOR",
	2: "PERMIT_TYPE_VIEWER",
}
var PermitType_value = map[string]int32{
	"PERMIT_TYPE_OWNER":  0,
	"PERMIT_TYPE_EDITOR": 1,
	"PERMIT_TYPE_VIEWER": 2,
}

func (x PermitType) String() string {
	return proto.EnumName(PermitType_name, int32(x))
}
func (PermitType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{0}
}

type TodoGroup struct {
	TodoGroupId          string     `protobuf:"bytes,1,opt,name=todo_group_id,json=todoGroupId,proto3" json:"todo_group_id,omitempty"`
	Title                string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedBy            string     `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt            int64      `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64      `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Order                string     `protobuf:"bytes,7,opt,name=order,proto3" json:"order,omitempty"`
	PermitType           PermitType `protobuf:"varint,8,opt,name=permit_type,json=permitType,proto3,enum=pb.family.todogroups.PermitType" json:"permit_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TodoGroup) Reset()         { *m = TodoGroup{} }
func (m *TodoGroup) String() string { return proto.CompactTextString(m) }
func (*TodoGroup) ProtoMessage()    {}
func (*TodoGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{0}
}
func (m *TodoGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoGroup.Unmarshal(m, b)
}
func (m *TodoGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoGroup.Marshal(b, m, deterministic)
}
func (dst *TodoGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoGroup.Merge(dst, src)
}
func (m *TodoGroup) XXX_Size() int {
	return xxx_messageInfo_TodoGroup.Size(m)
}
func (m *TodoGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoGroup.DiscardUnknown(m)
}

var xxx_messageInfo_TodoGroup proto.InternalMessageInfo

func (m *TodoGroup) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

func (m *TodoGroup) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TodoGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TodoGroup) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *TodoGroup) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *TodoGroup) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *TodoGroup) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *TodoGroup) GetPermitType() PermitType {
	if m != nil {
		return m.PermitType
	}
	return PermitType_PERMIT_TYPE_OWNER
}

type CreateTodoGroupRequest struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup,proto3" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateTodoGroupRequest) Reset()         { *m = CreateTodoGroupRequest{} }
func (m *CreateTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoGroupRequest) ProtoMessage()    {}
func (*CreateTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{1}
}
func (m *CreateTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoGroupRequest.Unmarshal(m, b)
}
func (m *CreateTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoGroupRequest.Merge(dst, src)
}
func (m *CreateTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoGroupRequest.Size(m)
}
func (m *CreateTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoGroupRequest proto.InternalMessageInfo

func (m *CreateTodoGroupRequest) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type CreateTodoGroupResponse struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup,proto3" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateTodoGroupResponse) Reset()         { *m = CreateTodoGroupResponse{} }
func (m *CreateTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoGroupResponse) ProtoMessage()    {}
func (*CreateTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{2}
}
func (m *CreateTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoGroupResponse.Unmarshal(m, b)
}
func (m *CreateTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoGroupResponse.Merge(dst, src)
}
func (m *CreateTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoGroupResponse.Size(m)
}
func (m *CreateTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoGroupResponse proto.InternalMessageInfo

func (m *CreateTodoGroupResponse) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type ListTodoGroupsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoGroupsRequest) Reset()         { *m = ListTodoGroupsRequest{} }
func (m *ListTodoGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodoGroupsRequest) ProtoMessage()    {}
func (*ListTodoGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{3}
}
func (m *ListTodoGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoGroupsRequest.Unmarshal(m, b)
}
func (m *ListTodoGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoGroupsRequest.Marshal(b, m, deterministic)
}
func (dst *ListTodoGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoGroupsRequest.Merge(dst, src)
}
func (m *ListTodoGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodoGroupsRequest.Size(m)
}
func (m *ListTodoGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoGroupsRequest proto.InternalMessageInfo

type ListTodoGroupsResponse struct {
	TodoGroups           []*TodoGroup `protobuf:"bytes,1,rep,name=todo_groups,json=todoGroups,proto3" json:"todo_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListTodoGroupsResponse) Reset()         { *m = ListTodoGroupsResponse{} }
func (m *ListTodoGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoGroupsResponse) ProtoMessage()    {}
func (*ListTodoGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{4}
}
func (m *ListTodoGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoGroupsResponse.Unmarshal(m, b)
}
func (m *ListTodoGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoGroupsResponse.Marshal(b, m, deterministic)
}
func (dst *ListTodoGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoGroupsResponse.Merge(dst, src)
}
func (m *ListTodoGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoGroupsResponse.Size(m)
}
func (m *ListTodoGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoGroupsResponse proto.InternalMessageInfo

func (m *ListTodoGroupsResponse) GetTodoGroups() []*TodoGroup {
	if m != nil {
		return m.TodoGroups
	}
	return nil
}

type GetTodoGroupRequest struct {
	TodoGroupId          string   `protobuf:"bytes,1,opt,name=todo_group_id,json=todoGroupId,proto3" json:"todo_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoGroupRequest) Reset()         { *m = GetTodoGroupRequest{} }
func (m *GetTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoGroupRequest) ProtoMessage()    {}
func (*GetTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{5}
}
func (m *GetTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoGroupRequest.Unmarshal(m, b)
}
func (m *GetTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *GetTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoGroupRequest.Merge(dst, src)
}
func (m *GetTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoGroupRequest.Size(m)
}
func (m *GetTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoGroupRequest proto.InternalMessageInfo

func (m *GetTodoGroupRequest) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

type GetTodoGroupResponse struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup,proto3" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetTodoGroupResponse) Reset()         { *m = GetTodoGroupResponse{} }
func (m *GetTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoGroupResponse) ProtoMessage()    {}
func (*GetTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{6}
}
func (m *GetTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoGroupResponse.Unmarshal(m, b)
}
func (m *GetTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *GetTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoGroupResponse.Merge(dst, src)
}
func (m *GetTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoGroupResponse.Size(m)
}
func (m *GetTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoGroupResponse proto.InternalMessageInfo

func (m *GetTodoGroupResponse) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type UpdateTodoGroupRequest struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup,proto3" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateTodoGroupRequest) Reset()         { *m = UpdateTodoGroupRequest{} }
func (m *UpdateTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoGroupRequest) ProtoMessage()    {}
func (*UpdateTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{7}
}
func (m *UpdateTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoGroupRequest.Unmarshal(m, b)
}
func (m *UpdateTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoGroupRequest.Merge(dst, src)
}
func (m *UpdateTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoGroupRequest.Size(m)
}
func (m *UpdateTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoGroupRequest proto.InternalMessageInfo

func (m *UpdateTodoGroupRequest) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type UpdateTodoGroupResponse struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup,proto3" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateTodoGroupResponse) Reset()         { *m = UpdateTodoGroupResponse{} }
func (m *UpdateTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoGroupResponse) ProtoMessage()    {}
func (*UpdateTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{8}
}
func (m *UpdateTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoGroupResponse.Unmarshal(m, b)
}
func (m *UpdateTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoGroupResponse.Merge(dst, src)
}
func (m *UpdateTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoGroupResponse.Size(m)
}
func (m *UpdateTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoGroupResponse proto.InternalMessageInfo

func (m *UpdateTodoGroupResponse) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type DeleteTodoGroupRequest struct {
	TodoGroupId          string   `protobuf:"bytes,1,opt,name=todo_group_id,json=todoGroupId,proto3" json:"todo_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoGroupRequest) Reset()         { *m = DeleteTodoGroupRequest{} }
func (m *DeleteTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoGroupRequest) ProtoMessage()    {}
func (*DeleteTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{9}
}
func (m *DeleteTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoGroupRequest.Unmarshal(m, b)
}
func (m *DeleteTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoGroupRequest.Merge(dst, src)
}
func (m *DeleteTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoGroupRequest.Size(m)
}
func (m *DeleteTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoGroupRequest proto.InternalMessageInfo

func (m *DeleteTodoGroupRequest) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

type DeleteTodoGroupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoGroupResponse) Reset()         { *m = DeleteTodoGroupResponse{} }
func (m *DeleteTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoGroupResponse) ProtoMessage()    {}
func (*DeleteTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todogroups_5b19f9e6a1aae1c4, []int{10}
}
func (m *DeleteTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoGroupResponse.Unmarshal(m, b)
}
func (m *DeleteTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoGroupResponse.Merge(dst, src)
}
func (m *DeleteTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoGroupResponse.Size(m)
}
func (m *DeleteTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoGroupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TodoGroup)(nil), "pb.family.todogroups.TodoGroup")
	proto.RegisterType((*CreateTodoGroupRequest)(nil), "pb.family.todogroups.CreateTodoGroupRequest")
	proto.RegisterType((*CreateTodoGroupResponse)(nil), "pb.family.todogroups.CreateTodoGroupResponse")
	proto.RegisterType((*ListTodoGroupsRequest)(nil), "pb.family.todogroups.ListTodoGroupsRequest")
	proto.RegisterType((*ListTodoGroupsResponse)(nil), "pb.family.todogroups.ListTodoGroupsResponse")
	proto.RegisterType((*GetTodoGroupRequest)(nil), "pb.family.todogroups.GetTodoGroupRequest")
	proto.RegisterType((*GetTodoGroupResponse)(nil), "pb.family.todogroups.GetTodoGroupResponse")
	proto.RegisterType((*UpdateTodoGroupRequest)(nil), "pb.family.todogroups.UpdateTodoGroupRequest")
	proto.RegisterType((*UpdateTodoGroupResponse)(nil), "pb.family.todogroups.UpdateTodoGroupResponse")
	proto.RegisterType((*DeleteTodoGroupRequest)(nil), "pb.family.todogroups.DeleteTodoGroupRequest")
	proto.RegisterType((*DeleteTodoGroupResponse)(nil), "pb.family.todogroups.DeleteTodoGroupResponse")
	proto.RegisterEnum("pb.family.todogroups.PermitType", PermitType_name, PermitType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoGroupsServiceClient is the client API for TodoGroupsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoGroupsServiceClient interface {
	CreateTodoGroup(ctx context.Context, in *CreateTodoGroupRequest, opts ...grpc.CallOption) (*CreateTodoGroupResponse, error)
	ListTodoGroups(ctx context.Context, in *ListTodoGroupsRequest, opts ...grpc.CallOption) (*ListTodoGroupsResponse, error)
	GetTodoGroup(ctx context.Context, in *GetTodoGroupRequest, opts ...grpc.CallOption) (*GetTodoGroupResponse, error)
	UpdateTodoGroup(ctx context.Context, in *UpdateTodoGroupRequest, opts ...grpc.CallOption) (*UpdateTodoGroupResponse, error)
	DeleteTodoGroup(ctx context.Context, in *DeleteTodoGroupRequest, opts ...grpc.CallOption) (*DeleteTodoGroupResponse, error)
}

type todoGroupsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoGroupsServiceClient(cc *grpc.ClientConn) TodoGroupsServiceClient {
	return &todoGroupsServiceClient{cc}
}

func (c *todoGroupsServiceClient) CreateTodoGroup(ctx context.Context, in *CreateTodoGroupRequest, opts ...grpc.CallOption) (*CreateTodoGroupResponse, error) {
	out := new(CreateTodoGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.family.todogroups.TodoGroupsService/CreateTodoGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) ListTodoGroups(ctx context.Context, in *ListTodoGroupsRequest, opts ...grpc.CallOption) (*ListTodoGroupsResponse, error) {
	out := new(ListTodoGroupsResponse)
	err := c.cc.Invoke(ctx, "/pb.family.todogroups.TodoGroupsService/ListTodoGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) GetTodoGroup(ctx context.Context, in *GetTodoGroupRequest, opts ...grpc.CallOption) (*GetTodoGroupResponse, error) {
	out := new(GetTodoGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.family.todogroups.TodoGroupsService/GetTodoGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) UpdateTodoGroup(ctx context.Context, in *UpdateTodoGroupRequest, opts ...grpc.CallOption) (*UpdateTodoGroupResponse, error) {
	out := new(UpdateTodoGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.family.todogroups.TodoGroupsService/UpdateTodoGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) DeleteTodoGroup(ctx context.Context, in *DeleteTodoGroupRequest, opts ...grpc.CallOption) (*DeleteTodoGroupResponse, error) {
	out := new(DeleteTodoGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.family.todogroups.TodoGroupsService/DeleteTodoGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoGroupsServiceServer is the server API for TodoGroupsService service.
type TodoGroupsServiceServer interface {
	CreateTodoGroup(context.Context, *CreateTodoGroupRequest) (*CreateTodoGroupResponse, error)
	ListTodoGroups(context.Context, *ListTodoGroupsRequest) (*ListTodoGroupsResponse, error)
	GetTodoGroup(context.Context, *GetTodoGroupRequest) (*GetTodoGroupResponse, error)
	UpdateTodoGroup(context.Context, *UpdateTodoGroupRequest) (*UpdateTodoGroupResponse, error)
	DeleteTodoGroup(context.Context, *DeleteTodoGroupRequest) (*DeleteTodoGroupResponse, error)
}

func RegisterTodoGroupsServiceServer(s *grpc.Server, srv TodoGroupsServiceServer) {
	s.RegisterService(&_TodoGroupsService_serviceDesc, srv)
}

func _TodoGroupsService_CreateTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).CreateTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todogroups.TodoGroupsService/CreateTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).CreateTodoGroup(ctx, req.(*CreateTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_ListTodoGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).ListTodoGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todogroups.TodoGroupsService/ListTodoGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).ListTodoGroups(ctx, req.(*ListTodoGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_GetTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).GetTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todogroups.TodoGroupsService/GetTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).GetTodoGroup(ctx, req.(*GetTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_UpdateTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).UpdateTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todogroups.TodoGroupsService/UpdateTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).UpdateTodoGroup(ctx, req.(*UpdateTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_DeleteTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).DeleteTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todogroups.TodoGroupsService/DeleteTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).DeleteTodoGroup(ctx, req.(*DeleteTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoGroupsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.family.todogroups.TodoGroupsService",
	HandlerType: (*TodoGroupsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodoGroup",
			Handler:    _TodoGroupsService_CreateTodoGroup_Handler,
		},
		{
			MethodName: "ListTodoGroups",
			Handler:    _TodoGroupsService_ListTodoGroups_Handler,
		},
		{
			MethodName: "GetTodoGroup",
			Handler:    _TodoGroupsService_GetTodoGroup_Handler,
		},
		{
			MethodName: "UpdateTodoGroup",
			Handler:    _TodoGroupsService_UpdateTodoGroup_Handler,
		},
		{
			MethodName: "DeleteTodoGroup",
			Handler:    _TodoGroupsService_DeleteTodoGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/family/todogroups/todogroups.proto",
}

func init() {
	proto.RegisterFile("pb/family/todogroups/todogroups.proto", fileDescriptor_todogroups_5b19f9e6a1aae1c4)
}

var fileDescriptor_todogroups_5b19f9e6a1aae1c4 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x4d, 0x6f, 0x12, 0x4f,
	0x1c, 0xc7, 0xff, 0x43, 0xff, 0xad, 0xf2, 0x43, 0x4b, 0x3b, 0x52, 0x58, 0x89, 0xa6, 0xb8, 0x89,
	0xb1, 0x62, 0xbb, 0x1b, 0x6b, 0x3c, 0xf8, 0x10, 0x63, 0x6b, 0x49, 0x43, 0xb4, 0xb6, 0xd9, 0x62,
	0x9f, 0x2e, 0x64, 0x61, 0x47, 0x3a, 0x09, 0x30, 0xe3, 0xee, 0xd0, 0x84, 0x18, 0x2f, 0xc6, 0x93,
	0x07, 0x2f, 0x1e, 0x3c, 0x79, 0xf4, 0x15, 0xf9, 0x16, 0x7c, 0x15, 0x9e, 0xcc, 0xce, 0x2e, 0x2c,
	0x2c, 0x83, 0x6e, 0x62, 0xbd, 0xcd, 0xce, 0xef, 0x61, 0x3e, 0x7c, 0xe7, 0xfb, 0x63, 0xe0, 0x26,
	0x6f, 0x98, 0xaf, 0xed, 0x0e, 0x6d, 0xf7, 0x4d, 0xc1, 0x1c, 0xd6, 0x72, 0x59, 0x8f, 0x7b, 0x23,
	0x4b, 0x83, 0xbb, 0x4c, 0x30, 0x9c, 0xe3, 0x0d, 0x23, 0x48, 0x33, 0xa2, 0x58, 0xf1, 0x5a, 0x8b,
	0xb1, 0x56, 0x9b, 0x98, 0x36, 0xa7, 0xa6, 0xdd, 0xed, 0x32, 0x61, 0x0b, 0xca, 0xba, 0x61, 0x8d,
	0xfe, 0x35, 0x05, 0xe9, 0x1a, 0x73, 0xd8, 0xb6, 0x9f, 0x8c, 0x75, 0xb8, 0xec, 0x57, 0xd6, 0x65,
	0x69, 0x9d, 0x3a, 0x1a, 0x2a, 0xa1, 0x95, 0xb4, 0x95, 0x11, 0x83, 0x8c, 0xaa, 0x83, 0x73, 0x30,
	0x2b, 0xa8, 0x68, 0x13, 0x2d, 0x25, 0x63, 0xc1, 0x07, 0x2e, 0x41, 0xc6, 0x21, 0x5e, 0xd3, 0xa5,
	0xdc, 0xef, 0xae, 0xcd, 0x04, 0x75, 0x23, 0x5b, 0xf8, 0x3a, 0x40, 0xd3, 0x25, 0xb6, 0x20, 0x4e,
	0xbd, 0xd1, 0xd7, 0xfe, 0x97, 0x09, 0xe9, 0x70, 0x67, 0xb3, 0x3f, 0x1a, 0xb6, 0x85, 0x36, 0x5b,
	0x42, 0x2b, 0x33, 0xc3, 0xf0, 0x86, 0xf0, 0xc3, 0x3d, 0xee, 0x0c, 0xc2, 0x73, 0x41, 0x38, 0xdc,
	0xd9, 0x10, 0x3e, 0x14, 0x73, 0x1d, 0xe2, 0x6a, 0x17, 0x02, 0x28, 0xf9, 0x81, 0x37, 0x20, 0xc3,
	0x89, 0xdb, 0xa1, 0xa2, 0x2e, 0xfa, 0x9c, 0x68, 0x17, 0x4b, 0x68, 0x65, 0x7e, 0xbd, 0x64, 0xa8,
	0x64, 0x32, 0xf6, 0x64, 0x62, 0xad, 0xcf, 0x89, 0x05, 0x7c, 0xb8, 0xd6, 0x8f, 0x20, 0xff, 0x4c,
	0x42, 0x0c, 0x45, 0xb2, 0xc8, 0x9b, 0x1e, 0xf1, 0x04, 0x7e, 0x02, 0x10, 0x69, 0x25, 0x85, 0xca,
	0xac, 0x2f, 0xab, 0x7b, 0x47, 0xb5, 0xe9, 0xa1, 0x92, 0xfa, 0x31, 0x14, 0x26, 0x3a, 0x7b, 0x9c,
	0x75, 0x3d, 0xf2, 0xd7, 0xad, 0x0b, 0xb0, 0xf4, 0x82, 0x7a, 0x62, 0x18, 0xf3, 0x42, 0x66, 0xfd,
	0x04, 0xf2, 0xf1, 0x40, 0x78, 0xe4, 0x53, 0xc8, 0x44, 0x47, 0x7a, 0x1a, 0x2a, 0xcd, 0x24, 0x39,
	0x13, 0x86, 0x67, 0x7a, 0xfa, 0x03, 0xb8, 0xb2, 0x4d, 0xc4, 0x84, 0x4c, 0x09, 0x2c, 0xa5, 0x1f,
	0x40, 0x6e, 0xbc, 0xf4, 0x9c, 0x74, 0x38, 0x82, 0xfc, 0x2b, 0x69, 0x91, 0x7f, 0x71, 0x79, 0x13,
	0x9d, 0xcf, 0x09, 0xfa, 0x31, 0xe4, 0xb7, 0x48, 0x9b, 0x28, 0xa0, 0x93, 0x48, 0x79, 0x15, 0x0a,
	0x13, 0xd5, 0x01, 0x58, 0x79, 0x1f, 0x20, 0x32, 0x39, 0x5e, 0x82, 0xc5, 0xbd, 0x8a, 0xb5, 0x53,
	0xad, 0xd5, 0x6b, 0xc7, 0x7b, 0x95, 0xfa, 0xee, 0xe1, 0xcb, 0x8a, 0xb5, 0xf0, 0x1f, 0xce, 0x03,
	0x1e, 0xdd, 0xae, 0x6c, 0x55, 0x6b, 0xbb, 0xd6, 0x02, 0x8a, 0xef, 0x1f, 0x54, 0x2b, 0x87, 0x15,
	0x6b, 0x21, 0xb5, 0xfe, 0x73, 0x16, 0x16, 0x23, 0x3b, 0xed, 0x13, 0xf7, 0x8c, 0x36, 0x09, 0xfe,
	0x88, 0x20, 0x1b, 0x33, 0x37, 0x5e, 0x55, 0x6b, 0xa0, 0x9e, 0xae, 0xe2, 0x5a, 0xc2, 0xec, 0xe0,
	0xb7, 0xe9, 0xc5, 0xf7, 0xdf, 0x7f, 0x7c, 0x4e, 0xe5, 0xf4, 0xac, 0x79, 0x76, 0x57, 0xfe, 0x31,
	0xae, 0x05, 0x05, 0x0f, 0x51, 0x19, 0x7f, 0x40, 0x30, 0x3f, 0xee, 0x7a, 0x7c, 0x47, 0xdd, 0x5d,
	0x39, 0x34, 0xc5, 0xd5, 0x64, 0xc9, 0x21, 0x49, 0x41, 0x92, 0x2c, 0xe2, 0x38, 0x09, 0xfe, 0x84,
	0xe0, 0xd2, 0xa8, 0xcb, 0xf1, 0x6d, 0x75, 0x5f, 0xc5, 0x10, 0x15, 0xcb, 0x49, 0x52, 0x43, 0x80,
	0x5b, 0x12, 0xe0, 0x06, 0x5e, 0x8e, 0x01, 0x98, 0x6f, 0xc7, 0xcc, 0xf3, 0x0e, 0x7f, 0x43, 0x90,
	0x8d, 0x99, 0x78, 0xda, 0x25, 0xa9, 0xa7, 0x68, 0xda, 0x25, 0x4d, 0x99, 0x0c, 0xfd, 0xbe, 0x24,
	0x33, 0x8b, 0xe5, 0xdf, 0x90, 0x19, 0xe3, 0x90, 0xfe, 0xfd, 0x7d, 0x41, 0x90, 0x8d, 0x79, 0x7a,
	0x1a, 0xa7, 0x7a, 0x70, 0xa6, 0x71, 0x4e, 0x19, 0x94, 0x81, 0x82, 0xe5, 0x3f, 0x29, 0xb8, 0xb9,
	0x73, 0xf2, 0xbc, 0x45, 0xc5, 0x69, 0xaf, 0x61, 0x34, 0x59, 0xc7, 0x14, 0x36, 0x39, 0x65, 0x6b,
	0x94, 0x0d, 0x9e, 0x6a, 0xea, 0xb4, 0xcd, 0x16, 0xe9, 0x12, 0xd7, 0x7f, 0xa3, 0xcc, 0x16, 0x33,
	0x55, 0xaf, 0xf8, 0xa3, 0x68, 0xd9, 0x98, 0x93, 0x4f, 0xf2, 0xbd, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x26, 0x4f, 0x68, 0xc8, 0xef, 0x07, 0x00, 0x00,
}
