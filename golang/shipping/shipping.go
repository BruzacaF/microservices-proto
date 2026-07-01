package shipping

import (
	"context"
	"google.golang.org/grpc"
)


// ShippingClient is the client API for Shipping service.
type ShippingClient interface {
	Create(ctx context.Context, in *CreateShippingRequest, opts ...grpc.CallOption) (*CreateShippingResponse, error)
}

type shippingClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingClient(cc grpc.ClientConnInterface) ShippingClient {
	return &shippingClient{cc}
}

func (c *shippingClient) Create(ctx context.Context, in *CreateShippingRequest, opts ...grpc.CallOption) (*CreateShippingResponse, error) {
	out := new(CreateShippingResponse)
	err := c.cc.Invoke(ctx, "/Shipping/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServer is the server API for Shipping service.
type ShippingServer interface {
	Create(context.Context, *CreateShippingRequest) (*CreateShippingResponse, error)
	mustEmbedUnimplementedShippingServer()
}

type UnimplementedShippingServer struct{}

func (UnimplementedShippingServer) Create(context.Context, *CreateShippingRequest) (*CreateShippingResponse, error) {
	return nil, nil
}

func (UnimplementedShippingServer) mustEmbedUnimplementedShippingServer() {}

// RegisterShippingServer registers the Shipping service with a gRPC server.
func RegisterShippingServer(s grpc.ServiceRegistrar, srv ShippingServer) {
	s.RegisterService(&Shipping_ServiceDesc, srv)
}

var Shipping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Shipping",
	HandlerType: (*ShippingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Shipping_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipping/shipping.proto",
}

func _Shipping_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShippingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/Shipping/Create"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).Create(ctx, req.(*CreateShippingRequest))
	}
	return interceptor(ctx, in, info, handler)
}
