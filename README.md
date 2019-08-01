# Orders API  

## Dependencies  
* GO   
* Mux  

# Resources  
https://golang.org/doc/install  
https://tutorialedge.net/golang/creating-restful-api-with-golang/  
https://github.com/gorilla/mux  

## Commands  
go build  
./golangapi  
OR  
go build && ./golangapi  
OR  
go run main.go  

## API  

**Definition**  
`GET /orders`  
localhost:8081/api/orders  
**Response**  
- `200 OK` on success  
```json
[
  {
    "id": 1,
    "quantity": 2,
    "productName": "pen"
  },
  {
    "id": 2,
    "quantity": 3,
    "productName": "tape"
  },
  {
    "id": 3,
    "quantity": 1,
    "productName": "glue"
  }
]
```  

**Definition**  
`GET /orders{id}`  
localhost:8081/api/orders/1  
**Response**  
- `200 OK` on success  
- `404 Not Found` if order doesn't exist  
```json
{
  "id": 1,
  "quantity": 2,
  "productName": "pen"
}
```
  
**Definition**  
- `201 OK` on success  
- `409 Conflict` when id already exists  
`POST /orders`  
localhost:8081/api/orders  
**Body**  
```json 
{
  "quantity": 2,
  "productName": "pen"
}
```  
**Response**  
```json
{
  "id": 1,
  "quantity": 2,
  "productName": "pen"
}
```
  
**Definition**  
`DELETE /orders{id}`  
localhost:8081/api/orders/1  
**Response**  
- `204 No Content` on success  
- `404 Not Found` if order doesn't exist  
```json
{
  "id": 1,
  "quantity": 2,
  "productName": "pen"
}
```  