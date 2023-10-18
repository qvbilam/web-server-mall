package validate

type GoodsList struct {
	CategoryId *string `form:"category_id" json:"category_id" binding:"omitempty,numeric"`
	Page       *string `form:"page" json:"page" binding:"omitempty,numeric"`
	PerPage    *string `form:"per_page" json:"per_page" binding:"omitempty,numeric"`
}

type GoodsSell struct {
	Count int64 `form:"count" json:"count" binding:"required,numeric"`
}
