# Go Programming Language
Go Programming Playlist https://youtube.com/playlist?list=PLyZTXfAT27ib7T9Eg3qhvDE5rgvjQk4OL

### Website
* Golang https://golang.org
* Package https://pkg.go.dev
* Go standard library https://pkg.go.dev/std
* Language Specification https://golang.org/ref/spec
* Awesome Go https://github.com/avelino/awesome-go
* Uber Style Guide https://github.com/uber-go/guide
* Go modules cheat sheet https://encore.dev/guide/go.mod
* VSCode Settings https://github.com/golang/vscode-go/blob/master/docs/settings.md
* SQL database drivers https://github.com/golang/go/wiki/SQLDrivers
* Redis https://redis.io
* K6 https://k6.io
* Influx https://influxdata.com
* Grafana https://grafana.com
* Learn Go with Tests https://quii.gitbook.io
* Kafka https://kafka.apache.org
* gRPC https://grpc.io
* Protocol Buffers Release https://github.com/protocolbuffers/protobuf/releases
* Protocol Buffers Doctument https://developers.google.com/protocol-buffers/docs/proto3
* evans gRPC Client https://github.com/ktr0731/evans
* Protocol Buffer Library https://github.com/protocolbuffers/protobuf/tree/main/src/google/protobuf

### Go on macOS
```sh
curl -OL https://golang.org/dl/go1.17.darwin-amd64.tar.gz
tar xfz go1.17.darwin-amd64.tar.gz
sudo mv go /usr/local
sudo ln /usr/local/go/bin/go /usr/local/bin/go
```

### Go on Windows with Windows Package Manager
* [Windows Terminal](https://www.microsoft.com/store/productId/9N0DX20HK701)
* [Windows Package Manager Insiders Program](https://forms.microsoft.com/Pages/ResponsePage.aspx?id=v4j5cvGGr0GRqy180BHbR-NSOqDz219PqoOqk5qxQEZUMVVCT1IwVEpLSklZS0dDRFZEUjZUOU9ZWi4u)
* [App Installer (winget)](https://www.microsoft.com/store/productId/9NBLGGH4NNS1)

### Visual Studio Code and Extension 
* VSCode https://code.visualstudio.com
* Go https://marketplace.visualstudio.com/items?itemName=golang.Go
* Error Lens https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens
* Docker https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker
* MySQL https://marketplace.visualstudio.com/items?itemName=formulahendry.vscode-mysql
* K6 snippets https://marketplace.visualstudio.com/items?itemName=mJacobson.snippets-for-k6
* vscode-proto3 https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3

### Go Package
* database/sql https://pkg.go.dev/database/sql
* jmoiron/sqlx https://pkg.go.dev/github.com/jmoiron/sqlx
* denisenkom/go-mssqldb https://pkg.go.dev/github.com/denisenkom/go-mssqldb
* go-sql-driver/mysql https://pkg.go.dev/github.com/go-sql-driver/mysql
* gorilla/mux https://pkg.go.dev/github.com/gorilla/mux
* go<span></span>.uber.org/zap https://pkg.go.dev/go.uber.org/zap
* spf13/viper https://pkg.go.dev/github.com/spf13/viper
* gofiber/fiber/v2 https://pkg.go.dev/github.com/gofiber/fiber/v2
* gofiber/jwt/v2 https://pkg.go.dev/github.com/gofiber/jwt/v2
* gofiber/adaptor/v2 https://pkg.go.dev/github.com/gofiber/adaptor/v2
* mock/gomock https://pkg.go.dev/github.com/golang/mock/gomock
* net/http https://pkg.go.dev/net/http
* net/http/httptest https://pkg.go.dev/net/http/httptest
* crypto/bcrypt https://pkg.go.dev/golang.org/x/crypto/bcrypt
* dgrijalva/jwt-go https://pkg.go.dev/github.com/dgrijalva/jwt-go
* gorm<span></span>.io/gorm https://pkg.go.dev/gorm.io/gorm
* gorm<span></span>.io/driver/mysql https://pkg.go.dev/gorm.io/driver/mysql
* go-redis/redis/v8 https://pkg.go.dev/github.com/go-redis/redis/v8
* afex/hystrix-go/hystrix https://pkg.go.dev/github.com/afex/hystrix-go/hystrix
* stretchr/testify https://pkg.go.dev/github.com/stretchr/testify
* godoc https://pkg.go.dev/golang.org/x/tools/cmd/godoc
* Shopify/sarama https://pkg.go.dev/github.com/Shopify/sarama
* google/uuid https://pkg.go.dev/github.com/google/uuid
* grpc https://pkg.go.dev/google.golang.org/grpc
* protobuf https://pkg.go.dev/google.golang.org/protobuf
* protoc-gen-go https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go
* protoc-gen-go-grpc https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc

#### Docker Image
* redis https://hub.docker.com/_/redis
* loadimpact/k6 https://hub.docker.com/r/loadimpact/k6
* influxdb https://hub.docker.com/_/influxdb
* mariadb https://hub.docker.com/_/mariadb
* grafana/grafana https://hub.docker.com/r/grafana/grafana
* mlabouardy/hystrix-dashboard https://hub.docker.com/r/mlabouardy/hystrix-dashboard
* zookeeper https://hub.docker.com/_/zookeeper
* kafka https://hub.docker.com/r/bitnami/kafka
* vscode-proto3 https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3

#### Redis Configuration
``` 
bind 0.0.0.0
appendonly yes
SAVE ""
```

#### Unit Test VS Code Configuration
```json
"go.coverOnSave": true,
"go.coverOnSingleTest": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
}
```

### Go gRPC for macOS
1) Install Protobuf on macOS
```
brew install protobuf
```
2) Install Evans gRPC client for macOS
```
brew tap ktr0731/evans
brew install evans
```
3) Install vscode-proto3 for VSCode https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3

4) Go get package in project
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
5) Install gRPC tool in project
```
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```


### Follow me
* **Page:** [https://fb.com/CodeBangkok​](https://fb.com/CodeBangkok​)
* **Group:** [https://fb.com/groups/msdevth​](https://fb.com/groups/msdevth​)
* **Blog:** [https://dev.to/codebangkok](https://dev.to/codebangkok)
* **YouTube:** [https://youtube.com/CodeBangkok](https://youtube.com/CodeBangkok)