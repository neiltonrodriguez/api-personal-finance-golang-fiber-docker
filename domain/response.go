package domain

import (
	"encoding/json"
)

type Response struct {
	Meta  Meta        `json:"meta,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"errors,omitempty"`
}

type ResponseCheck struct {
	Check bool      `json:"check,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type Authorization struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data,omitempty"`
}

type Meta struct {
	Count      int                 `json:"count"`
	Pagination *paginationResponse `json:"pagination,omitempty"`
}

type Check struct {
	Status bool `json:"status"`
}

type LegacyResponse struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func ToJson(v interface{}) string {
	jsonToReturn, e := json.Marshal(v)
	if e != nil {
		return ""
	} else {
		return string(jsonToReturn)
	}
}

func FromJson(v interface{}, jsonString string) error {
	return json.Unmarshal([]byte(jsonString), &v)
}
