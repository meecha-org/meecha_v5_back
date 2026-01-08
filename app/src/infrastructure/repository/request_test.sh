export DB_HOST=postgres-db
export DB_TEST_HOST=postgres-db
export DB_PORT=5432
export DB_USER=meecha
export DB_PASSWORD=meecha_pass
export DB_NAME=meecha
export DATABASE_URL="host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable TimeZone=Asia/Tokyo"

go test ./infrastructure/repository/request_test.go
