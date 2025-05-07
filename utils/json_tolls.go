package utils

import "encoding/json"

// map转json
func ToJsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// 泛型
func StructToJson(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// json转map
func StringToJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

// json转IntList
func ToIntList(data string) []int {
	var tmp = make([]int, 0)
	json.Unmarshal([]byte(data), &tmp)
	return tmp
}
