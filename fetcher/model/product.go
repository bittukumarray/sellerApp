package model

type Url struct {
	URL string `json:"url"`
}

type Product struct {
	ImageUrl string `json:"imageUrl"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price string `json:"price"`
	TotalReviews int `json:"totalReviews"`

}

type ResponseModel struct {
	URL string `json:"url"`
	Product Product `json:"product"`
}
