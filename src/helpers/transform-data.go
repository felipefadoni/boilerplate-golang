package helpers

import "encoding/json"

func TransformData(dto interface{}, entity interface{}) {
	jsonM, _ := json.Marshal(dto)
	json.Unmarshal(jsonM, &entity)
}
