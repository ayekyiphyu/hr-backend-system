# Define the end points of this API


```bash

| Methods           | Endpoints         | Function                                      | Authentication    |
| :---------------- | :---------------- |  :----------------                            | :---------------- |
| GET               |   /v1/health      | Server Health Check                           | -                 |
| GET               |   /v1/health_db   | DB Health Check                               | -                 |
| GET               |   /v1/users       | Getting a list of members                     | 〇                |
| POST              |   /v1/users/:id   | Edit member information                       | 〇                |
| GET               |   /v1/users/:id   | Obtaining detailed information for each member| 〇                |
| DELETE            |   /v1/users/:id   | Deletion of member information                | 〇                |



```

### step by step swag in GO
# how to install swag

```bash
* go install github.com/swaggo/swag/cmd/swag@latest

* go get -u github.com/swaggo/gin-swagger

* go get -u github.com/swaggo/files

```

# how to add bash in swag
```bash
 export PATH=$PATH:$(go env GOPATH)/bin

 * out will be docs.go
 * swagger.json
 * swagger.yaml

※plese check main.go file inside setting.

```

```bash
*   go run main.go
*   http://localhost:8080/swagger/index.html
```