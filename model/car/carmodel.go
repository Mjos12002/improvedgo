package model

type CarModel struct {
	Type        string
	NumberPlate string
}

func (c CarModel) CheckType(t string) bool {
	return len(t) > 0
}
