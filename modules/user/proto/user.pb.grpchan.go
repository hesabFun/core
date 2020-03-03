// Code generated by protoc-gen-grpchan. DO NOT EDIT.
// source: modules/user/proto/user.proto

package userpb

import "github.com/fullstorydev/grpchan"
import "golang.org/x/net/context"
import "google.golang.org/grpc"

func RegisterHandlerUserSystem(reg grpchan.ServiceRegistry, srv UserSystemServer) {
	reg.RegisterService(&_UserSystem_serviceDesc, srv)
}

type userSystemChannelClient struct {
	ch grpchan.Channel
}

func NewUserSystemChannelClient(ch grpchan.Channel) UserSystemClient {
	return &userSystemChannelClient{ch: ch}
}

func (c *userSystemChannelClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) ChangeDisplayName(ctx context.Context, in *ChangeDisplayNameRequest, opts ...grpc.CallOption) (*ChangeDisplayNameResponse, error) {
	out := new(ChangeDisplayNameResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/ChangeDisplayName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemChannelClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.ch.Invoke(ctx, "/user.UserSystem/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}