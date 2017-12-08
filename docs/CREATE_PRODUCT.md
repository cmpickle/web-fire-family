# API
## Requests
### **POST** - /product/create
## Create Product  
Creates a product, is very particular about the fields coming it, must be JSON and have ALL of the fields.

### Example Request
`POST /product/create`
`content-type: application/json`
```
{
    "productname": "Swing",
    "notificationquantity": 10,
    "color": "test",
    "trimcolor": "test",
    "size": "test",
    "price": 5.99,
    "dimensions": "test",
    "sku": 1
}

```

### Example Response
`200 OK`

```
{
    "ProductId": 15
}
```