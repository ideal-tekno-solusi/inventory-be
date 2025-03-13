package entity

import database "app/database/main"

type CategoryResponse struct {
	TotalData   int                          `json:"totalData"`
	CurrentPage int                          `json:"currentPage"`
	TotalPage   int                          `json:"totalPage"`
	Categories  *[]database.FetchCategoryRow `json:"categories"`
}
