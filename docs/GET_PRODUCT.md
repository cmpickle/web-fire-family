# API
## Requests
### **GET** - /product/{sku}
## Get Product  
Allows you to search a product by its specific SKU, where the word in brackets is just an int value. Returns a JSON array where the only element is the found object. 

### Example Request
`GET /product/1`
`content-type: application/json`


### Example Response
`200 OK`

```
[
    {
        "productid": 2,
        "productname": "Swing",
        "notificationquantity": 10,
        "color": "test",
        "trimcolor": "test",
        "size": "test",
        "price": 1,
        "dimensions": "test",
        "sku": 1,
        "quantity": 5
    }
]
```