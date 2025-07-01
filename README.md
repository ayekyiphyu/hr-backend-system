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



### Reference Link


```bash

https://zenn.dev/kurusugawa/articles/golang-env-lib

```

---
