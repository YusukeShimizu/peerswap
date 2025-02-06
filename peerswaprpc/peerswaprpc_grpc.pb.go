// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: peerswaprpc.proto

package peerswaprpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PeerSwap_SwapOut_FullMethodName             = "/peerswap.PeerSwap/SwapOut"
	PeerSwap_SwapIn_FullMethodName              = "/peerswap.PeerSwap/SwapIn"
	PeerSwap_GetSwap_FullMethodName             = "/peerswap.PeerSwap/GetSwap"
	PeerSwap_ListSwaps_FullMethodName           = "/peerswap.PeerSwap/ListSwaps"
	PeerSwap_ListPeers_FullMethodName           = "/peerswap.PeerSwap/ListPeers"
	PeerSwap_ListRequestedSwaps_FullMethodName  = "/peerswap.PeerSwap/ListRequestedSwaps"
	PeerSwap_ListActiveSwaps_FullMethodName     = "/peerswap.PeerSwap/ListActiveSwaps"
	PeerSwap_AllowSwapRequests_FullMethodName   = "/peerswap.PeerSwap/AllowSwapRequests"
	PeerSwap_ReloadPolicyFile_FullMethodName    = "/peerswap.PeerSwap/ReloadPolicyFile"
	PeerSwap_AddPeer_FullMethodName             = "/peerswap.PeerSwap/AddPeer"
	PeerSwap_RemovePeer_FullMethodName          = "/peerswap.PeerSwap/RemovePeer"
	PeerSwap_AddSusPeer_FullMethodName          = "/peerswap.PeerSwap/AddSusPeer"
	PeerSwap_RemoveSusPeer_FullMethodName       = "/peerswap.PeerSwap/RemoveSusPeer"
	PeerSwap_LiquidGetAddress_FullMethodName    = "/peerswap.PeerSwap/LiquidGetAddress"
	PeerSwap_LiquidGetBalance_FullMethodName    = "/peerswap.PeerSwap/LiquidGetBalance"
	PeerSwap_LiquidSendToAddress_FullMethodName = "/peerswap.PeerSwap/LiquidSendToAddress"
	PeerSwap_Stop_FullMethodName                = "/peerswap.PeerSwap/Stop"
)

