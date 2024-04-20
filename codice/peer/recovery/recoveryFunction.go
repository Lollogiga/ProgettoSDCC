package recovery

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var RecoveryString string
var RecoveryId = -1

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func GetId(filePath string) int {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error during read file")
	}

	var yamlData struct {
		Id int `yaml:"Id"`
	}

	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		log.Fatalf("Error during unmarshaling")
	}
	return yamlData.Id
}

func SaveId(id int32) {

	type Data struct {
		Id int32 `yaml:"Id"`
	}

	data := Data{
		Id: id,
	}

	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatalf("failed to marshal data to YAML: %v", err)
	}

	if err := os.WriteFile("recovery/Id.yaml", yamlData, 0644); err != nil {
		log.Fatalf("failed to write YAML to file: %v", err)
	}

	fmt.Println("YAML file generated successfully.")
}
