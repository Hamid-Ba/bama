package common

import "encoding/json"

func TypeConverter[T any](data any) (T, error) {
	var result T
	json_data, err := json.Marshal(&data)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(json_data, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
