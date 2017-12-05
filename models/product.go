package models

//Matches our product table
type Product struct {
	ProductID            int     `json:"productid,omitempty"`
	ProductName          string  `json:"productname,omitempty"`
	NotificationQuantity int     `json:"notificationquantity, omitempty"`
	Color                string  `json:"color,omitempty"`
	TrimColor            string  `json:"trimcolor,omitempty"`
	Size                 string  `json:"size,omitempty"`
	Price                float32 `json:"price,omitempty"`
	Dimensions           string  `json:"dimensions,omitempty"`
	SKU                  int     `json:"sku,omitempty"`
	Deleted              int     `json:"deleted,omitempty"`
	Quantity        	 int     `json:"quantity"`
}
