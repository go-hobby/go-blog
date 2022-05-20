package model

import (
	"github.com/casbin/casbin/v2/model"
	"go-blog/pkg/path"
	"path/filepath"
)

type Config struct {
	Path string
}

func New(config *Config) (model.Model, error) {
	dir := path.GetInstance().Path
	return model.NewModelFromFile(filepath.Join(dir, config.Path))
}
