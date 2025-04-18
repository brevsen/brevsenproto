// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: distribution/v1/distribution.proto

package distributionv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/brevsen/brevsenproto/gen/go/distribution/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// DistributionServiceName is the fully-qualified name of the DistributionService service.
	DistributionServiceName = "distribution.v1.DistributionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DistributionServiceFetchOperationsProcedure is the fully-qualified name of the
	// DistributionService's FetchOperations RPC.
	DistributionServiceFetchOperationsProcedure = "/distribution.v1.DistributionService/FetchOperations"
)

// DistributionServiceClient is a client for the distribution.v1.DistributionService service.
type DistributionServiceClient interface {
	FetchOperations(context.Context, *connect.Request[v1.FetchOperationsRequest]) (*connect.ServerStreamForClient[v1.FetchOperationsResponse], error)
}

// NewDistributionServiceClient constructs a client for the distribution.v1.DistributionService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDistributionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DistributionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	distributionServiceMethods := v1.File_distribution_v1_distribution_proto.Services().ByName("DistributionService").Methods()
	return &distributionServiceClient{
		fetchOperations: connect.NewClient[v1.FetchOperationsRequest, v1.FetchOperationsResponse](
			httpClient,
			baseURL+DistributionServiceFetchOperationsProcedure,
			connect.WithSchema(distributionServiceMethods.ByName("FetchOperations")),
			connect.WithClientOptions(opts...),
		),
	}
}

// distributionServiceClient implements DistributionServiceClient.
type distributionServiceClient struct {
	fetchOperations *connect.Client[v1.FetchOperationsRequest, v1.FetchOperationsResponse]
}

// FetchOperations calls distribution.v1.DistributionService.FetchOperations.
func (c *distributionServiceClient) FetchOperations(ctx context.Context, req *connect.Request[v1.FetchOperationsRequest]) (*connect.ServerStreamForClient[v1.FetchOperationsResponse], error) {
	return c.fetchOperations.CallServerStream(ctx, req)
}

// DistributionServiceHandler is an implementation of the distribution.v1.DistributionService
// service.
type DistributionServiceHandler interface {
	FetchOperations(context.Context, *connect.Request[v1.FetchOperationsRequest], *connect.ServerStream[v1.FetchOperationsResponse]) error
}

// NewDistributionServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDistributionServiceHandler(svc DistributionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	distributionServiceMethods := v1.File_distribution_v1_distribution_proto.Services().ByName("DistributionService").Methods()
	distributionServiceFetchOperationsHandler := connect.NewServerStreamHandler(
		DistributionServiceFetchOperationsProcedure,
		svc.FetchOperations,
		connect.WithSchema(distributionServiceMethods.ByName("FetchOperations")),
		connect.WithHandlerOptions(opts...),
	)
	return "/distribution.v1.DistributionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DistributionServiceFetchOperationsProcedure:
			distributionServiceFetchOperationsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDistributionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDistributionServiceHandler struct{}

func (UnimplementedDistributionServiceHandler) FetchOperations(context.Context, *connect.Request[v1.FetchOperationsRequest], *connect.ServerStream[v1.FetchOperationsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("distribution.v1.DistributionService.FetchOperations is not implemented"))
}
