package model

type Size struct {
	Lenght int `json:"lenght"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSize(lenght int, width int, height int) *Size {
	return &Size{
		Lenght: lenght,
		Width:  width,
		Height: height,
	}
}
