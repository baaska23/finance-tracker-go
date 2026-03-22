package subcategories

type SubCategory struct {
	SubCategoryId   int    `json:"sub_category_id"`
	SubCategoryName string `json:"sub_category_name"`
}

type SubCategoryPatch struct {
	SubCategoryId   *int    `json:"sub_category_id"`
	SubCategoryName *string `json:"sub_category_name"`
}
