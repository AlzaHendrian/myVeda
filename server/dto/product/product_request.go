package productdto

type CreateProduct struct {
	Title     string `json:"title" form:"title" validate:"required"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
	Image     string `json:"image" form:"image" validate:"required"`
	Price     int    `json:"price" form:"price" validate:"required"`
	Size      string `json:"size" form:"size" validate:"required"`
	Style     string `json:"style" form:"style" validate:"required"`
	Material  string `json:"material" form:"material" validate:"required"`
	Color     string `json:"color" form:"color" validate:"required"`
	Condition string `json:"condition" form:"condition" validate:"required"`
}

type UpdateProduct struct {
	Title     string `json:"title" form:"title"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
	Image     string `json:"image" form:"image"`
	Price     int    `json:"price" form:"price"`
	Size      string `json:"size" form:"size"`
	Style     string `json:"style" form:"style"`
	Material  string `json:"material" form:"material"`
	Color     string `json:"color" form:"color"`
	Condition string `json:"condition" form:"condition"`
}
