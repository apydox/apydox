package plugins

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-plugin"
)

// Globs provides a mapping of plugin type
// to the glob used to discover plugins of each type.
var Globs = map[string]string{
	"configstore": "apydox-config-store-*",
}

func Discover(glob string) ([]string, error) {
	directories, err := getDirectories()
	if err != nil {
		return nil, err
	}

	plugins := []string{}
	for _, directory := range directories {
		discovered, err := plugin.Discover(glob, directory)
		if err != nil {
			return nil, err
		}
		plugins = append(plugins, discovered...)
	}

	return plugins, nil
}

func getDirectories() ([]string, error) {
	// For now we return a single location where plugins
	// are expected to be stored.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	pluginsPathForOS := filepath.FromSlash(".apydox/plugins")
	pluginsPath := fmt.Sprintf("%s%c%s", homeDir, filepath.Separator, pluginsPathForOS)
	return []string{pluginsPath}, nil
}
