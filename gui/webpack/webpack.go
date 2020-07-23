package webpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Webpack is a webpack integration
type Webpack struct {
	Manifest Manifest
}

// Manifest reflects the structure of asset-manifest.json
type Manifest struct {
	Files map[string]string `json:"files"`
	Entrypoints Entrypoints
}

type Entrypoints []string

func New(buildPath string) (*Webpack, error) {
	webpack := &Webpack{}
	assetsManifestPath := path.Join(buildPath, "asset-manifest.json")

	if _, err := os.Stat(assetsManifestPath); os.IsNotExist(err) {
		return webpack, nil
	}

	content, err := ioutil.ReadFile(assetsManifestPath)

	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file %s: %w", assetsManifestPath, err)
	}

	if err = json.Unmarshal(content, &webpack.Manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest file %s: %w", assetsManifestPath, err)
	}

	return webpack, nil
}

func (e Entrypoints) Scripts() Entrypoints {
	var scripts Entrypoints

	for _, f := range e {
		if strings.HasSuffix(f,".js") {
			scripts = append(scripts, f)
		}
	}

	return scripts
}

func (e Entrypoints) Styles() Entrypoints {
	var styles Entrypoints

	for _, f := range e {
		if strings.HasSuffix(f,".css") {
			styles = append(styles, f)
		}
	}

	return styles
}

