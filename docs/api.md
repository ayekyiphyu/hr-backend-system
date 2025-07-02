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


# how to install swag in GO

```bash
* go install github.com/swaggo/swag/cmd/swag@latest

* go get -u github.com/swaggo/gin-swagger

* go get -u github.com/swaggo/files

```

# how to add bash in swag
```bash
 export PATH=$PATH:$(go env GOPATH)/bin
```
