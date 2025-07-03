##  Start Golang REST API with Gin

###  Install Go

#### macOS (with Homebrew)

```bash
brew install go
```

#### Windows (WSL / Ubuntu)

```bash
sudo apt update
sudo apt install golang-go
```

---

##  Check Go version

```bash
go version
```

---

##  Create Go Project

```bash
mkdir project_name
cd project_name
go mod init my-rest-api
```

---

##  Install Gin Framework (Recommended)

```bash
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors   # For CORS support
```

---


## Folder structure description
### this is rule of folder structure
```bash

project/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── handlers/
│   ├── auth.go          # Authentication handlers
│   ├── health.go
│   └── users.go
├── middleware/
│   ├── auth.go          # JWT authentication middleware
│   └── cors.go
├── models/
│   ├── auth.go          # Auth request/response models
│   ├── response.go
│   └── users.go         # Renamed from usersmodel.go for consistency
├── routes/
│   └── routes.go
├── storage/
│   ├── interface.go     # New: Storage interface
│   ├── memory.go
│   └── postgres.go      # New: Database implementation (optional)
├── utils/
│   ├── jwt.go           # New: JWT utilities
│   └── password.go      # New: Password hashing utilities
├── .env
├── go.mod
└── go.sum
```

##  Create a Complete REST API

Create a `main.go` file:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()
    r.Use(cors.Default())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8080") // Run on localhost:8080
}
```

---

## ▶ Run the server

```bash
go run main.go
```

Open your browser or use curl/postman to check:

```
http://localhost:8080/ping
```

You should see:

```json
{"message":"pong"}
```

---

##  Done!

You now have:

* Go installed
* Go module initialized
* Gin + CORS installed
* A minimal REST API running with Gin.

---

### Next Steps (Optional)

* Structure your project into `/controllers`, `/services`, `/models`, `/routes`.
* Add CRUD endpoints with Gin.
* Add Docker support for containerization.
* Add unit tests with `testing` and `httptest`.

※ I had development window version


### database most popular
* PostgreSQL
* MySQL
* SQLite
* MariaDB Server



### Folder sturcture description


```bash

* handlers/ - HTTP request handlers
* models/ - Data structures
* storage/ - Data persistence layer
* middleware/ - Reusable middleware
* routes/ - Route definitions

```

```bash

* go mod init your-api  (use your project name)
* go mod tidy (download dependencies)
* go run main.go ( start the server)


## password Auth (JWT token or Google auth)

```bash
## JWT
* go get -u github.com/golang-jwt/jwt


## google auth
*go get golang.org/x/oauth2
*go get golang.org/x/oauth2/google
*go get github.com/gin-gonic/gin
```



# studying reference Page
```

Refernce : https://zenn.dev/o_ga/books/dc6c7b055b65a3/viewer/chapter2

---
