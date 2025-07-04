// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.34
// source: demo.proto

package pb

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_demo_proto struct{}

func (drpcEncoding_File_demo_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_demo_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_demo_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_demo_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_demo_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCStatsClient interface {
	DRPCConn() drpc.Conn

	ProcessDemo(ctx context.Context, in *ProcessDemoReq) (*ProcessDemoResponse, error)
}

type drpcStatsClient struct {
	cc drpc.Conn
}

func NewDRPCStatsClient(cc drpc.Conn) DRPCStatsClient {
	return &drpcStatsClient{cc}
}

func (c *drpcStatsClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcStatsClient) ProcessDemo(ctx context.Context, in *ProcessDemoReq) (*ProcessDemoResponse, error) {
	out := new(ProcessDemoResponse)
	err := c.cc.Invoke(ctx, "/pb.Stats/ProcessDemo", drpcEncoding_File_demo_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCStatsServer interface {
	ProcessDemo(context.Context, *ProcessDemoReq) (*ProcessDemoResponse, error)
}

type DRPCStatsUnimplementedServer struct{}

func (s *DRPCStatsUnimplementedServer) ProcessDemo(context.Context, *ProcessDemoReq) (*ProcessDemoResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCStatsDescription struct{}

func (DRPCStatsDescription) NumMethods() int { return 1 }

func (DRPCStatsDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/pb.Stats/ProcessDemo", drpcEncoding_File_demo_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCStatsServer).
					ProcessDemo(
						ctx,
						in1.(*ProcessDemoReq),
					)
			}, DRPCStatsServer.ProcessDemo, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterStats(mux drpc.Mux, impl DRPCStatsServer) error {
	return mux.Register(impl, DRPCStatsDescription{})
}

type DRPCStats_ProcessDemoStream interface {
	drpc.Stream
	SendAndClose(*ProcessDemoResponse) error
}

type drpcStats_ProcessDemoStream struct {
	drpc.Stream
}

func (x *drpcStats_ProcessDemoStream) SendAndClose(m *ProcessDemoResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_demo_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
