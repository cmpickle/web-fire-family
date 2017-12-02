# API
## Requests
### **PUT** - /product/update/{sku}
## Update Product  
Updates a product. It is just as particular as the create route, and is also identical. It requires all fields to be overwritten and does not load the old default values.  You can do this however on the front end by doing a get route hit to populate the form fields if you want.

### Example Request
`PUT /product/update/1`
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
`202 Accepted`
