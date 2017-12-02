# web-fire-family
Web App


On the front end side of things that we show the user, we might not want to show the productid and inventoryid stuff in 
the fields, everything works off of sku anyway.

/product - GET.
returns a JSON array of all products in the DB not flagged as deleted. 
Fields: 
     - productid, 
     - productname, 
     - notificationquantity, 
     - color, 
     - trimcolor, 
     - size, 
     - price, 
     - dimensions, 
     - sku, 
     - quantity. Quantity is only a returned field when it isn't 0.


/product/{sku} - GET. 
allows you to search a product by its specific SKU, where the word in brackets is just an int value. 
Returns a JSON array where the only element is the found object. EX: /product/3. 
Fields: 
    - productid, 
    - productname, 
    - notificationquantity, 
    - color, 
    - trimcolor, 
    - size, 
    - price, 
    - dimensions, 
    - sku, 
    - quantity. Quantity is only a returned field when it isn't 0.


/product/create - POST. 
Creates a product, is very particular about the fields coming it, must be JSON and have the following 
fields: 
    - "productname": string value, 
    - "notificationquantity": int value, 
    - "color": string value, 
    - "trimcolor": string value, 
    - "size": string value, 
    - "price": decimal/float value, 
    - "dimensions": string value, 
    - "sku": int


/product/update/{sku} - PUT. 
Updates a product. It is just as particular as the create route, and is also identical. 
It requires all fields to be overwritten and does not load the old default values.  
You can do this however on the front end by doing a get route hit to populate the form fields if you want.


/product/delete/[sku} - DELETE. 
Isn't really a delete. Just goes in and toggles a column in the database from 0 to 1 so it is effectively just archived. 
Nothing fancy here.


/inventories - GET. 
returns a JSON array of all inventories in the DB not flagged as deleted. 
Fields: 
    - inventoryid, 
    - quantity, 
    - datelastupdated, 
    - productid,  
    - sku.


/inventory/{sku} - GET. 
returns a JSON array that contains all rows associated to the specified SKU. 
Fields: 
    - inventoryid, 
    - quantity, 
    - datelastupdated, 
    - productid,  
    - sku.


/inventory/update/{sku}/{quantity} - PUT. 
Far less picky than its product cousins. No input json. 
Changes the quantity field of all inventory rows associated to the given SKU to the given quantity.


/inventory/increment/{sku} - PUT. 
Increases the quantity column by one for all rows associated to that SKU. Designed for use with scanner. (Hopefully)


/inventory/decrement/{sku} - PUT. 
Decreases the quantity column by one for all rows associated to that SKU. Designed for use with scanner. (Hopefully)