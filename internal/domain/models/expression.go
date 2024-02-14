package models

type Expression struct {
	Expression         string `json:"expression"`
	AdditionTime       int    `json:"addition_time"`
	SubtractionTime    int    `json:"subtraction_time"`
	MultiplicationTime int    `json:"multiplication_time"`
	DivisionTime       int    `json:"division_time"`
}
