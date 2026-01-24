package model

type Size struct {
	lenght int
	width  int
	height int
}

func NewSize(lenght int, width int, height int) *Size {
	return &Size{
		lenght: lenght,
		width:  width,
		height: height,
	}
}
