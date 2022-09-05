# GO-CRM-REST-API

### Features

All REST APIs (GET, POST, PATCH, DELETE)
Start server inside the project root directory
`go run main.go`

The application handles the following 5 operations for customers in the "database":
> 
    Getting a single customer through a /customers/{id} path
    Getting all customers through a the /customers path
    Creating a customer through a /customers path
    Updating a customer through a /customers/{id} path
    Deleting a customer through a /customers/{id} path

### API

Get All customers

```Json
URL: GET localhost:8080/customers
```

Response:

```Json
[
    {
    "id": "0",
    "name": "John Doe",
    "role": "Software Developer 53",
    "email": "mockEmail@anymail.com",
    "phone": 5550199,
    "contacted": true
    },
    {
    "id": "3",
    "name": "John Doe",
    "role": "Software Developer 46",
    "email": "mockEmail@anymail.com",
    "phone": 5550199,
    "contacted": true
    },
    {
    "id": "4",
    "name": "Example Name",
    "role": "Example Role",
    "email": "Example Email",
    "phone": 5550199,
    "contacted": true
    },
    {
    "id": "5",
    "name": "Example Name",
    "role": "Example Role",
    "email": "Example Email",
    "phone": 5550199,
    "contacted": true
    },
    {
    "id": "6",
    "name": "Example Name",
    "role": "Example Role",
    "email": "Example Email",
    "phone": 5550199,
    "contacted": true
    }
]
```

Get Customer by ID
```Json
URL: GET localhost:8080/customers/0
```

Response:
```Json
 {
    "id": "0",
    "name": "John Doe",
    "role": "Software Developer 53",
    "email": "mockEmail@anymail.com",
    "phone": 5550199,
    "contacted": true
    },
```

Add new Customer
```Json
URL: POST localhost:8080/customers
```

Body
```Json
{
"name": "John Doe",
"role": "Software Developer 53",
"email": "mockEmail@anymail.com",
"phone": 5550199,
"contacted": true
}
```
Response:
If the customer array length is `0` then `id:1`
```Json

{
"id":"1",
"name": "John Doe",
"role": "Software Developer 53",
"email": "mockEmail@anymail.com",
"phone": 5550199,
"contacted": true
}
```
Update a Customer
```Json
URL: PUT localhost:8080/customers/1
```
Body
```Json
{
"id": "1",
"name": "Gyanendra Verma"
}
```
Response:
```Json
{
"id":"1",
"name": "Gyanendra Verma",
"role": "Software Developer 53",
"email": "mockEmail@anymail.com",
"phone": 5550199,
"contacted": true
}
```
Delete a Customer
```Json
URL: DELETE localhost:8080/customer/1
```

```
Response:

```Json
The event with ID 1 has been deleted successfully
```
