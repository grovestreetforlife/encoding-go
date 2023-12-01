package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	fileJSON, err := os.ReadFile(j.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при открытии JSON-файла: %s", err.Error())
	}

	if err := json.Unmarshal(fileJSON, &j.DockerCompose); err != nil {
		return fmt.Errorf("ошибка при десериализации JSON-файла: %s", err.Error())
	}

	data, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации YAML-файла: %s", err.Error())
	}

	fileYAML, err := os.Create(j.FileOutput)
	if err != nil {
		return fmt.Errorf("ошибка при создании YAML-файла: %s", err.Error())
	}
	defer fileYAML.Close()

	if _, err = fileYAML.Write(data); err != nil {
		return fmt.Errorf("ошибка при записи YAML в файл: %s", err.Error())
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	fileYAML, err := os.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при открытии YAML-файла: %s", err.Error())
	}

	if err := yaml.Unmarshal(fileYAML, &y.DockerCompose); err != nil {
		return fmt.Errorf("ошибка при десериализации YAML-файла: %s", err.Error())
	}

	data, err := json.MarshalIndent(&y.DockerCompose, "", "    ")
	if err != nil {
		return fmt.Errorf("ошибка при сериализации JSON-файла: %s", err.Error())
	}

	fileJSON, err := os.Create(y.FileOutput)
	if err != nil {
		return fmt.Errorf("ошибка при создании JSON-файла: %s", err.Error())
	}
	defer fileJSON.Close()

	if _, err = fileJSON.Write(data); err != nil {
		return fmt.Errorf("ошибка при записи JSON в файл: %s", err.Error())
	}

	return nil
}
