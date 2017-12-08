# API
## Requests
### **POST** - /product/delete/{sku}
## Delete Product  
Isn't really a delete. Just goes in and toggles a column in the database from 0 to 1 so it is effectively just archived. Nothing fancy here.

### Example Request
`POST /product/delete/1`

### Example Response
`200 OK`
```
{
    "deleted": "true"
}
```
