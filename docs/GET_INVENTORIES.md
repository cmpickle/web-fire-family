# API
## Requests
### **GET** - /inventories
## Get Inventories 
Returns a JSON array of all inventories in the DB not flagged as deleted. 

### Example Request
`GET /inventories`

### Example Response
`200 OK`

```
[
    {
        "inventoryid": 4,
        "quantity": 9,
        "datelastupdated": "2017-11-21 05:58:08",
        "productid": 4,
        "sku": 3
    },
    {
        "inventoryid": 5,
        "quantity": 5,
        "datelastupdated": "2017-12-01 00:02:22",
        "productid": 2,
        "sku": 1
    }
]
```