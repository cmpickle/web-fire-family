# API
## Requests
### **POST** - /inventory/update/{sku}/{quantity}
## Update Inventory
Far less picky than its product cousins. No input json. Changes the quantity field of all inventory rows associated to the given SKU to the given quantity.

### Example Request
`POST /inventory/update/3/20`

### Example Response
`200 OK`
