package budgets

type Budget struct {
	Id            int     `json:"id"`
	SubCategoryId string  `json:"sub_category_id"`
	Amount        float64 `json:"amount"`
	StartMonth    string  `json:"start_month"`
	Years         int     `json:"years"`
}
