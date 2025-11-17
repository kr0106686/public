package filer

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Node map[string]interface{}

func (f *Filer) ReadFreeYaml(fileName string) map[string]interface{} {
	filePath := path.Join(f.docsPath, fileName)
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var tree Node
	if err := yaml.Unmarshal(yamlFile, &tree); err != nil {
		panic(err)
	}
	fmt.Println("Project structure generated!")

	return tree
}
