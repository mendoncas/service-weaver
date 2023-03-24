package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/mendoncas/service-weaver/DbClient",
		Iface: reflect.TypeOf((*DbClient)(nil)).Elem(),
		New:   func() any { return &dbClient{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return dbClient_local_stub{impl: impl.(DbClient), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return dbClient_client_stub{stub: stub, startConnectionMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/service-weaver/DbClient", Method: "StartConnection"}), insertMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/service-weaver/DbClient", Method: "Insert"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return dbClient_server_stub{impl: impl.(DbClient), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/mendoncas/service-weaver/Reverser",
		Iface: reflect.TypeOf((*Reverser)(nil)).Elem(),
		New:   func() any { return &reverser{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return reverser_local_stub{impl: impl.(Reverser), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return reverser_client_stub{stub: stub, reverseMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/service-weaver/Reverser", Method: "Reverse"}), reverseControllerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/service-weaver/Reverser", Method: "ReverseController"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return reverser_server_stub{impl: impl.(Reverser), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type dbClient_local_stub struct {
	impl   DbClient
	tracer trace.Tracer
}

func (s dbClient_local_stub) StartConnection(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.DbClient.StartConnection", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.StartConnection(ctx)
}

func (s dbClient_local_stub) Insert(ctx context.Context, a0 string, a1 string, a2 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.DbClient.Insert", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Insert(ctx, a0, a1, a2)
}

type reverser_local_stub struct {
	impl   Reverser
	tracer trace.Tracer
}

func (s reverser_local_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Reverse(ctx, a0)
}

func (s reverser_local_stub) ReverseController(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Reverser.ReverseController", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ReverseController(ctx)
}

// Client stub implementations.

type dbClient_client_stub struct {
	stub                   codegen.Stub
	startConnectionMetrics *codegen.MethodMetrics
	insertMetrics          *codegen.MethodMetrics
}

func (s dbClient_client_stub) StartConnection(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.startConnectionMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.DbClient.StartConnection", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.startConnectionMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.startConnectionMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.startConnectionMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	if err != nil {
		return
	}
	s.startConnectionMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s dbClient_client_stub) Insert(ctx context.Context, a0 string, a1 string, a2 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.insertMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.DbClient.Insert", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.insertMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.insertMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	size += (4 + len(a2))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	enc.String(a2)
	var shardKey uint64

	// Call the remote method.
	s.insertMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.insertMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

type reverser_client_stub struct {
	stub                     codegen.Stub
	reverseMetrics           *codegen.MethodMetrics
	reverseControllerMetrics *codegen.MethodMetrics
}

func (s reverser_client_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.reverseMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.reverseMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.reverseMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.reverseMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.reverseMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

func (s reverser_client_stub) ReverseController(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.reverseControllerMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Reverser.ReverseController", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.reverseControllerMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.reverseControllerMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.reverseControllerMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	if err != nil {
		return
	}
	s.reverseControllerMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Server stub implementations.

type dbClient_server_stub struct {
	impl    DbClient
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s dbClient_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "StartConnection":
		return s.startConnection
	case "Insert":
		return s.insert
	default:
		return nil
	}
}

func (s dbClient_server_stub) startConnection(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.StartConnection(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s dbClient_server_stub) insert(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()
	var a2 string
	a2 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Insert(ctx, a0, a1, a2)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type reverser_server_stub struct {
	impl    Reverser
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s reverser_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Reverse":
		return s.reverse
	case "ReverseController":
		return s.reverseController
	default:
		return nil
	}
}

func (s reverser_server_stub) reverse(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Reverse(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s reverser_server_stub) reverseController(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.ReverseController(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}
