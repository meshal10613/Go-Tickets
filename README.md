### Load Testing a Protected POST Endpoint

Example using `hey` to send 500 requests concurrently while authenticating with a JWT stored in a cookie.

```bash
hey -n 500 -c 500 \
  -m POST \
  -H "Content-Type: application/json" \
  -H "Cookie: token=YOUR_JWT_TOKEN" \
  -d '{"event_id":1,"quantity":50}' \
  http://localhost:8080/api/orders
```

#### Parameters

- `-n 500` → Total number of requests
- `-c 500` → Number of concurrent workers
- `-m POST` → HTTP method
- `-H "Content-Type: application/json"` → JSON request body
- `-H "Cookie: token=YOUR_JWT_TOKEN"` → JWT authentication via cookie
- `-d` → Request payload

#### Sample Request Body

```json
{
  "event_id": 1,
  "quantity": 50
}
```


### Build go project 

```
go build -o bin/go-tickets ./cmd

bin/go-tickets -> bin: which folder; go-tickets: file name;
./cmd -> that folder where the main.go file
```