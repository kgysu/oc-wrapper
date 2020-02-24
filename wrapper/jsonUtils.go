package wrapper

import "encoding/json"

func ParseJsonString(jsonContent string, v interface{}) error {
	return json.Unmarshal([]byte(jsonContent), v)
}

func ObjectToJsonString(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
