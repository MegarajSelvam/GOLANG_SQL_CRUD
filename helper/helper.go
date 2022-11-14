package helper

import (
	"GO_SQL_CRUD/types"
	"encoding/json"
	"log"
	"strconv"
)

func StringToInt(s string) (int, any) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, types.CreateErrorMessage(err.Error())
	}
	return i, nil
}

func ErrorToBytes(error types.Error) []byte {
	result, err := json.Marshal(error)
	if err != nil {
		log.Println(err)
	}
	return result
}

func AnyToErrorBytes(result any) []byte {
	resultBytes, err := json.Marshal(result.(types.Error))
	if err != nil {
		log.Println(err)
	}
	return resultBytes
}

func AnyToEmployeeBytes(result any) []byte {
	resultBytes, err := json.Marshal(result.(types.Employee))
	if err != nil {
		log.Println(err)
	}
	return resultBytes
}

func AnyToEmployeeListBytes(result any) []byte {
	resultBytes, err := json.Marshal(result.([]types.Employee))
	if err != nil {
		log.Println(err)
	}
	return resultBytes
}
