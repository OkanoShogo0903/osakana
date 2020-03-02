package model

type AquariumModel struct {
	Text string `json:"text"`
}

func New() *AquariumModel {
	return &AquariumModel{}
}
