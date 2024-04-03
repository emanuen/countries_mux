package models

type Country struct {
	Name   Name   `json:"name"`
	Region string `json:"region"`
}

type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}
