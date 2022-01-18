package cell

var cell bool

func InitCell() bool {
	cell = false
	return cell
}

func SetCellValue(value bool) bool {
	cell = value //endret verdi
	return cell  // retunerer den endrete verdien
}

func GetCellValue() bool {
	return cell //sjekker verdien på variabelen i pakke-avgrensning
}
