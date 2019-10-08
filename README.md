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

### implement Auth

1. Update the schema and ORM model

	在 schema  還有 ORM model 裡面都加上 User，欄位我只填最基本的 username/password，然後記得如果在 production 實作的話，密碼存入 DB 之前要加密！

	`directive @isAuthenticated on FIELD_DEFINITION`

	`createCategory(input: CategoryInput): Category! @isAuthenticated`

	如此一來，如果想要 createCategory，就需要先通過 isAuthenticated 的驗證！

2. Run gqlgen to generate new codes based on the new schema

3. Write a auth middleware to handle incoming request

	```Go
	package auth

	import (
		"context"
		"net/http"

		"github.com/wtlin1228/go-gql-server/internal/gql/resolvers"
	)

	// Middleware is used to handle auth logic
	func Middleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			auth := r.Header.Get("Authorization")
			if auth != "" {
				// Write your fancy token introspection logic here and if valid user then pass appropriate key in header
				// IMPORTANT: DO NOT HANDLE UNAUTHORIZED USER HERE
				ctx = context.WithValue(ctx, resolvers.UserIDCtxKey, auth)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	```

	並且在 resolvers/main.go 裡面加上

	```Go
	type contextKey string

	var (
		UserIDCtxKey = contextKey("userID")
	)

	func NewRootResolvers(orm *orm.ORM) generated.Config {
		c := generated.Config{
			Resolvers: &Resolver{
				ORM: orm, // pass in the ORM instance in the resolvers to be used
			},
		}

		// Schema Directive
		c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			ctxUserID := ctx.Value(UserIDCtxKey)
			if ctxUserID == nil {
				return nil, errors.UnAuthorizedError
			}
			return next(ctx)
		}

		return c
	}
	```

4. Wrap the http handle function with auth middleware

	將 GraphqlHandler 改成這樣

	```Go
	func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {

		h := auth.Middleware(
			handler.GraphQL(gql.NewExecutableSchema(resolvers.NewRootResolvers(orm))),
		)

		return func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		}
	}
	```

	記得加上自己的 login logic，產生你用來做驗證的 token，然後就可以拿這個 token 給 auth middleware 驗證，範例可以看 `/internal/gql/resolvers/users.go`