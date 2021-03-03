package jsonutil

import "encoding/json"

func ToJson(o interface{}) string {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(jsonBytes)
	}
}

func ToJsonPretty(o interface{}) string {
	jsonBytes, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return ""
	} else {
		return string(jsonBytes)
	}
}
