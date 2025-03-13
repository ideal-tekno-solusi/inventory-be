package entity

import database "app/database/main"

type InventoryResponse struct {
	TotalData   int                                `json:"totalData"`
	CurrentPage int                                `json:"currentPage"`
	TotalPage   int                                `json:"totalPage"`
	Items       *[]database.FetchInventoryItemsRow `json:"items"`
}
