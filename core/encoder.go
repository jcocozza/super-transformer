package core

import "encoding/json"

func EncodeJson(m map[string]any) ([]byte, error) {
	return json.Marshal(m)
}
