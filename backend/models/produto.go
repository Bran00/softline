package models
// Produto representa a estrutura de um produto
type Produto struct {
	ID           int     `json:"id"`
	Nome         string  `json:"nome"`
	Descricao    string  `json:"descricao"`
	CodigoBarras string  `json:"codigo_barras"`
	ValorVenda   float64 `json:"valor_venda"`
	PesoBruto    float64 `json:"peso_bruto"`
	PesoLiquido  float64 `json:"peso_liquido"`
}
