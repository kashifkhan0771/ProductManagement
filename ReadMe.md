<pre>Product Management System v 1.1

Language Used : GO Lang

Note : 
    Feel free to open an issue if you find any.
    Feel free to improve code by new PR's. 

______________________________________

Features : 
    1. Add Product 
    2. Delete Product 
    3. Update/Edit Product 
    4. Get Product 

Database :
    1. MongoDB

API's : 
    - /addProduct  (post request)   - to add a new product(JSON Format)
    - /product{ID} (get request)    - to get a product by id
    - /product{ID} (delete request) - to delete a product
    - /product{ID} (put request)    - to update/edit a product(JSON Format)
    - /listProducts (get request)   - to get all products(read code for further information)

Product Model :
    Product:
        type: "object"
    properties:
        ID:
            type: "string"
        ProductID:
            type: "string"
        Name:
            type: "string"
        Price:
            type: "integer"
        Status:
            type: "string"
        Quantity:
            type: "integer"

Docker : 
    1. cd root_of_project
    2. sudo make test (Command)

Local API Testing : 
    1. cd root_of_project/cmd/product-management-server
    2. go run main.go
    3. USE Postman Or Swagger Editor To Test API's.
</pre>
