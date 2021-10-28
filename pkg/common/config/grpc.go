package config

import (
	context "context"
)

// GRPCClient is an implementation of the config store that talks over RPC.
type GRPCClient struct {
	client ConfigStoreClient
}

func (m *GRPCClient) Get(key string) (interface{}, error) {
	resp, err := m.client.Get(context.Background(), &GetRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	return getValue(resp), nil
}

func getValue(resp *GetResponse) interface{} {
	switch resp.GetValueUnion().(type) {
	case *GetResponse_BoolValue:
		return resp.GetBoolValue()
	case *GetResponse_FloatValue:
		return resp.GetFloatValue()
	case *GetResponse_IntValue:
		return resp.GetIntValue()
	default:
		return resp.GetStringValue()
	}
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// We must embed the unimplemented interface
	// for a config store server.
	UnimplementedConfigStoreServer
	// This is the real implementation
	Impl Provider
}

func (m *GRPCServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	v, err := m.Impl.Get(req.Key)
	if err != nil {
		return nil, err
	}
	return &GetResponse{ValueUnion: toValueUnion(v)}, nil
}

func toValueUnion(value interface{}) isGetResponse_ValueUnion {
	switch final := value.(type) {
	case bool:
		return &GetResponse_BoolValue{BoolValue: final}
	case int:
	case int32:
	case int64:
		return &GetResponse_IntValue{IntValue: int64(final)}
	case float32:
	case float64:
		return &GetResponse_FloatValue{FloatValue: float32(final)}
	case string:
		return &GetResponse_StringValue{StringValue: final}
	}
	return nil
}
