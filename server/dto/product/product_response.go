package productdto

type ResponseProduct struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Image     string `json:"image"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Style     string `json:"style"`
	Material  string `json:"material"`
	Color     string `json:"color" `
	Condition string `json:"condition"`
}
