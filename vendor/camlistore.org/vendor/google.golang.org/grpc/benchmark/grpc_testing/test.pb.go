// Code generated by protoc-gen-go.
// source: benchmark/grpc_testing/test.proto
// DO NOT EDIT!

/*
Package grpc_testing is a generated protocol buffer package.

It is generated from these files:
	benchmark/grpc_testing/test.proto

It has these top-level messages:
	StatsRequest
	ServerStats
	Payload
	HistogramData
	ClientConfig
	Mark
	ClientArgs
	ClientStats
	ClientStatus
	ServerConfig
	ServerArgs
	ServerStatus
	SimpleRequest
	SimpleResponse
*/
package grpc_testing

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
const _ = proto.ProtoPackageIsVersion1

type PayloadType int32

const (
	// Compressable text format.
	PayloadType_COMPRESSABLE PayloadType = 0
	// Uncompressable binary format.
	PayloadType_UNCOMPRESSABLE PayloadType = 1
	// Randomly chosen from all other formats defined in this enum.
	PayloadType_RANDOM PayloadType = 2
)

var PayloadType_name = map[int32]string{
	0: "COMPRESSABLE",
	1: "UNCOMPRESSABLE",
	2: "RANDOM",
}
var PayloadType_value = map[string]int32{
	"COMPRESSABLE":   0,
	"UNCOMPRESSABLE": 1,
	"RANDOM":         2,
}

