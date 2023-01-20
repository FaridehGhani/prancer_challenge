# pancer online code challenge

## How to run

```bash
go mod tidy 
```

```bash
go run cmd/main.go
```
```bash
curl --request POST \
  --url http://localhost:8080/deliver_point \
  --header 'Content-Type: application/json' \
  --data '{
	"x": 10,
	"y": 20
}'
```