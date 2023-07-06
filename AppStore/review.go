package appstore


type Review struct {
	ReviewText string  `json:"review"`
	Rating     float64 `json:"rating"`
	Date       string  `json:"date"`
	Title      string  `json:"title"`
	IsEdited   bool  `json:"isEdited"`
	UserName   string  `json:"userName"`

}



