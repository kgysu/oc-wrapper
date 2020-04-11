package files

import (
	"io/ioutil"
	"strings"
)

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func CreateFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func ReplaceEnvs(content string, envs map[string]string) string {
	result := content
	for key, value := range envs {
		result = strings.ReplaceAll(result, "${"+key+"}", value)
	}
	return result
}
