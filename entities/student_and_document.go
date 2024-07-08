package entities

type DocumentAndStudent struct {
	Data     Student  `json:"data"`
	Document Document `json:"document"`
}
