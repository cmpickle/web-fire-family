# API
## Requests
### **POST** - /inventory/increment/{sku}
## Increment Inventory
Increases the quantity column by one for all rows associated to that SKU. Designed for use with scanner.

### Example Request
`POST /inventory/increment/3`

### Example Response
`200 OK`
