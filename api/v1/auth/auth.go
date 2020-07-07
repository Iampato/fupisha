package auth

import (
	"github.com/nairobi-gophers/fupisha/internal/config"
	"github.com/nairobi-gophers/fupisha/internal/store"
)

const (
	//version of api provided by the server
	apiVersion = "v1"
)

//Resource defines dependencies for auth handlers.
type Resource struct {
	Store  store.Store
	Config *config.Config
}

//NewResource returns a configured authentication resource.
func NewResource(store store.Store, cfg *config.Config) *Resource {
	return &Resource{
		Store:  store,
		Config: cfg,
	}
}