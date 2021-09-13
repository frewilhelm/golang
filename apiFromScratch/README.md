# apiFromScratch

The code in this repository is based on or inspired by
https://www.youtube.com/watch?v=2v11Ym6Ct9Q

## Usage

```sh
# Run on host
ADMIN_PASSWORD=<secret> go run server.go

# POST (example)
curl localhost:8080/progLangs -X POST -d "$(jq '.progLang[0]' progLang.json)" -H "Content-Type: application/json"

# GET all
curl localhost:8080/progLangs | jq

# GET id
curl localhost:8080/progLangs/<id> | jq

# GET random
curl localhost:8080/progLangs/random -L | jq

# Access /admin
curl localhost:8080/admin -u admin:<secret>
```
