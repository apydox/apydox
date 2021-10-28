package config

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"

	"github.com/apydox/apydox/pkg/common/plugins"
)

// Provider is an interface for a config
// service that you can retrieve secrets
// and configuration from.
type Provider interface {
	Get(key string) (interface{}, error)
}

type providerImpl struct {
	configCache       map[string]interface{}
	pluginClient      *plugin.Client
	pluginConfigStore Provider
}

func NewDefaultProvider() (Provider, error) {
	provider := &providerImpl{}
	err := provider.discoverAndStartPlugins()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *providerImpl) discoverAndStartPlugins() error {
	discoveredPlugins, err := plugins.Discover(plugins.Globs["configstore"])
	if err != nil {
		return err
	}
	if len(discoveredPlugins) > 0 {
		// We can only have a single config store plugin so we'll
		// take the first of the discovery results.
		configStorePlugin := discoveredPlugins[0]
		p.pluginClient = plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig:  plugins.Handshake,
			Plugins:          plugins.PluginMap(),
			Cmd:              exec.Command(configStorePlugin),
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		})
		rpcClient, err := p.pluginClient.Client()
		if err != nil {
			return err
		}
		raw, err := rpcClient.Dispense("configstore")
		if err != nil {
			return err
		}
		p.pluginConfigStore = raw.(Provider)
	}
	return nil
}

func (p *providerImpl) Get(key string) (interface{}, error) {
	return nil, nil
}
