package models

type Users struct {
	Uid    string `json:"uid,omitempty"`
	Resume string `json:"resume"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type IDs struct {
	Id string `json:"id"`
}
