package cards

type card struct {
	suit  string
	name  string
	value int
}

type deck struct {
	cards []card
}
