package repository

import (
	"github.com/google/wire"
	"go-blog/internal/app/repository/greet"
	"go-blog/internal/app/repository/user"
)

var ProviderSet = wire.NewSet(
	wire.NewSet(wire.Bind(new(user.RepositoryInterface), new(*user.Repository)), user.NewRepository),
	wire.NewSet(wire.Bind(new(greet.RepositoryInterface), new(*greet.Repository)), greet.NewRepository),
)
