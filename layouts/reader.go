package layouts

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Is this way as there was trouble doing it right
type YamlLayoutOuter struct {
	YamlLayout []struct {
		Name          string           `yaml:"name"`
		MinGridWidth  int              `yaml:"minGridWidth"`
		MinGridHeight int              `yaml:"minGridHeight"`
		RowSpans      []int            `yaml:"rowSpans"`
		ColumnSpans   []int            `yaml:"columnSpans"`
		Items         map[string][]int `yaml:"items"`
	} `yaml:"yamlLayout"`
}

// Parse yaml config of layouts
func NewYamlLayoutOuter(path string) *YamlLayoutOuter {
	// Make sure file exists
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse into above schema
	config := YamlLayoutOuter{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}
