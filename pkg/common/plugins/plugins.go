package plugins

import (
	"github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion: 1,
	// The magic cookie values should NEVER be changed.
	MagicCookieKey:   "APYDOX_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "2c875e90-3878-4566-99f4-ab4f4e9fc322",
}

// The map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	// "configstore": &config.ConfigStorePlugin{},
}

// PluginMap provides the current version of the plugin
// map that spans all types of plugins in apydox.
func PluginMap() map[string]plugin.Plugin {
	return pluginMap
}

// RegisterPluginMapping allows packages to register
// their own plugin mappings during initialisation.
// This most only be used during package initialisation.
func RegisterPluginMapping(name string, stub plugin.Plugin) {
	pluginMap[name] = stub
}
