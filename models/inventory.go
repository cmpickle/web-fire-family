package models

// Inventory - The inventory database model
type Inventory struct {
	InventoryID     int    `json:"inventoryid,omitempty"`
	Quantity        int    `json:"quantity"`
	DateLastUpdated string `json:"datelastupdated, omitempty"`
	ProductID       int    `json:"productid,omitempty"`
	Deleted         int    `json:"deleted,omitempty"`
	SKU             int    `json:"sku,omitempty"`
}
