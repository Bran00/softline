package models

type Cliente struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Fantasia  string `json:"fantasia"`
	Documento string `json:"documento"`
	Endereco  string `json:"endereco"`
}