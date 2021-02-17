// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserGroupsClient is the client API for UserGroups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGroupsClient interface {
	UserGroups(ctx context.Context, in *AllUGsRequest, opts ...grpc.CallOption) (*AllUGsResult, error)
	CreateUserGroup(ctx context.Context, in *CreateUGRequest, opts ...grpc.CallOption) (*CreateUGResult, error)
	UpdateUserGroup(ctx context.Context, in *UpdateUGRequest, opts ...grpc.CallOption) (*UpdateUGResult, error)
	DeleteUserGroup(ctx context.Context, in *DeleteUGRequest, opts ...grpc.CallOption) (*DeleteUGResult, error)
}

type userGroupsClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGroupsClient(cc grpc.ClientConnInterface) UserGroupsClient {
	return &userGroupsClient{cc}
}

func (c *userGroupsClient) UserGroups(ctx context.Context, in *AllUGsRequest, opts ...grpc.CallOption) (*AllUGsResult, error) {
	out := new(AllUGsResult)
	err := c.cc.Invoke(ctx, "/api.UserGroups/UserGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupsClient) CreateUserGroup(ctx context.Context, in *CreateUGRequest, opts ...grpc.CallOption) (*CreateUGResult, error) {
	out := new(CreateUGResult)
	err := c.cc.Invoke(ctx, "/api.UserGroups/CreateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupsClient) UpdateUserGroup(ctx context.Context, in *UpdateUGRequest, opts ...grpc.CallOption) (*UpdateUGResult, error) {
	out := new(UpdateUGResult)
	err := c.cc.Invoke(ctx, "/api.UserGroups/UpdateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupsClient) DeleteUserGroup(ctx context.Context, in *DeleteUGRequest, opts ...grpc.CallOption) (*DeleteUGResult, error) {
	out := new(DeleteUGResult)
	err := c.cc.Invoke(ctx, "/api.UserGroups/DeleteUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGroupsServer is the server API for UserGroups service.
// All implementations must embed UnimplementedUserGroupsServer
// for forward compatibility
type UserGroupsServer interface {
	UserGroups(context.Context, *AllUGsRequest) (*AllUGsResult, error)
	CreateUserGroup(context.Context, *CreateUGRequest) (*CreateUGResult, error)
	UpdateUserGroup(context.Context, *UpdateUGRequest) (*UpdateUGResult, error)
	DeleteUserGroup(context.Context, *DeleteUGRequest) (*DeleteUGResult, error)
	mustEmbedUnimplementedUserGroupsServer()
}

// UnimplementedUserGroupsServer must be embedded to have forward compatible implementations.
type UnimplementedUserGroupsServer struct {
}

func (UnimplementedUserGroupsServer) UserGroups(context.Context, *AllUGsRequest) (*AllUGsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGroups not implemented")
}
func (UnimplementedUserGroupsServer) CreateUserGroup(context.Context, *CreateUGRequest) (*CreateUGResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserGroup not implemented")
}
func (UnimplementedUserGroupsServer) UpdateUserGroup(context.Context, *UpdateUGRequest) (*UpdateUGResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroup not implemented")
}
func (UnimplementedUserGroupsServer) DeleteUserGroup(context.Context, *DeleteUGRequest) (*DeleteUGResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGroup not implemented")
}
func (UnimplementedUserGroupsServer) mustEmbedUnimplementedUserGroupsServer() {}

// UnsafeUserGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGroupsServer will
// result in compilation errors.
type UnsafeUserGroupsServer interface {
	mustEmbedUnimplementedUserGroupsServer()
}

func RegisterUserGroupsServer(s grpc.ServiceRegistrar, srv UserGroupsServer) {
	s.RegisterService(&UserGroups_ServiceDesc, srv)
}

func _UserGroups_UserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllUGsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupsServer).UserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserGroups/UserGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupsServer).UserGroups(ctx, req.(*AllUGsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroups_CreateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupsServer).CreateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserGroups/CreateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupsServer).CreateUserGroup(ctx, req.(*CreateUGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroups_UpdateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupsServer).UpdateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserGroups/UpdateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupsServer).UpdateUserGroup(ctx, req.(*UpdateUGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroups_DeleteUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupsServer).DeleteUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserGroups/DeleteUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupsServer).DeleteUserGroup(ctx, req.(*DeleteUGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGroups_ServiceDesc is the grpc.ServiceDesc for UserGroups service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGroups_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserGroups",
	HandlerType: (*UserGroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserGroups",
			Handler:    _UserGroups_UserGroups_Handler,
		},
		{
			MethodName: "CreateUserGroup",
			Handler:    _UserGroups_CreateUserGroup_Handler,
		},
		{
			MethodName: "UpdateUserGroup",
			Handler:    _UserGroups_UpdateUserGroup_Handler,
		},
		{
			MethodName: "DeleteUserGroup",
			Handler:    _UserGroups_DeleteUserGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user_group.proto",
}