func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}
func (PayloadType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClientType int32

const (
	ClientType_SYNCHRONOUS_CLIENT ClientType = 0
	ClientType_ASYNC_CLIENT       ClientType = 1
)

var ClientType_name = map[int32]string{
	0: "SYNCHRONOUS_CLIENT",
	1: "ASYNC_CLIENT",
}
var ClientType_value = map[string]int32{
	"SYNCHRONOUS_CLIENT": 0,
	"ASYNC_CLIENT":       1,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}
func (ClientType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServerType int32

const (
	ServerType_SYNCHRONOUS_SERVER ServerType = 0
	ServerType_ASYNC_SERVER       ServerType = 1
)

var ServerType_name = map[int32]string{
	0: "SYNCHRONOUS_SERVER",
	1: "ASYNC_SERVER",
}
var ServerType_value = map[string]int32{
	"SYNCHRONOUS_SERVER": 0,
	"ASYNC_SERVER":       1,
}

func (x ServerType) String() string {
	return proto.EnumName(ServerType_name, int32(x))
}
func (ServerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RpcType int32

const (
	RpcType_UNARY     RpcType = 0
	RpcType_STREAMING RpcType = 1
)

var RpcType_name = map[int32]string{
	0: "UNARY",
	1: "STREAMING",
}
var RpcType_value = map[string]int32{
	"UNARY":     0,
	"STREAMING": 1,
}

func (x RpcType) String() string {
	return proto.EnumName(RpcType_name, int32(x))
}
func (RpcType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StatsRequest struct {
	// run number
	TestNum int32 `protobuf:"varint,1,opt,name=test_num" json:"test_num,omitempty"`
}

func (m *StatsRequest) Reset()                    { *m = StatsRequest{} }
func (m *StatsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()               {}
func (*StatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerStats struct {
	// wall clock time
	TimeElapsed float64 `protobuf:"fixed64,1,opt,name=time_elapsed" json:"time_elapsed,omitempty"`
	// user time used by the server process and threads
	TimeUser float64 `protobuf:"fixed64,2,opt,name=time_user" json:"time_user,omitempty"`
	// server time used by the server process and all threads
	TimeSystem float64 `protobuf:"fixed64,3,opt,name=time_system" json:"time_system,omitempty"`
}

func (m *ServerStats) Reset()                    { *m = ServerStats{} }
func (m *ServerStats) String() string            { return proto.CompactTextString(m) }
func (*ServerStats) ProtoMessage()               {}
func (*ServerStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Payload struct {
	// The type of data in body.
	Type PayloadType `protobuf:"varint,1,opt,name=type,enum=grpc.testing.PayloadType" json:"type,omitempty"`
	// Primary contents of payload.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type HistogramData struct {
	Bucket       []uint32 `protobuf:"varint,1,rep,name=bucket" json:"bucket,omitempty"`
	MinSeen      float64  `protobuf:"fixed64,2,opt,name=min_seen" json:"min_seen,omitempty"`
	MaxSeen      float64  `protobuf:"fixed64,3,opt,name=max_seen" json:"max_seen,omitempty"`
	Sum          float64  `protobuf:"fixed64,4,opt,name=sum" json:"sum,omitempty"`
	SumOfSquares float64  `protobuf:"fixed64,5,opt,name=sum_of_squares" json:"sum_of_squares,omitempty"`
	Count        float64  `protobuf:"fixed64,6,opt,name=count" json:"count,omitempty"`
}

func (m *HistogramData) Reset()                    { *m = HistogramData{} }
func (m *HistogramData) String() string            { return proto.CompactTextString(m) }
func (*HistogramData) ProtoMessage()               {}
func (*HistogramData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ClientConfig struct {
	ServerTargets             []string   `protobuf:"bytes,1,rep,name=server_targets" json:"server_targets,omitempty"`
	ClientType                ClientType `protobuf:"varint,2,opt,name=client_type,enum=grpc.testing.ClientType" json:"client_type,omitempty"`
	EnableSsl                 bool       `protobuf:"varint,3,opt,name=enable_ssl" json:"enable_ssl,omitempty"`
	OutstandingRpcsPerChannel int32      `protobuf:"varint,4,opt,name=outstanding_rpcs_per_channel" json:"outstanding_rpcs_per_channel,omitempty"`
	ClientChannels            int32      `protobuf:"varint,5,opt,name=client_channels" json:"client_channels,omitempty"`
	PayloadSize               int32      `protobuf:"varint,6,opt,name=payload_size" json:"payload_size,omitempty"`
	// only for async client:
	AsyncClientThreads int32   `protobuf:"varint,7,opt,name=async_client_threads" json:"async_client_threads,omitempty"`
	RpcType            RpcType `protobuf:"varint,8,opt,name=rpc_type,enum=grpc.testing.RpcType" json:"rpc_type,omitempty"`
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Request current stats
type Mark struct {
}

func (m *Mark) Reset()                    { *m = Mark{} }
func (m *Mark) String() string            { return proto.CompactTextString(m) }
func (*Mark) ProtoMessage()               {}
func (*Mark) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ClientArgs struct {
	// Types that are valid to be assigned to Argtype:
	//	*ClientArgs_Setup
	//	*ClientArgs_Mark
	Argtype isClientArgs_Argtype `protobuf_oneof:"argtype"`
}

func (m *ClientArgs) Reset()                    { *m = ClientArgs{} }
func (m *ClientArgs) String() string            { return proto.CompactTextString(m) }
func (*ClientArgs) ProtoMessage()               {}
func (*ClientArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isClientArgs_Argtype interface {
	isClientArgs_Argtype()
}

type ClientArgs_Setup struct {
	Setup *ClientConfig `protobuf:"bytes,1,opt,name=setup,oneof"`
}
type ClientArgs_Mark struct {
	Mark *Mark `protobuf:"bytes,2,opt,name=mark,oneof"`
}

func (*ClientArgs_Setup) isClientArgs_Argtype() {}
func (*ClientArgs_Mark) isClientArgs_Argtype()  {}

func (m *ClientArgs) GetArgtype() isClientArgs_Argtype {
	if m != nil {
		return m.Argtype
	}
	return nil
}

func (m *ClientArgs) GetSetup() *ClientConfig {
	if x, ok := m.GetArgtype().(*ClientArgs_Setup); ok {
		return x.Setup
	}
	return nil
}

func (m *ClientArgs) GetMark() *Mark {
	if x, ok := m.GetArgtype().(*ClientArgs_Mark); ok {
		return x.Mark
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientArgs) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientArgs_OneofMarshaler, _ClientArgs_OneofUnmarshaler, _ClientArgs_OneofSizer, []interface{}{
		(*ClientArgs_Setup)(nil),
		(*ClientArgs_Mark)(nil),
	}
}

func _ClientArgs_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ClientArgs_Setup:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Setup); err != nil {
			return err
		}
	case *ClientArgs_Mark:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mark); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientArgs.Argtype has unexpected type %T", x)
	}
	return nil
}

func _ClientArgs_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientArgs)
	switch tag {
	case 1: // argtype.setup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientConfig)
		err := b.DecodeMessage(msg)
		m.Argtype = &ClientArgs_Setup{msg}
		return true, err
	case 2: // argtype.mark
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mark)
		err := b.DecodeMessage(msg)
		m.Argtype = &ClientArgs_Mark{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientArgs_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ClientArgs_Setup:
		s := proto.Size(x.Setup)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientArgs_Mark:
		s := proto.Size(x.Mark)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ClientStats struct {
	Latencies   *HistogramData `protobuf:"bytes,1,opt,name=latencies" json:"latencies,omitempty"`
	TimeElapsed float64        `protobuf:"fixed64,3,opt,name=time_elapsed" json:"time_elapsed,omitempty"`
	TimeUser    float64        `protobuf:"fixed64,4,opt,name=time_user" json:"time_user,omitempty"`
	TimeSystem  float64        `protobuf:"fixed64,5,opt,name=time_system" json:"time_system,omitempty"`
}

func (m *ClientStats) Reset()                    { *m = ClientStats{} }
func (m *ClientStats) String() string            { return proto.CompactTextString(m) }
func (*ClientStats) ProtoMessage()               {}
func (*ClientStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ClientStats) GetLatencies() *HistogramData {
	if m != nil {
		return m.Latencies
	}
	return nil
}

type ClientStatus struct {
	Stats *ClientStats `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
}

func (m *ClientStatus) Reset()                    { *m = ClientStatus{} }
func (m *ClientStatus) String() string            { return proto.CompactTextString(m) }
func (*ClientStatus) ProtoMessage()               {}
func (*ClientStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ClientStatus) GetStats() *ClientStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type ServerConfig struct {
	ServerType ServerType `protobuf:"varint,1,opt,name=server_type,enum=grpc.testing.ServerType" json:"server_type,omitempty"`
	Threads    int32      `protobuf:"varint,2,opt,name=threads" json:"threads,omitempty"`
	EnableSsl  bool       `protobuf:"varint,3,opt,name=enable_ssl" json:"enable_ssl,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ServerArgs struct {
	// Types that are valid to be assigned to Argtype:
	//	*ServerArgs_Setup
	//	*ServerArgs_Mark
	Argtype isServerArgs_Argtype `protobuf_oneof:"argtype"`
}

func (m *ServerArgs) Reset()                    { *m = ServerArgs{} }
func (m *ServerArgs) String() string            { return proto.CompactTextString(m) }
func (*ServerArgs) ProtoMessage()               {}
func (*ServerArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type isServerArgs_Argtype interface {
	isServerArgs_Argtype()
}

type ServerArgs_Setup struct {
	Setup *ServerConfig `protobuf:"bytes,1,opt,name=setup,oneof"`
}
type ServerArgs_Mark struct {
	Mark *Mark `protobuf:"bytes,2,opt,name=mark,oneof"`
}

func (*ServerArgs_Setup) isServerArgs_Argtype() {}
func (*ServerArgs_Mark) isServerArgs_Argtype()  {}

func (m *ServerArgs) GetArgtype() isServerArgs_Argtype {
	if m != nil {
		return m.Argtype
	}
	return nil
}

func (m *ServerArgs) GetSetup() *ServerConfig {
	if x, ok := m.GetArgtype().(*ServerArgs_Setup); ok {
		return x.Setup
	}
	return nil
}

func (m *ServerArgs) GetMark() *Mark {
	if x, ok := m.GetArgtype().(*ServerArgs_Mark); ok {
		return x.Mark
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerArgs) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerArgs_OneofMarshaler, _ServerArgs_OneofUnmarshaler, _ServerArgs_OneofSizer, []interface{}{
		(*ServerArgs_Setup)(nil),
		(*ServerArgs_Mark)(nil),
	}
}

func _ServerArgs_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ServerArgs_Setup:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Setup); err != nil {
			return err
		}
	case *ServerArgs_Mark:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mark); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerArgs.Argtype has unexpected type %T", x)
	}
	return nil
}

func _ServerArgs_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerArgs)
	switch tag {
	case 1: // argtype.setup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerConfig)
		err := b.DecodeMessage(msg)
		m.Argtype = &ServerArgs_Setup{msg}
		return true, err
	case 2: // argtype.mark
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mark)
		err := b.DecodeMessage(msg)
		m.Argtype = &ServerArgs_Mark{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerArgs_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ServerArgs_Setup:
		s := proto.Size(x.Setup)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerArgs_Mark:
		s := proto.Size(x.Mark)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ServerStatus struct {
	Stats *ServerStats `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
	Port  int32        `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ServerStatus) GetStats() *ServerStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type SimpleRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, server randomly chooses one from other formats.
	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Desired payload size in the response from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	ResponseSize int32 `protobuf:"varint,2,opt,name=response_size" json:"response_size,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
}

func (m *SimpleRequest) Reset()                    { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string            { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()               {}
func (*SimpleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SimpleRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SimpleResponse struct {
	Payload *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SimpleResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*StatsRequest)(nil), "grpc.testing.StatsRequest")
	proto.RegisterType((*ServerStats)(nil), "grpc.testing.ServerStats")
	proto.RegisterType((*Payload)(nil), "grpc.testing.Payload")
	proto.RegisterType((*HistogramData)(nil), "grpc.testing.HistogramData")
	proto.RegisterType((*ClientConfig)(nil), "grpc.testing.ClientConfig")
	proto.RegisterType((*Mark)(nil), "grpc.testing.Mark")
	proto.RegisterType((*ClientArgs)(nil), "grpc.testing.ClientArgs")
	proto.RegisterType((*ClientStats)(nil), "grpc.testing.ClientStats")
	proto.RegisterType((*ClientStatus)(nil), "grpc.testing.ClientStatus")
	proto.RegisterType((*ServerConfig)(nil), "grpc.testing.ServerConfig")
	proto.RegisterType((*ServerArgs)(nil), "grpc.testing.ServerArgs")
	proto.RegisterType((*ServerStatus)(nil), "grpc.testing.ServerStatus")
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
	proto.RegisterEnum("grpc.testing.PayloadType", PayloadType_name, PayloadType_value)
	proto.RegisterEnum("grpc.testing.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("grpc.testing.ServerType", ServerType_name, ServerType_value)
	proto.RegisterEnum("grpc.testing.RpcType", RpcType_name, RpcType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for TestService service

type TestServiceClient interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	StreamingCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) StreamingCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/grpc.testing.TestService/StreamingCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingCallClient{stream}
	return x, nil
}

type TestService_StreamingCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceStreamingCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceStreamingCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	StreamingCall(TestService_StreamingCallServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_StreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingCall(&testServiceStreamingCallServer{stream})
}

type TestService_StreamingCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type testServiceStreamingCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceStreamingCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingCall",
			Handler:       _TestService_StreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

// Client API for Worker service

type WorkerClient interface {
	// Start test with specified workload
	RunTest(ctx context.Context, opts ...grpc.CallOption) (Worker_RunTestClient, error)
	// Start test with specified workload
	RunServer(ctx context.Context, opts ...grpc.CallOption) (Worker_RunServerClient, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) RunTest(ctx context.Context, opts ...grpc.CallOption) (Worker_RunTestClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Worker_serviceDesc.Streams[0], c.cc, "/grpc.testing.Worker/RunTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerRunTestClient{stream}
	return x, nil
}

type Worker_RunTestClient interface {
	Send(*ClientArgs) error
	Recv() (*ClientStatus, error)
	grpc.ClientStream
}

type workerRunTestClient struct {
	grpc.ClientStream
}

func (x *workerRunTestClient) Send(m *ClientArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerRunTestClient) Recv() (*ClientStatus, error) {
	m := new(ClientStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) RunServer(ctx context.Context, opts ...grpc.CallOption) (Worker_RunServerClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Worker_serviceDesc.Streams[1], c.cc, "/grpc.testing.Worker/RunServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerRunServerClient{stream}
	return x, nil
}

type Worker_RunServerClient interface {
	Send(*ServerArgs) error
	Recv() (*ServerStatus, error)
	grpc.ClientStream
}

type workerRunServerClient struct {
	grpc.ClientStream
}

func (x *workerRunServerClient) Send(m *ServerArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerRunServerClient) Recv() (*ServerStatus, error) {
	m := new(ServerStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Worker service

type WorkerServer interface {
	// Start test with specified workload
	RunTest(Worker_RunTestServer) error
	// Start test with specified workload
	RunServer(Worker_RunServerServer) error
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_RunTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).RunTest(&workerRunTestServer{stream})
}

type Worker_RunTestServer interface {
	Send(*ClientStatus) error
	Recv() (*ClientArgs, error)
	grpc.ServerStream
}

type workerRunTestServer struct {
	grpc.ServerStream
}

func (x *workerRunTestServer) Send(m *ClientStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerRunTestServer) Recv() (*ClientArgs, error) {
	m := new(ClientArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Worker_RunServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).RunServer(&workerRunServerServer{stream})
}

type Worker_RunServerServer interface {
	Send(*ServerStatus) error
	Recv() (*ServerArgs, error)
	grpc.ServerStream
}

type workerRunServerServer struct {
	grpc.ServerStream
}

func (x *workerRunServerServer) Send(m *ServerStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerRunServerServer) Recv() (*ServerArgs, error) {
	m := new(ServerArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunTest",
			Handler:       _Worker_RunTest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunServer",
			Handler:       _Worker_RunServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 867 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0x51, 0x8f, 0xdb, 0x44,
	0x10, 0x3e, 0xe7, 0xe2, 0xe4, 0x3c, 0x4e, 0xd2, 0xb0, 0xb4, 0x25, 0x3d, 0xee, 0xa1, 0x18, 0x44,
	0x4f, 0x87, 0x48, 0xab, 0x20, 0x21, 0xde, 0x20, 0x4d, 0x53, 0xae, 0xa2, 0xe7, 0xab, 0xec, 0x3b,
	0x50, 0x9f, 0xac, 0x8d, 0xb3, 0x4d, 0xac, 0x73, 0xd6, 0xae, 0x77, 0x0d, 0x04, 0x5e, 0x78, 0xe1,
	0x1f, 0xf0, 0x17, 0xf8, 0x9f, 0xec, 0x8e, 0x6d, 0xe1, 0x04, 0x83, 0x0e, 0xf5, 0x29, 0xf2, 0xec,
	0xcc, 0xb7, 0xdf, 0x37, 0xf3, 0xcd, 0x06, 0x3e, 0x5a, 0x30, 0x1e, 0xae, 0x37, 0x34, 0xbb, 0x79,
	0xbc, 0xca, 0xd2, 0x30, 0x90, 0x4c, 0xc8, 0x88, 0xaf, 0x1e, 0xeb, 0xdf, 0x71, 0x9a, 0x25, 0x32,
	0x21, 0x3d, 0x7d, 0x30, 0x2e, 0x0f, 0x9c, 0x87, 0xd0, 0xf3, 0x25, 0x95, 0xc2, 0x63, 0x6f, 0x73,
	0x15, 0x22, 0x43, 0x38, 0xd2, 0x47, 0x01, 0xcf, 0x37, 0x23, 0xe3, 0xa1, 0x71, 0x6a, 0x3a, 0xdf,
	0x81, 0xed, 0xb3, 0xec, 0x47, 0x96, 0x61, 0x1e, 0xb9, 0x0b, 0x3d, 0x19, 0x6d, 0x58, 0xc0, 0x62,
	0x9a, 0x0a, 0xb6, 0xc4, 0x24, 0x83, 0xbc, 0x07, 0x16, 0x46, 0x73, 0xc1, 0xb2, 0x51, 0x0b, 0x43,
	0xef, 0x83, 0x8d, 0x21, 0xb1, 0x15, 0x92, 0x6d, 0x46, 0x87, 0x3a, 0xe8, 0x7c, 0x03, 0xdd, 0x57,
	0x74, 0x1b, 0x27, 0x74, 0x49, 0x1e, 0x41, 0x5b, 0x6e, 0x53, 0x86, 0x00, 0x83, 0xc9, 0x83, 0x71,
	0x9d, 0xd6, 0xb8, 0x4c, 0xba, 0x52, 0x09, 0xa4, 0x07, 0xed, 0x45, 0xb2, 0xdc, 0x22, 0x6c, 0xcf,
	0xf9, 0x09, 0xfa, 0xe7, 0x91, 0x90, 0xc9, 0x2a, 0xa3, 0x9b, 0x67, 0x54, 0x52, 0x32, 0x80, 0xce,
	0x22, 0x0f, 0x6f, 0x98, 0x54, 0x48, 0x87, 0xa7, 0x7d, 0xad, 0x60, 0x13, 0xf1, 0x40, 0x30, 0xc6,
	0x4b, 0x26, 0x3a, 0x42, 0x7f, 0x2e, 0x22, 0x48, 0x83, 0xd8, 0x70, 0x28, 0x94, 0xc0, 0x36, 0x7e,
	0xdc, 0x87, 0x81, 0xfa, 0x08, 0x92, 0x37, 0x81, 0x78, 0x9b, 0xd3, 0x8c, 0x89, 0x91, 0x89, 0xf1,
	0x3e, 0x98, 0x61, 0x92, 0x73, 0x39, 0xea, 0x20, 0xf5, 0xdf, 0x5b, 0xd0, 0x9b, 0xc5, 0x11, 0xe3,
	0x72, 0x96, 0xf0, 0x37, 0xd1, 0x0a, 0xeb, 0xb0, 0x31, 0x81, 0xa4, 0xd9, 0x8a, 0x49, 0x81, 0x04,
	0x2c, 0xf2, 0x39, 0xd8, 0x21, 0xe6, 0x05, 0xa8, 0xaf, 0x85, 0xfa, 0x46, 0xbb, 0xfa, 0x0a, 0x20,
	0x94, 0x47, 0x00, 0x18, 0xa7, 0x8b, 0x58, 0x75, 0x4a, 0xc4, 0xc8, 0xef, 0x88, 0x7c, 0x02, 0x27,
	0x49, 0x2e, 0x85, 0xa4, 0x7c, 0xa9, 0xb2, 0x03, 0x55, 0x29, 0x82, 0x54, 0x5d, 0x14, 0xae, 0x29,
	0xe7, 0x2c, 0x46, 0xe2, 0x26, 0xf9, 0x00, 0xee, 0x94, 0x17, 0x95, 0xf1, 0x82, 0xb9, 0xa9, 0x67,
	0x94, 0x16, 0x0d, 0x0c, 0x44, 0xf4, 0x0b, 0x43, 0x01, 0x26, 0x39, 0x81, 0xbb, 0x54, 0x6c, 0x79,
	0x18, 0x54, 0xec, 0xd6, 0x19, 0xa3, 0x4b, 0x31, 0xea, 0xe2, 0xe9, 0x23, 0x38, 0x42, 0xc3, 0x68,
	0xca, 0x47, 0x48, 0xf9, 0xde, 0x2e, 0x65, 0x2f, 0x0d, 0x35, 0x5f, 0xa7, 0x03, 0xed, 0x0b, 0xe5,
	0x2f, 0x67, 0x0d, 0x50, 0xa8, 0x98, 0x66, 0x2b, 0x41, 0x3e, 0x03, 0x53, 0x30, 0x99, 0xa7, 0x38,
	0x4e, 0x7b, 0x72, 0xdc, 0x24, 0xb7, 0xe8, 0xdb, 0xf9, 0x01, 0x71, 0xa0, 0xad, 0x2d, 0x8a, 0xad,
	0xb1, 0x27, 0x64, 0x37, 0x57, 0x83, 0x9f, 0x1f, 0x3c, 0xb5, 0xa0, 0xab, 0xba, 0xaa, 0xe9, 0x38,
	0xbf, 0x82, 0x5d, 0x00, 0x14, 0x0e, 0x1c, 0x83, 0x15, 0x53, 0xa9, 0x7c, 0x1e, 0x31, 0x51, 0x5e,
	0xf7, 0xe1, 0x2e, 0xc4, 0xae, 0x41, 0xf6, 0x1d, 0x7b, 0xf8, 0x4f, 0xc7, 0xb6, 0x9b, 0x1c, 0x8b,
	0x2e, 0x70, 0xbe, 0xaa, 0xa6, 0xae, 0x2f, 0xcf, 0x05, 0x39, 0x55, 0x42, 0x35, 0x8d, 0xf2, 0xe6,
	0x07, 0x4d, 0x42, 0x91, 0xa7, 0xb3, 0x50, 0xab, 0x85, 0xfe, 0x28, 0xfd, 0xa2, 0x7c, 0x51, 0xf9,
	0xe5, 0x6f, 0xdf, 0xef, 0xf9, 0xa2, 0x28, 0x40, 0x5f, 0xdc, 0x81, 0x6e, 0x35, 0xa1, 0x16, 0x4e,
	0xa8, 0xc1, 0x28, 0x7a, 0x08, 0x45, 0xc9, 0x2d, 0x86, 0x50, 0x27, 0xf3, 0xff, 0x87, 0xf0, 0xbc,
	0x52, 0x73, 0xab, 0x3e, 0xd4, 0x5f, 0x0c, 0xb5, 0xbf, 0x69, 0x92, 0xc9, 0x42, 0x85, 0xf3, 0x9b,
	0x01, 0x7d, 0x3f, 0xda, 0xa4, 0x31, 0xab, 0x9e, 0x9c, 0x27, 0xd0, 0x57, 0x4b, 0x97, 0x26, 0x5c,
	0xb0, 0xe0, 0x76, 0x2f, 0xc2, 0xbd, 0x5a, 0x05, 0x1a, 0xbc, 0x68, 0xd0, 0xa7, 0xd0, 0x2d, 0x6d,
	0x8f, 0xdd, 0xb1, 0xf7, 0x1d, 0x5c, 0x42, 0xa8, 0x91, 0x0e, 0x2a, 0x06, 0x05, 0x48, 0xbd, 0xd2,
	0xf8, 0x8f, 0xca, 0xb3, 0xaf, 0xc1, 0xae, 0xf3, 0x18, 0x2a, 0x6f, 0x5c, 0x5e, 0xbc, 0xf2, 0xe6,
	0xbe, 0x3f, 0x7d, 0xfa, 0x72, 0x3e, 0x3c, 0x50, 0x33, 0x1a, 0x5c, 0xbb, 0x3b, 0x31, 0x43, 0x8d,
	0xad, 0xe3, 0x4d, 0xdd, 0x67, 0x97, 0x17, 0xc3, 0xd6, 0xd9, 0x97, 0xd5, 0xd2, 0x60, 0xfd, 0x7d,
	0x20, 0xfe, 0x6b, 0x77, 0x76, 0xee, 0x5d, 0xba, 0x97, 0xd7, 0x7e, 0x30, 0x7b, 0xf9, 0x62, 0xee,
	0x5e, 0x29, 0x14, 0x85, 0x3b, 0xd5, 0x07, 0x55, 0xc4, 0xd0, 0x75, 0x35, 0x6b, 0xec, 0xd5, 0xf9,
	0x73, 0xef, 0xfb, 0xb9, 0x57, 0xaf, 0x2b, 0x23, 0xc6, 0xd9, 0xc7, 0xd0, 0x2d, 0xf7, 0x96, 0x58,
	0x60, 0x5e, 0xbb, 0x53, 0xef, 0xb5, 0xca, 0xeb, 0x83, 0xe5, 0x5f, 0x79, 0xf3, 0xe9, 0xc5, 0x0b,
	0xf7, 0xdb, 0xa1, 0x31, 0xf9, 0xd3, 0x00, 0xfb, 0x4a, 0x29, 0xd5, 0x37, 0x44, 0x21, 0x23, 0xcf,
	0xc1, 0xba, 0xe6, 0x34, 0xdb, 0xce, 0x68, 0x1c, 0x93, 0xbd, 0xd5, 0xda, 0x19, 0xdd, 0xf1, 0x49,
	0xf3, 0x61, 0xd9, 0x55, 0x57, 0x4d, 0x5a, 0x2a, 0x07, 0xab, 0xe7, 0x78, 0xf5, 0x8e, 0x58, 0xa7,
	0xc6, 0x13, 0x63, 0xf2, 0x87, 0x01, 0x9d, 0x1f, 0x92, 0xec, 0x86, 0x65, 0x64, 0xa6, 0x74, 0xe5,
	0x5c, 0x93, 0x26, 0x8d, 0x2f, 0xab, 0x5e, 0x87, 0xe3, 0xe3, 0x7f, 0xdb, 0xcd, 0x5c, 0x68, 0x3c,
	0x32, 0x07, 0x4b, 0x81, 0x14, 0x7d, 0x25, 0x8d, 0x8b, 0xd8, 0x04, 0x53, 0xdf, 0x02, 0x0d, 0xb3,
	0xe8, 0xe0, 0xff, 0xea, 0x17, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xdf, 0x5b, 0x40, 0x7c,
	0x07, 0x00, 0x00,
}