// PeerSwapClient is the client API for PeerSwap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerSwapClient interface {
	SwapOut(ctx context.Context, in *SwapOutRequest, opts ...grpc.CallOption) (*SwapResponse, error)
	SwapIn(ctx context.Context, in *SwapInRequest, opts ...grpc.CallOption) (*SwapResponse, error)
	GetSwap(ctx context.Context, in *GetSwapRequest, opts ...grpc.CallOption) (*SwapResponse, error)
	ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error)
	ListPeers(ctx context.Context, in *ListPeersRequest, opts ...grpc.CallOption) (*ListPeersResponse, error)
	ListRequestedSwaps(ctx context.Context, in *ListRequestedSwapsRequest, opts ...grpc.CallOption) (*ListRequestedSwapsResponse, error)
	ListActiveSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error)
	// policy
	AllowSwapRequests(ctx context.Context, in *AllowSwapRequestsRequest, opts ...grpc.CallOption) (*Policy, error)
	ReloadPolicyFile(ctx context.Context, in *ReloadPolicyFileRequest, opts ...grpc.CallOption) (*Policy, error)
	AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*Policy, error)
	RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*Policy, error)
	AddSusPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*Policy, error)
	RemoveSusPeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*Policy, error)
	// Liquid Stuff
	LiquidGetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	LiquidGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	LiquidSendToAddress(ctx context.Context, in *SendToAddressRequest, opts ...grpc.CallOption) (*SendToAddressResponse, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type peerSwapClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerSwapClient(cc grpc.ClientConnInterface) PeerSwapClient {
	return &peerSwapClient{cc}
}

func (c *peerSwapClient) SwapOut(ctx context.Context, in *SwapOutRequest, opts ...grpc.CallOption) (*SwapResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SwapResponse)
	err := c.cc.Invoke(ctx, PeerSwap_SwapOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) SwapIn(ctx context.Context, in *SwapInRequest, opts ...grpc.CallOption) (*SwapResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SwapResponse)
	err := c.cc.Invoke(ctx, PeerSwap_SwapIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) GetSwap(ctx context.Context, in *GetSwapRequest, opts ...grpc.CallOption) (*SwapResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SwapResponse)
	err := c.cc.Invoke(ctx, PeerSwap_GetSwap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSwapsResponse)
	err := c.cc.Invoke(ctx, PeerSwap_ListSwaps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListPeers(ctx context.Context, in *ListPeersRequest, opts ...grpc.CallOption) (*ListPeersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPeersResponse)
	err := c.cc.Invoke(ctx, PeerSwap_ListPeers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListRequestedSwaps(ctx context.Context, in *ListRequestedSwapsRequest, opts ...grpc.CallOption) (*ListRequestedSwapsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRequestedSwapsResponse)
	err := c.cc.Invoke(ctx, PeerSwap_ListRequestedSwaps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListActiveSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSwapsResponse)
	err := c.cc.Invoke(ctx, PeerSwap_ListActiveSwaps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) AllowSwapRequests(ctx context.Context, in *AllowSwapRequestsRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, PeerSwap_AllowSwapRequests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ReloadPolicyFile(ctx context.Context, in *ReloadPolicyFileRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, PeerSwap_ReloadPolicyFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, PeerSwap_AddPeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, PeerSwap_RemovePeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) AddSusPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, PeerSwap_AddSusPeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) RemoveSusPeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, PeerSwap_RemoveSusPeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) LiquidGetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, PeerSwap_LiquidGetAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) LiquidGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, PeerSwap_LiquidGetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) LiquidSendToAddress(ctx context.Context, in *SendToAddressRequest, opts ...grpc.CallOption) (*SendToAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendToAddressResponse)
	err := c.cc.Invoke(ctx, PeerSwap_LiquidSendToAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, PeerSwap_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerSwapServer is the server API for PeerSwap service.
// All implementations must embed UnimplementedPeerSwapServer
// for forward compatibility.
type PeerSwapServer interface {
	SwapOut(context.Context, *SwapOutRequest) (*SwapResponse, error)
	SwapIn(context.Context, *SwapInRequest) (*SwapResponse, error)
	GetSwap(context.Context, *GetSwapRequest) (*SwapResponse, error)
	ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error)
	ListPeers(context.Context, *ListPeersRequest) (*ListPeersResponse, error)
	ListRequestedSwaps(context.Context, *ListRequestedSwapsRequest) (*ListRequestedSwapsResponse, error)
	ListActiveSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error)
	// policy
	AllowSwapRequests(context.Context, *AllowSwapRequestsRequest) (*Policy, error)
	ReloadPolicyFile(context.Context, *ReloadPolicyFileRequest) (*Policy, error)
	AddPeer(context.Context, *AddPeerRequest) (*Policy, error)
	RemovePeer(context.Context, *RemovePeerRequest) (*Policy, error)
	AddSusPeer(context.Context, *AddPeerRequest) (*Policy, error)
	RemoveSusPeer(context.Context, *RemovePeerRequest) (*Policy, error)
	// Liquid Stuff
	LiquidGetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	LiquidGetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	LiquidSendToAddress(context.Context, *SendToAddressRequest) (*SendToAddressResponse, error)
	Stop(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedPeerSwapServer()
}

// UnimplementedPeerSwapServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPeerSwapServer struct{}

func (UnimplementedPeerSwapServer) SwapOut(context.Context, *SwapOutRequest) (*SwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapOut not implemented")
}
func (UnimplementedPeerSwapServer) SwapIn(context.Context, *SwapInRequest) (*SwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapIn not implemented")
}
func (UnimplementedPeerSwapServer) GetSwap(context.Context, *GetSwapRequest) (*SwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwap not implemented")
}
func (UnimplementedPeerSwapServer) ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSwaps not implemented")
}
func (UnimplementedPeerSwapServer) ListPeers(context.Context, *ListPeersRequest) (*ListPeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeers not implemented")
}
func (UnimplementedPeerSwapServer) ListRequestedSwaps(context.Context, *ListRequestedSwapsRequest) (*ListRequestedSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequestedSwaps not implemented")
}
func (UnimplementedPeerSwapServer) ListActiveSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveSwaps not implemented")
}
func (UnimplementedPeerSwapServer) AllowSwapRequests(context.Context, *AllowSwapRequestsRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowSwapRequests not implemented")
}
func (UnimplementedPeerSwapServer) ReloadPolicyFile(context.Context, *ReloadPolicyFileRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadPolicyFile not implemented")
}
func (UnimplementedPeerSwapServer) AddPeer(context.Context, *AddPeerRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedPeerSwapServer) RemovePeer(context.Context, *RemovePeerRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}
func (UnimplementedPeerSwapServer) AddSusPeer(context.Context, *AddPeerRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSusPeer not implemented")
}
func (UnimplementedPeerSwapServer) RemoveSusPeer(context.Context, *RemovePeerRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSusPeer not implemented")
}
func (UnimplementedPeerSwapServer) LiquidGetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidGetAddress not implemented")
}
func (UnimplementedPeerSwapServer) LiquidGetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidGetBalance not implemented")
}
func (UnimplementedPeerSwapServer) LiquidSendToAddress(context.Context, *SendToAddressRequest) (*SendToAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidSendToAddress not implemented")
}
func (UnimplementedPeerSwapServer) Stop(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedPeerSwapServer) mustEmbedUnimplementedPeerSwapServer() {}
func (UnimplementedPeerSwapServer) testEmbeddedByValue()                  {}

// UnsafePeerSwapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerSwapServer will
// result in compilation errors.
type UnsafePeerSwapServer interface {
	mustEmbedUnimplementedPeerSwapServer()
}

func RegisterPeerSwapServer(s grpc.ServiceRegistrar, srv PeerSwapServer) {
	// If the following call pancis, it indicates UnimplementedPeerSwapServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PeerSwap_ServiceDesc, srv)
}

func _PeerSwap_SwapOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwapOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).SwapOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_SwapOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).SwapOut(ctx, req.(*SwapOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_SwapIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwapInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).SwapIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_SwapIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).SwapIn(ctx, req.(*SwapInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_GetSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).GetSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_GetSwap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).GetSwap(ctx, req.(*GetSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_ListSwaps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListSwaps(ctx, req.(*ListSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_ListPeers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListPeers(ctx, req.(*ListPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListRequestedSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequestedSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListRequestedSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_ListRequestedSwaps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListRequestedSwaps(ctx, req.(*ListRequestedSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListActiveSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListActiveSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_ListActiveSwaps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListActiveSwaps(ctx, req.(*ListSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_AllowSwapRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowSwapRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).AllowSwapRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_AllowSwapRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).AllowSwapRequests(ctx, req.(*AllowSwapRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ReloadPolicyFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadPolicyFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ReloadPolicyFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_ReloadPolicyFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ReloadPolicyFile(ctx, req.(*ReloadPolicyFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_AddPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).AddPeer(ctx, req.(*AddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_RemovePeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).RemovePeer(ctx, req.(*RemovePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_AddSusPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).AddSusPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_AddSusPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).AddSusPeer(ctx, req.(*AddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_RemoveSusPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).RemoveSusPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_RemoveSusPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).RemoveSusPeer(ctx, req.(*RemovePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_LiquidGetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).LiquidGetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_LiquidGetAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).LiquidGetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_LiquidGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).LiquidGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_LiquidGetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).LiquidGetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_LiquidSendToAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).LiquidSendToAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_LiquidSendToAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).LiquidSendToAddress(ctx, req.(*SendToAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeerSwap_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerSwap_ServiceDesc is the grpc.ServiceDesc for PeerSwap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerSwap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "peerswap.PeerSwap",
	HandlerType: (*PeerSwapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SwapOut",
			Handler:    _PeerSwap_SwapOut_Handler,
		},
		{
			MethodName: "SwapIn",
			Handler:    _PeerSwap_SwapIn_Handler,
		},
		{
			MethodName: "GetSwap",
			Handler:    _PeerSwap_GetSwap_Handler,
		},
		{
			MethodName: "ListSwaps",
			Handler:    _PeerSwap_ListSwaps_Handler,
		},
		{
			MethodName: "ListPeers",
			Handler:    _PeerSwap_ListPeers_Handler,
		},
		{
			MethodName: "ListRequestedSwaps",
			Handler:    _PeerSwap_ListRequestedSwaps_Handler,
		},
		{
			MethodName: "ListActiveSwaps",
			Handler:    _PeerSwap_ListActiveSwaps_Handler,
		},
		{
			MethodName: "AllowSwapRequests",
			Handler:    _PeerSwap_AllowSwapRequests_Handler,
		},
		{
			MethodName: "ReloadPolicyFile",
			Handler:    _PeerSwap_ReloadPolicyFile_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _PeerSwap_AddPeer_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _PeerSwap_RemovePeer_Handler,
		},
		{
			MethodName: "AddSusPeer",
			Handler:    _PeerSwap_AddSusPeer_Handler,
		},
		{
			MethodName: "RemoveSusPeer",
			Handler:    _PeerSwap_RemoveSusPeer_Handler,
		},
		{
			MethodName: "LiquidGetAddress",
			Handler:    _PeerSwap_LiquidGetAddress_Handler,
		},
		{
			MethodName: "LiquidGetBalance",
			Handler:    _PeerSwap_LiquidGetBalance_Handler,
		},
		{
			MethodName: "LiquidSendToAddress",
			Handler:    _PeerSwap_LiquidSendToAddress_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _PeerSwap_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peerswaprpc.proto",
}
