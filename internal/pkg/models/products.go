package models

type Product struct {
	Product_name string `json:"name" gorm:"product_name"`
	Article_id   string `json:"art_id" gorm:"article_id"`
	Amount_of    string `json:"amount_of" gorm:"amount_of"`
}

type ProductsRequest struct {
	Products []struct {
		Name            string `json:"name"`
		ContainArticles []struct {
			ArtID    string `json:"art_id"`
			AmountOf string `json:"amount_of"`
		} `json:"contain_articles"`
	} `json:"products"`
}
