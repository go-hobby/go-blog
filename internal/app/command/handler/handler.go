package handler

import (
	"github.com/google/wire"
	"go-blog/internal/app/command/handler/greet"
)

var ProviderSet = wire.NewSet(greet.NewHandler)
