package models

type Expression struct {
	Expression      string `json:"expression"`
	AdditionTime    int    `json:"addition_time"`
	SubtractionTime int    `json:"subtraction_time"`
	MultiplyingTime int    `json:"multiplying_time"`
	DivisionTime    int    `json:"division_time"`
}

type ProcessedExpression struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}
