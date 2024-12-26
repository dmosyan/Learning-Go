module github.com/dmosyan/Learning-Go/apis/banking-auth

go 1.23.2

require (
	github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib v0.0.0-20241221214953-bbb1ec3b9c82
	github.com/golang-jwt/jwt v3.2.2+incompatible
)

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.1
)

require filippo.io/edwards25519 v1.1.0 // indirect

require (
	github.com/jmoiron/sqlx v1.4.0
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)
