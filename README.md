
### start golang
#### macOS with Homebrew
brew install go

# window
sudo apt install golang-go


# check go version
go version



## Create Go Project
mkdir "project_name"
cd p project_name
go mod init my-rest-api

## Recommend install Gin Framework

go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors  # For CORS support

## Create a Complete REST API
