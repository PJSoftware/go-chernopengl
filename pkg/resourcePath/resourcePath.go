package resourcePath

import (
	"fmt"
	"os"
)

func locateResource(rType string, file string) (string, error) {
	path := "res/" + rType + "/" + file
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("resource file '%s' not found", path)
	}
	return path, nil
}

func Shader(file string) (string, error)  { return locateResource("shaders", file) }
func Texture(file string) (string, error) { return locateResource("textures", file) }
