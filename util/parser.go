package util

import (
	"encoding/json"
	"fmt"
	err "github.com/asifhajiyev/matching-api/error"
	"strconv"
	"strings"
)

func StringToFloat(s string) (float64, *err.Error) {
	if strings.TrimSpace(s) == "" {
		return 0.0, err.ParsingError("argument should not be empty")
	}
	f, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return 0.0, err.ParsingError("argument could not be parsed to float number")
	}
	return f, nil
}

func StringToInt(s string) (int, *err.Error) {
	i, e := strconv.Atoi(s)
	if e != nil {
		message := fmt.Sprintf("%s could not be parsed to int number", s)
		return 0, err.ParsingError(message)
	}
	return i, nil
}

func InterfaceToStruct(from interface{}, to interface{}) *err.Error {
	js, _ := json.Marshal(from)
	e := json.Unmarshal(js, to)

	if e != nil {
		return err.ServerError(e.Error())
	}
	return nil
}

func JsonToStruct(from []byte, to interface{}) *err.Error {
	e := json.Unmarshal(from, to)

	if e != nil {
		return err.ServerError(e.Error())
	}
	return nil
}
