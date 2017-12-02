# API
## Requests
### **GET** - /inventory/{sku}
## Get Inventory
Returns a JSON array that contains all rows associated to the specified SKU. 

### Example Request
`GET /inventory/3`

### Example Response
`200 OK`

```
[
    {
        "inventoryid": 4,
        "quantity": 9,
        "datelastupdated": "2017-11-21 05:58:08",
        "deleted": 4,
        "sku": 3
    }
]
```