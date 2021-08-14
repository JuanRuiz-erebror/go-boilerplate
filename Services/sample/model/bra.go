package model

type BraMetadata struct {
	Other string `json:"other" example:"bbbb"`
}

type Bra struct {
	Name     string      `json:"name" example:"aaaa"`
	Metadata BraMetadata `json:"metadata" example:"metadataaaa"`
}
