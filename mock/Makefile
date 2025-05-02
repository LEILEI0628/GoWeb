.PHONY: mock
mock:
	@mockgen -source=../internal/service/user.go -package=svcmocks -destination=../internal/service/mocks/user.mock.go
	@mockgen -source=../internal/repository/user.go -package=repomocks -destination=../internal/repository/mocks/user.mock.go
	@mockgen -source=../internal/repository/dao/user.go -package=daomocks -destination=../internal/repository/dao/mocks/user.mock.go
	@mockgen -source=../internal/repository/cache/user.go -package=cachemocks -destination=../internal/repository/cache/mocks/user.mock.go
	@mockgen -source=../internal/service/article.go -package=svcmocks -destination=../internal/service/mocks/article.mock.go
	@mockgen -source=../internal/repository/article.go -package=repomocks -destination=../internal/repository/mocks/article.mock.go
	@mockgen -source=../internal/repository/dao/article.go -package=daomocks -destination=../internal/repository/dao/mocks/article.mock.go
	@go mod tidy