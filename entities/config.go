package entities

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YAMLStruct struct {
	PackageName struct {
		Description string `yaml:"description"`
		Directories []struct {
			Path string `yaml:"path"`
			Perm uint32 `yaml:"perm"`
		} `yaml:"directories"`
	} `yaml:"project"`
}

func NewConfig() *YAMLStruct {
	return &YAMLStruct{}
}

func (yml *YAMLStruct) FromYAML(file string) (*YAMLStruct, error) {

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return yml, err
	}

	if err := yaml.Unmarshal([]byte(yamlFile), &yml); err != nil {
		return yml, err
	}
	return yml, nil
}
