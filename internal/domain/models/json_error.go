package models

type JsonErr struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}
