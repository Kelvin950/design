
## User Authentication API

---

### POST `/createUser`

Create a new user account.

#### Request

#### Body


```json
{
  "email": "kdldss@gmail.com",
  "fullname": "elvin cdc",
  "user_type": "normal",
  "user_role": "user",
  "avatar": "dfvd",
  "username": "xdsd",
  "password": "testing12"
}
```

POST /loginUser
Content-Type: application/json

```json 
{
  "email": "kdldss@gmail.com",
  "password": "testing12"
}
```

## Products API

---

### POST `/products`

Create a new product.

#### Request

#### Body
```json
{
  "name": "Cool Hoodie 3",
  "color": [16, 17],
  "sizes": [13, 14],
  "quantity": 100,
  "min_price": 29.99,
  "max_price": 49.99,
  "brand_id": 16,
  "brand": {
    "id": 16,
    "name": "Nike"
  },
  "type_id": 10
}
```

GET  /products


## Orders API

---

### POST `/order`

Create a new order for a user.

#### Request


#### Body
```json
[
  {
    "product_id": 3,
    "quantity": 2,
    "price": 21.00
  }
]

```

GET /order?userId=int


## Wishlist API

### GET `/wishlist?userId=1`

Fetch wishlist items for a specific user.

#### Request

#### Query Parameters

| Name     | Type | Required | Description         |
|----------|------|----------|---------------------|
| userId   | int  | Yes      | ID of the user      |

#### Success Response

- **Code:** 200 OK
- **Content:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "product_id": 3,
    "product": {
      "id": 3,
      "name": "Cool T-Shirt",
      "color": [1, 2],
      "sizes": [1, 2, 3],
      "quantity": 100,
      "min_price": 20.0,
      "max_price": 25.0
    }
  }
]
```

POST  /wishlist
Content-Type: application/json


```json

{
  "user_id": 1,
  "product_id": 3
}

```