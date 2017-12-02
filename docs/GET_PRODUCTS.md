# API
## Requests
### **GET** - /product
## Get Products 
Returns a JSON array of all products in the DB not flagged as deleted. Fields: productid, productname, notificationquantity, color, trimcolor, size, price, dimensions, sku, quantity. Quantity is only a returned field when it isn't 0.

### Example Request
`GET /product`


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
    },
    {
        "productid": 3,
        "productname": "Chain",
        "notificationquantity": 15,
        "color": "test2",
        "trimcolor": "test2",
        "size": "test2",
        "price": 15.99,
        "dimensions": "test2",
        "sku": 2,
        "quantity": 5
    }
]
```