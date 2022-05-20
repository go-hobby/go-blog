package file

import (
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"go-blog/pkg/path"
	"path/filepath"
)

type Config struct {
	Path string
}

// New casin file adapter
func New(config *Config) *fileadapter.FilteredAdapter {
	dir := path.GetInstance().Path
	return fileadapter.NewFilteredAdapter(filepath.Join(dir, config.Path))
}
