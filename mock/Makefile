.PHONY: mock
mock:
	@mockgen -source=GoWeb/internal/service/user.go -package=svcmocks -destination=GoWeb/internal/service/mocks/user.mock.go
	@mockgen -source=GoWeb/internal/repository/user.go -package=repomocks -destination=GoWeb/internal/repository/mocks/user.mock.go
	@mockgen -source=GoWeb/internal/repository/dao/user.go -package=daomocks -destination=GoWeb/internal/repository/dao/mocks/user.mock.go
	@mockgen -source=GoWeb/internal/repository/cache/user.go -package=cachemocks -destination=GoWeb/internal/repository/cache/mocks/user.mock.go
	@go mod tidy