package cardgame

type Card struct {
	Id     int64  `json:"id" db:"id"`
	Text   string `json:"text" db:"text"`
	Type   int8   `json:"type" db:"type"`
	DeckId int64  `json:"deckId" db:"deckId"`
}

type Deck struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	AuthorId    int64  `json:"authorId" db:"authorId"`
}