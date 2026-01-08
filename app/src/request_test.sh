echo "Running repository tests..."
./infrastructure/repository/request_test.sh

echo "Running application tests..."
go test ./application/usecase/request_interactor_test.go

echo "Running domain tests..."
go test ./domain/request_entity_test.go
