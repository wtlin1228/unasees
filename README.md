# go-gql-server


使用 Golang 建立一個 GraphQL API Server

### How to run the server
`$ ./scripts/run.sh`

### Setup go mod
`$ go mod init github.com/wtlin1228/go-graphql-server`

### Install libraries
`$ go get -u github.com/gin-gonic/gin`

`$ go get -u github.com/99designs/gqlgen`

`$ go get -u github.com/jinzhu/gorm`

`$ go get gopkg.in/gormigrate.v1`

`go get github.com/gofrs/uuid`

### How to modify schema
1. 修改 `internal/gql/schemas`
2. 修改 `internal/orm/models`
3. 修改 `gqlgen.yml`
4. `./scripts/gqlgen.sh`
5. 修改 GQLGen 產生的 Resolvers

### How to create a postgreSQL DB for this server 
1. `$ psql` 進到 postgreSQL
2. `$ postgre=# CREATE DATABASE yourdbname;`
3. `$ postgre=# CREATE USER youruser WITH ENCRYPTED PASSWORD 'yourpass';`
4. `$ postgre=# GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;`

### How to setup ENV
可以參考 `.env.example`

```shell
# Web framework config
GIN_MODE=debug
GQL_SERVER_HOST=localhost
GQL_SERVER_PORT=7000

# GQLGen config
GQL_SERVER_GRAPHQL_PATH=/graphql
GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED=true
GQL_SERVER_GRAPHQL_PLAYGROUND_PATH=/

# GORM config
GORM_AUTOMIGRATE=true
GORM_SEED_DB=true
GORM_LOGMODE=true
GORM_DIALECT=postgres
GORM_CONNECTION_DSN=postgres://username:password@localhost:5432/gqltest?sslmode=disable
```

### How to replace the model's ID with uuid

寫一個 BaseModel 給所有的 Models 使用，取代 gorm 的 BaseModel

```golang
package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
	// ID should use uuid_generate_v4() for the pk's
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index"`
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}

```
