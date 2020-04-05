package files

import (
	"fmt"
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
	if strings.Contains(result, "$") {
		fmt.Println("WARN: content contains not calculated placeholders!")
	}
	return result
}
