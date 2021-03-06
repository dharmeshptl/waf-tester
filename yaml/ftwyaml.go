// Package yaml parses HTTP tests written in YAML format.
package yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// ParseFiles parses all YAML files in a directory.
func ParseFiles(path string) []Yaml {
	var yamls []Yaml

	walkFunc := func(path string, fi os.FileInfo, err error) error {
		if !fi.IsDir() && isYaml(path) {
			yaml, err := ParseFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warn parsing %s: %s\n", path, err)
				return err
			}
			yamls = append(yamls, yaml)
		}
		return nil
	}

	filepath.Walk(path, walkFunc)

	return yamls
}

// ParseFile parses a single YAML file.
func ParseFile(filename string) (Yaml, error) {
	var yamlConfig Yaml

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return yamlConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return yamlConfig, err
	}

	for i := range yamlConfig.Tests {
		yamlConfig.Tests[i].File = filename
	}

	return yamlConfig, nil
}

func isYaml(path string) bool {
	return filepath.Ext(path) == ".yaml" ||
		filepath.Ext(path) == ".yml" ||
		filepath.Ext(path) == ".YAML" ||
		filepath.Ext(path) == ".YML"
}
