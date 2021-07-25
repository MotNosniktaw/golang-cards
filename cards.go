package cards

type Card struct {
	suit  string
	name  string
	value int
}

type Deck struct {
	cards []card
}
