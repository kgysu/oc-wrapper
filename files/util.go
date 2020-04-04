package files

import "io/ioutil"

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func CreateFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}
