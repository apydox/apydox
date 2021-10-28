package config

import (
	context "context"

	"google.golang.org/grpc"

	"github.com/apydox/apydox/pkg/common/plugins"
	"github.com/hashicorp/go-plugin"
)

func init() {
	plugins.RegisterPluginMapping("configstore", &ConfigStorePlugin{})
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type ConfigStorePlugin struct {
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Provider
}

func (p *ConfigStorePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterConfigStoreServer(s, &GRPCServer{
		Impl: p.Impl,
	})
	return nil
}

func (p *ConfigStorePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: NewConfigStoreClient(c),
	}, nil
}
