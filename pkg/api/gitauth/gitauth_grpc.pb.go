//
// This file holds the protobuf definitions
// for the Weave Gitops Enterprise Git Provider Authentication API.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/gitauth/gitauth.proto

package api

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

const (
	GitAuth_Authenticate_FullMethodName              = "/gitauth.v1.GitAuth/Authenticate"
	GitAuth_GetGithubDeviceCode_FullMethodName       = "/gitauth.v1.GitAuth/GetGithubDeviceCode"
	GitAuth_GetGithubAuthStatus_FullMethodName       = "/gitauth.v1.GitAuth/GetGithubAuthStatus"
	GitAuth_GetGitlabAuthURL_FullMethodName          = "/gitauth.v1.GitAuth/GetGitlabAuthURL"
	GitAuth_GetBitbucketServerAuthURL_FullMethodName = "/gitauth.v1.GitAuth/GetBitbucketServerAuthURL"
	GitAuth_AuthorizeBitbucketServer_FullMethodName  = "/gitauth.v1.GitAuth/AuthorizeBitbucketServer"
	GitAuth_AuthorizeGitlab_FullMethodName           = "/gitauth.v1.GitAuth/AuthorizeGitlab"
	GitAuth_GetAzureDevOpsAuthURL_FullMethodName     = "/gitauth.v1.GitAuth/GetAzureDevOpsAuthURL"
	GitAuth_AuthorizeAzureDevOps_FullMethodName      = "/gitauth.v1.GitAuth/AuthorizeAzureDevOps"
	GitAuth_ParseRepoURL_FullMethodName              = "/gitauth.v1.GitAuth/ParseRepoURL"
	GitAuth_ValidateProviderToken_FullMethodName     = "/gitauth.v1.GitAuth/ValidateProviderToken"
)

// GitAuthClient is the client API for GitAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitAuthClient interface {
	//
	// Authenticate generates jwt token using git provider name
	// and git provider token arguments
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	//
	// GetGithubDeviceCode retrieves a temporary device code
	// for Github authentication.
	// This code is used to start the Github device-flow.
	GetGithubDeviceCode(ctx context.Context, in *GetGithubDeviceCodeRequest, opts ...grpc.CallOption) (*GetGithubDeviceCodeResponse, error)
	//
	// GetGithubAuthStatus gets the status of the Github device flow
	// authentication requests. Once the user has completed
	// the Github device flow, an access token will be returned.
	// This token will expire in 15 minutes,
	// after which the user will need to complete the flow again
	// to do Git Provider operations.
	GetGithubAuthStatus(ctx context.Context, in *GetGithubAuthStatusRequest, opts ...grpc.CallOption) (*GetGithubAuthStatusResponse, error)
	//
	// GetGitlabAuthURL returns the URL to initiate a GitLab OAuth PKCE flow.
	// The user must browse to the returned URL to authorize the OAuth callback
	// to the GitOps UI.
	// See the GitLab OAuth docs for more more information:
	// https://docs.gitlab.com/ee/api/oauth2.html#supported-oauth-20-flows
	GetGitlabAuthURL(ctx context.Context, in *GetGitlabAuthURLRequest, opts ...grpc.CallOption) (*GetGitlabAuthURLResponse, error)
	//
	GetBitbucketServerAuthURL(ctx context.Context, in *GetBitbucketServerAuthURLRequest, opts ...grpc.CallOption) (*GetBitbucketServerAuthURLResponse, error)
	//
	AuthorizeBitbucketServer(ctx context.Context, in *AuthorizeBitbucketServerRequest, opts ...grpc.CallOption) (*AuthorizeBitbucketServerResponse, error)
	//
	// AuthorizeGitlab exchanges a GitLab code obtained via OAuth callback.
	// The returned token is useable for authentication
	// with the GitOps server only.
	// See the GitLab OAuth docs for more more information:
	// https://docs.gitlab.com/ee/api/oauth2.html#supported-oauth-20-flows
	AuthorizeGitlab(ctx context.Context, in *AuthorizeGitlabRequest, opts ...grpc.CallOption) (*AuthorizeGitlabResponse, error)
	//
	// GetAzureDevOpsAuthURL returns the Azure DevOps authorization URL
	// used to initiate the OAuth flow.
	GetAzureDevOpsAuthURL(ctx context.Context, in *GetAzureDevOpsAuthURLRequest, opts ...grpc.CallOption) (*GetAzureDevOpsAuthURLResponse, error)
	//
	// AuthorizeAzureDevOps returns a token after a user authorizes
	// Azure DevOps to grant access to their account
	// on behalf of Weave GitOps Enterprise.
	AuthorizeAzureDevOps(ctx context.Context, in *AuthorizeAzureDevOpsRequest, opts ...grpc.CallOption) (*AuthorizeAzureDevOpsResponse, error)
	//
	// ParseRepoURL returns structured data about a git repository URL
	ParseRepoURL(ctx context.Context, in *ParseRepoURLRequest, opts ...grpc.CallOption) (*ParseRepoURLResponse, error)
	//
	// ValidateProviderToken check to see if the git provider token
	// is still valid
	ValidateProviderToken(ctx context.Context, in *ValidateProviderTokenRequest, opts ...grpc.CallOption) (*ValidateProviderTokenResponse, error)
}

type gitAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewGitAuthClient(cc grpc.ClientConnInterface) GitAuthClient {
	return &gitAuthClient{cc}
}

func (c *gitAuthClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, GitAuth_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) GetGithubDeviceCode(ctx context.Context, in *GetGithubDeviceCodeRequest, opts ...grpc.CallOption) (*GetGithubDeviceCodeResponse, error) {
	out := new(GetGithubDeviceCodeResponse)
	err := c.cc.Invoke(ctx, GitAuth_GetGithubDeviceCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) GetGithubAuthStatus(ctx context.Context, in *GetGithubAuthStatusRequest, opts ...grpc.CallOption) (*GetGithubAuthStatusResponse, error) {
	out := new(GetGithubAuthStatusResponse)
	err := c.cc.Invoke(ctx, GitAuth_GetGithubAuthStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) GetGitlabAuthURL(ctx context.Context, in *GetGitlabAuthURLRequest, opts ...grpc.CallOption) (*GetGitlabAuthURLResponse, error) {
	out := new(GetGitlabAuthURLResponse)
	err := c.cc.Invoke(ctx, GitAuth_GetGitlabAuthURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) GetBitbucketServerAuthURL(ctx context.Context, in *GetBitbucketServerAuthURLRequest, opts ...grpc.CallOption) (*GetBitbucketServerAuthURLResponse, error) {
	out := new(GetBitbucketServerAuthURLResponse)
	err := c.cc.Invoke(ctx, GitAuth_GetBitbucketServerAuthURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) AuthorizeBitbucketServer(ctx context.Context, in *AuthorizeBitbucketServerRequest, opts ...grpc.CallOption) (*AuthorizeBitbucketServerResponse, error) {
	out := new(AuthorizeBitbucketServerResponse)
	err := c.cc.Invoke(ctx, GitAuth_AuthorizeBitbucketServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) AuthorizeGitlab(ctx context.Context, in *AuthorizeGitlabRequest, opts ...grpc.CallOption) (*AuthorizeGitlabResponse, error) {
	out := new(AuthorizeGitlabResponse)
	err := c.cc.Invoke(ctx, GitAuth_AuthorizeGitlab_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) GetAzureDevOpsAuthURL(ctx context.Context, in *GetAzureDevOpsAuthURLRequest, opts ...grpc.CallOption) (*GetAzureDevOpsAuthURLResponse, error) {
	out := new(GetAzureDevOpsAuthURLResponse)
	err := c.cc.Invoke(ctx, GitAuth_GetAzureDevOpsAuthURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) AuthorizeAzureDevOps(ctx context.Context, in *AuthorizeAzureDevOpsRequest, opts ...grpc.CallOption) (*AuthorizeAzureDevOpsResponse, error) {
	out := new(AuthorizeAzureDevOpsResponse)
	err := c.cc.Invoke(ctx, GitAuth_AuthorizeAzureDevOps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) ParseRepoURL(ctx context.Context, in *ParseRepoURLRequest, opts ...grpc.CallOption) (*ParseRepoURLResponse, error) {
	out := new(ParseRepoURLResponse)
	err := c.cc.Invoke(ctx, GitAuth_ParseRepoURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitAuthClient) ValidateProviderToken(ctx context.Context, in *ValidateProviderTokenRequest, opts ...grpc.CallOption) (*ValidateProviderTokenResponse, error) {
	out := new(ValidateProviderTokenResponse)
	err := c.cc.Invoke(ctx, GitAuth_ValidateProviderToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitAuthServer is the server API for GitAuth service.
// All implementations must embed UnimplementedGitAuthServer
// for forward compatibility
type GitAuthServer interface {
	//
	// Authenticate generates jwt token using git provider name
	// and git provider token arguments
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	//
	// GetGithubDeviceCode retrieves a temporary device code
	// for Github authentication.
	// This code is used to start the Github device-flow.
	GetGithubDeviceCode(context.Context, *GetGithubDeviceCodeRequest) (*GetGithubDeviceCodeResponse, error)
	//
	// GetGithubAuthStatus gets the status of the Github device flow
	// authentication requests. Once the user has completed
	// the Github device flow, an access token will be returned.
	// This token will expire in 15 minutes,
	// after which the user will need to complete the flow again
	// to do Git Provider operations.
	GetGithubAuthStatus(context.Context, *GetGithubAuthStatusRequest) (*GetGithubAuthStatusResponse, error)
	//
	// GetGitlabAuthURL returns the URL to initiate a GitLab OAuth PKCE flow.
	// The user must browse to the returned URL to authorize the OAuth callback
	// to the GitOps UI.
	// See the GitLab OAuth docs for more more information:
	// https://docs.gitlab.com/ee/api/oauth2.html#supported-oauth-20-flows
	GetGitlabAuthURL(context.Context, *GetGitlabAuthURLRequest) (*GetGitlabAuthURLResponse, error)
	//
	GetBitbucketServerAuthURL(context.Context, *GetBitbucketServerAuthURLRequest) (*GetBitbucketServerAuthURLResponse, error)
	//
	AuthorizeBitbucketServer(context.Context, *AuthorizeBitbucketServerRequest) (*AuthorizeBitbucketServerResponse, error)
	//
	// AuthorizeGitlab exchanges a GitLab code obtained via OAuth callback.
	// The returned token is useable for authentication
	// with the GitOps server only.
	// See the GitLab OAuth docs for more more information:
	// https://docs.gitlab.com/ee/api/oauth2.html#supported-oauth-20-flows
	AuthorizeGitlab(context.Context, *AuthorizeGitlabRequest) (*AuthorizeGitlabResponse, error)
	//
	// GetAzureDevOpsAuthURL returns the Azure DevOps authorization URL
	// used to initiate the OAuth flow.
	GetAzureDevOpsAuthURL(context.Context, *GetAzureDevOpsAuthURLRequest) (*GetAzureDevOpsAuthURLResponse, error)
	//
	// AuthorizeAzureDevOps returns a token after a user authorizes
	// Azure DevOps to grant access to their account
	// on behalf of Weave GitOps Enterprise.
	AuthorizeAzureDevOps(context.Context, *AuthorizeAzureDevOpsRequest) (*AuthorizeAzureDevOpsResponse, error)
	//
	// ParseRepoURL returns structured data about a git repository URL
	ParseRepoURL(context.Context, *ParseRepoURLRequest) (*ParseRepoURLResponse, error)
	//
	// ValidateProviderToken check to see if the git provider token
	// is still valid
	ValidateProviderToken(context.Context, *ValidateProviderTokenRequest) (*ValidateProviderTokenResponse, error)
	mustEmbedUnimplementedGitAuthServer()
}

// UnimplementedGitAuthServer must be embedded to have forward compatible implementations.
type UnimplementedGitAuthServer struct {
}

func (UnimplementedGitAuthServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedGitAuthServer) GetGithubDeviceCode(context.Context, *GetGithubDeviceCodeRequest) (*GetGithubDeviceCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubDeviceCode not implemented")
}
func (UnimplementedGitAuthServer) GetGithubAuthStatus(context.Context, *GetGithubAuthStatusRequest) (*GetGithubAuthStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubAuthStatus not implemented")
}
func (UnimplementedGitAuthServer) GetGitlabAuthURL(context.Context, *GetGitlabAuthURLRequest) (*GetGitlabAuthURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGitlabAuthURL not implemented")
}
func (UnimplementedGitAuthServer) GetBitbucketServerAuthURL(context.Context, *GetBitbucketServerAuthURLRequest) (*GetBitbucketServerAuthURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBitbucketServerAuthURL not implemented")
}
func (UnimplementedGitAuthServer) AuthorizeBitbucketServer(context.Context, *AuthorizeBitbucketServerRequest) (*AuthorizeBitbucketServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeBitbucketServer not implemented")
}
func (UnimplementedGitAuthServer) AuthorizeGitlab(context.Context, *AuthorizeGitlabRequest) (*AuthorizeGitlabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeGitlab not implemented")
}
func (UnimplementedGitAuthServer) GetAzureDevOpsAuthURL(context.Context, *GetAzureDevOpsAuthURLRequest) (*GetAzureDevOpsAuthURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAzureDevOpsAuthURL not implemented")
}
func (UnimplementedGitAuthServer) AuthorizeAzureDevOps(context.Context, *AuthorizeAzureDevOpsRequest) (*AuthorizeAzureDevOpsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeAzureDevOps not implemented")
}
func (UnimplementedGitAuthServer) ParseRepoURL(context.Context, *ParseRepoURLRequest) (*ParseRepoURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseRepoURL not implemented")
}
func (UnimplementedGitAuthServer) ValidateProviderToken(context.Context, *ValidateProviderTokenRequest) (*ValidateProviderTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProviderToken not implemented")
}
func (UnimplementedGitAuthServer) mustEmbedUnimplementedGitAuthServer() {}

// UnsafeGitAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitAuthServer will
// result in compilation errors.
type UnsafeGitAuthServer interface {
	mustEmbedUnimplementedGitAuthServer()
}

func RegisterGitAuthServer(s grpc.ServiceRegistrar, srv GitAuthServer) {
	s.RegisterService(&GitAuth_ServiceDesc, srv)
}

func _GitAuth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_GetGithubDeviceCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGithubDeviceCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).GetGithubDeviceCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_GetGithubDeviceCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).GetGithubDeviceCode(ctx, req.(*GetGithubDeviceCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_GetGithubAuthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGithubAuthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).GetGithubAuthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_GetGithubAuthStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).GetGithubAuthStatus(ctx, req.(*GetGithubAuthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_GetGitlabAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGitlabAuthURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).GetGitlabAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_GetGitlabAuthURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).GetGitlabAuthURL(ctx, req.(*GetGitlabAuthURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_GetBitbucketServerAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBitbucketServerAuthURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).GetBitbucketServerAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_GetBitbucketServerAuthURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).GetBitbucketServerAuthURL(ctx, req.(*GetBitbucketServerAuthURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_AuthorizeBitbucketServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeBitbucketServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).AuthorizeBitbucketServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_AuthorizeBitbucketServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).AuthorizeBitbucketServer(ctx, req.(*AuthorizeBitbucketServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_AuthorizeGitlab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeGitlabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).AuthorizeGitlab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_AuthorizeGitlab_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).AuthorizeGitlab(ctx, req.(*AuthorizeGitlabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_GetAzureDevOpsAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAzureDevOpsAuthURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).GetAzureDevOpsAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_GetAzureDevOpsAuthURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).GetAzureDevOpsAuthURL(ctx, req.(*GetAzureDevOpsAuthURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_AuthorizeAzureDevOps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeAzureDevOpsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).AuthorizeAzureDevOps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_AuthorizeAzureDevOps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).AuthorizeAzureDevOps(ctx, req.(*AuthorizeAzureDevOpsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_ParseRepoURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRepoURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).ParseRepoURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_ParseRepoURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).ParseRepoURL(ctx, req.(*ParseRepoURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitAuth_ValidateProviderToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateProviderTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitAuthServer).ValidateProviderToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitAuth_ValidateProviderToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitAuthServer).ValidateProviderToken(ctx, req.(*ValidateProviderTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GitAuth_ServiceDesc is the grpc.ServiceDesc for GitAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitauth.v1.GitAuth",
	HandlerType: (*GitAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _GitAuth_Authenticate_Handler,
		},
		{
			MethodName: "GetGithubDeviceCode",
			Handler:    _GitAuth_GetGithubDeviceCode_Handler,
		},
		{
			MethodName: "GetGithubAuthStatus",
			Handler:    _GitAuth_GetGithubAuthStatus_Handler,
		},
		{
			MethodName: "GetGitlabAuthURL",
			Handler:    _GitAuth_GetGitlabAuthURL_Handler,
		},
		{
			MethodName: "GetBitbucketServerAuthURL",
			Handler:    _GitAuth_GetBitbucketServerAuthURL_Handler,
		},
		{
			MethodName: "AuthorizeBitbucketServer",
			Handler:    _GitAuth_AuthorizeBitbucketServer_Handler,
		},
		{
			MethodName: "AuthorizeGitlab",
			Handler:    _GitAuth_AuthorizeGitlab_Handler,
		},
		{
			MethodName: "GetAzureDevOpsAuthURL",
			Handler:    _GitAuth_GetAzureDevOpsAuthURL_Handler,
		},
		{
			MethodName: "AuthorizeAzureDevOps",
			Handler:    _GitAuth_AuthorizeAzureDevOps_Handler,
		},
		{
			MethodName: "ParseRepoURL",
			Handler:    _GitAuth_ParseRepoURL_Handler,
		},
		{
			MethodName: "ValidateProviderToken",
			Handler:    _GitAuth_ValidateProviderToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/gitauth/gitauth.proto",
}
