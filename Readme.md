### Backend The Graduate 

# Step By Step
- go mod download 
- download dependency golang migrate: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- jalankan migrasi: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
- jalankan file: go run main.go 


## Hal yang dibutuhkan
- golang 
- postgreSQL


## ENV VARIABLE

DB_PASSWORD=
DB_USERNAME=
DB_HOST=
DB_PORT=
DB_NAME=
JWT_SECRET_KEY=
LDAP_DOMAIN= <<OPTIONAL FOR NOW>>
LDAP_BASEDN=dc= <<OPTIONAL FOR NOW>>