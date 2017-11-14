package models

// Inventory - The inventory database model
type Inventory struct {
	InventoryID     int    `json:"inventoryid,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	DateLastUpdated string `json:"datelastupdated, omitempty"`
	ProductID       int    `json:"productid,omitempty"`
	Deleted         int    `json:"deleted,omitempty"`
}
