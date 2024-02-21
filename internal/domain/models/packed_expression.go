package models

type StoredExpression struct {
	Original string `json:"orig"`
	RPN      string `json:"rpn"`
	Now      string `json:"now"`
	Queue    []int  `json:"queue"`
}
