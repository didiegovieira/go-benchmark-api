package di

import (
	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api/middleware"
	"github.com/google/wire"
)

var apiMiddlewaresSet = wire.NewSet(
	wire.Struct(new(middleware.Cors), "*"),
)
