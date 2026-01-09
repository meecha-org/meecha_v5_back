export DB_HOST=mysql-db
export DB_TEST_HOST=mysql-test
export DB_PORT=3306
export DB_USER=meecha
export DB_PASSWORD=meecha_pass
export DB_NAME=meecha

export DATABASE_URL="${DB_USER}:${DB_PASSWORD}@tcp(${DB_TEST_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local"

go test ./infrastructure/repository/request_test.go
