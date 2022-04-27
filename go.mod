module github.com/torwald-sergesson/app-a/v2

go 1.17

require (
	github.com/torwald-sergesson/app-a/pkg/client/v2 v2.0.0
	github.com/torwald-sergesson/app-a/pkg/dto/v2 v2.1.3
)

replace (
	github.com/torwald-sergesson/app-a/pkg/client/v2 => ./pkg/client
	github.com/torwald-sergesson/app-a/pkg/dto/v2 => ./pkg/dto
)
