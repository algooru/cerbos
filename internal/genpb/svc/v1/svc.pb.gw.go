// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: svc/v1/svc.proto

/*
Package svcv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package svcv1

import (
	"context"
	"io"
	"net/http"

	"github.com/cerbos/cerbos/internal/genpb/request/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_CerbosService_CheckResourceSet_0(ctx context.Context, marshaler runtime.Marshaler, client CerbosServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.CheckResourceSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CheckResourceSet(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CerbosService_CheckResourceSet_0(ctx context.Context, marshaler runtime.Marshaler, server CerbosServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.CheckResourceSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CheckResourceSet(ctx, &protoReq)
	return msg, metadata, err

}

func request_CerbosService_CheckResourceBatch_0(ctx context.Context, marshaler runtime.Marshaler, client CerbosServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.CheckResourceBatchRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CheckResourceBatch(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CerbosService_CheckResourceBatch_0(ctx context.Context, marshaler runtime.Marshaler, server CerbosServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.CheckResourceBatchRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CheckResourceBatch(ctx, &protoReq)
	return msg, metadata, err

}

func request_CerbosAdminService_AddOrUpdatePolicy_0(ctx context.Context, marshaler runtime.Marshaler, client CerbosAdminServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.AddOrUpdatePolicyRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.AddOrUpdatePolicy(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CerbosAdminService_AddOrUpdatePolicy_0(ctx context.Context, marshaler runtime.Marshaler, server CerbosAdminServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.AddOrUpdatePolicyRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.AddOrUpdatePolicy(ctx, &protoReq)
	return msg, metadata, err

}

func request_CerbosAdminService_AddOrUpdatePolicy_1(ctx context.Context, marshaler runtime.Marshaler, client CerbosAdminServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.AddOrUpdatePolicyRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.AddOrUpdatePolicy(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CerbosAdminService_AddOrUpdatePolicy_1(ctx context.Context, marshaler runtime.Marshaler, server CerbosAdminServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.AddOrUpdatePolicyRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.AddOrUpdatePolicy(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_CerbosAdminService_ListAuditLogEntries_0 = &utilities.DoubleArray{Encoding: map[string]int{"kind": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_CerbosAdminService_ListAuditLogEntries_0(ctx context.Context, marshaler runtime.Marshaler, client CerbosAdminServiceClient, req *http.Request, pathParams map[string]string) (CerbosAdminService_ListAuditLogEntriesClient, runtime.ServerMetadata, error) {
	var protoReq requestv1.ListAuditLogEntriesRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["kind"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "kind")
	}

	e, err = runtime.Enum(val, requestv1.ListAuditLogEntriesRequest_Kind_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "kind", err)
	}

	protoReq.Kind = requestv1.ListAuditLogEntriesRequest_Kind(e)

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_CerbosAdminService_ListAuditLogEntries_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.ListAuditLogEntries(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_CerbosPlaygroundService_PlaygroundValidate_0(ctx context.Context, marshaler runtime.Marshaler, client CerbosPlaygroundServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.PlaygroundValidateRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PlaygroundValidate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CerbosPlaygroundService_PlaygroundValidate_0(ctx context.Context, marshaler runtime.Marshaler, server CerbosPlaygroundServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.PlaygroundValidateRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.PlaygroundValidate(ctx, &protoReq)
	return msg, metadata, err

}

func request_CerbosPlaygroundService_PlaygroundEvaluate_0(ctx context.Context, marshaler runtime.Marshaler, client CerbosPlaygroundServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.PlaygroundEvaluateRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PlaygroundEvaluate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CerbosPlaygroundService_PlaygroundEvaluate_0(ctx context.Context, marshaler runtime.Marshaler, server CerbosPlaygroundServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq requestv1.PlaygroundEvaluateRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.PlaygroundEvaluate(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCerbosServiceHandlerServer registers the http handlers for service CerbosService to "mux".
// UnaryRPC     :call CerbosServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCerbosServiceHandlerFromEndpoint instead.
func RegisterCerbosServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CerbosServiceServer) error {

	mux.Handle("POST", pattern_CerbosService_CheckResourceSet_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/svc.v1.CerbosService/CheckResourceSet")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CerbosService_CheckResourceSet_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosService_CheckResourceSet_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CerbosService_CheckResourceBatch_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/svc.v1.CerbosService/CheckResourceBatch")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CerbosService_CheckResourceBatch_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosService_CheckResourceBatch_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCerbosAdminServiceHandlerServer registers the http handlers for service CerbosAdminService to "mux".
// UnaryRPC     :call CerbosAdminServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCerbosAdminServiceHandlerFromEndpoint instead.
func RegisterCerbosAdminServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CerbosAdminServiceServer) error {

	mux.Handle("POST", pattern_CerbosAdminService_AddOrUpdatePolicy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/svc.v1.CerbosAdminService/AddOrUpdatePolicy")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CerbosAdminService_AddOrUpdatePolicy_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosAdminService_AddOrUpdatePolicy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_CerbosAdminService_AddOrUpdatePolicy_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/svc.v1.CerbosAdminService/AddOrUpdatePolicy")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CerbosAdminService_AddOrUpdatePolicy_1(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosAdminService_AddOrUpdatePolicy_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CerbosAdminService_ListAuditLogEntries_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterCerbosPlaygroundServiceHandlerServer registers the http handlers for service CerbosPlaygroundService to "mux".
// UnaryRPC     :call CerbosPlaygroundServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCerbosPlaygroundServiceHandlerFromEndpoint instead.
func RegisterCerbosPlaygroundServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CerbosPlaygroundServiceServer) error {

	mux.Handle("POST", pattern_CerbosPlaygroundService_PlaygroundValidate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/svc.v1.CerbosPlaygroundService/PlaygroundValidate")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CerbosPlaygroundService_PlaygroundValidate_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosPlaygroundService_PlaygroundValidate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CerbosPlaygroundService_PlaygroundEvaluate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/svc.v1.CerbosPlaygroundService/PlaygroundEvaluate")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CerbosPlaygroundService_PlaygroundEvaluate_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosPlaygroundService_PlaygroundEvaluate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCerbosServiceHandlerFromEndpoint is same as RegisterCerbosServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCerbosServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCerbosServiceHandler(ctx, mux, conn)
}

// RegisterCerbosServiceHandler registers the http handlers for service CerbosService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCerbosServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCerbosServiceHandlerClient(ctx, mux, NewCerbosServiceClient(conn))
}

// RegisterCerbosServiceHandlerClient registers the http handlers for service CerbosService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CerbosServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CerbosServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CerbosServiceClient" to call the correct interceptors.
func RegisterCerbosServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CerbosServiceClient) error {

	mux.Handle("POST", pattern_CerbosService_CheckResourceSet_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosService/CheckResourceSet")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosService_CheckResourceSet_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosService_CheckResourceSet_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CerbosService_CheckResourceBatch_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosService/CheckResourceBatch")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosService_CheckResourceBatch_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosService_CheckResourceBatch_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CerbosService_CheckResourceSet_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api", "check"}, ""))

	pattern_CerbosService_CheckResourceBatch_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "x", "check_resource_batch"}, ""))
)

var (
	forward_CerbosService_CheckResourceSet_0 = runtime.ForwardResponseMessage

	forward_CerbosService_CheckResourceBatch_0 = runtime.ForwardResponseMessage
)

// RegisterCerbosAdminServiceHandlerFromEndpoint is same as RegisterCerbosAdminServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCerbosAdminServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCerbosAdminServiceHandler(ctx, mux, conn)
}

// RegisterCerbosAdminServiceHandler registers the http handlers for service CerbosAdminService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCerbosAdminServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCerbosAdminServiceHandlerClient(ctx, mux, NewCerbosAdminServiceClient(conn))
}

// RegisterCerbosAdminServiceHandlerClient registers the http handlers for service CerbosAdminService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CerbosAdminServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CerbosAdminServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CerbosAdminServiceClient" to call the correct interceptors.
func RegisterCerbosAdminServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CerbosAdminServiceClient) error {

	mux.Handle("POST", pattern_CerbosAdminService_AddOrUpdatePolicy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosAdminService/AddOrUpdatePolicy")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosAdminService_AddOrUpdatePolicy_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosAdminService_AddOrUpdatePolicy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_CerbosAdminService_AddOrUpdatePolicy_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosAdminService/AddOrUpdatePolicy")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosAdminService_AddOrUpdatePolicy_1(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosAdminService_AddOrUpdatePolicy_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CerbosAdminService_ListAuditLogEntries_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosAdminService/ListAuditLogEntries")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosAdminService_ListAuditLogEntries_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosAdminService_ListAuditLogEntries_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CerbosAdminService_AddOrUpdatePolicy_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"admin", "policy"}, ""))

	pattern_CerbosAdminService_AddOrUpdatePolicy_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"admin", "policy"}, ""))

	pattern_CerbosAdminService_ListAuditLogEntries_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"admin", "auditlog", "list", "kind"}, ""))
)

var (
	forward_CerbosAdminService_AddOrUpdatePolicy_0 = runtime.ForwardResponseMessage

	forward_CerbosAdminService_AddOrUpdatePolicy_1 = runtime.ForwardResponseMessage

	forward_CerbosAdminService_ListAuditLogEntries_0 = runtime.ForwardResponseStream
)

// RegisterCerbosPlaygroundServiceHandlerFromEndpoint is same as RegisterCerbosPlaygroundServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCerbosPlaygroundServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCerbosPlaygroundServiceHandler(ctx, mux, conn)
}

// RegisterCerbosPlaygroundServiceHandler registers the http handlers for service CerbosPlaygroundService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCerbosPlaygroundServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCerbosPlaygroundServiceHandlerClient(ctx, mux, NewCerbosPlaygroundServiceClient(conn))
}

// RegisterCerbosPlaygroundServiceHandlerClient registers the http handlers for service CerbosPlaygroundService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CerbosPlaygroundServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CerbosPlaygroundServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CerbosPlaygroundServiceClient" to call the correct interceptors.
func RegisterCerbosPlaygroundServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CerbosPlaygroundServiceClient) error {

	mux.Handle("POST", pattern_CerbosPlaygroundService_PlaygroundValidate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosPlaygroundService/PlaygroundValidate")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosPlaygroundService_PlaygroundValidate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosPlaygroundService_PlaygroundValidate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CerbosPlaygroundService_PlaygroundEvaluate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/svc.v1.CerbosPlaygroundService/PlaygroundEvaluate")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CerbosPlaygroundService_PlaygroundEvaluate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CerbosPlaygroundService_PlaygroundEvaluate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CerbosPlaygroundService_PlaygroundValidate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "playground", "validate"}, ""))

	pattern_CerbosPlaygroundService_PlaygroundEvaluate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "playground", "evaluate"}, ""))
)

var (
	forward_CerbosPlaygroundService_PlaygroundValidate_0 = runtime.ForwardResponseMessage

	forward_CerbosPlaygroundService_PlaygroundEvaluate_0 = runtime.ForwardResponseMessage
)
