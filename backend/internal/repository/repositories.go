package repository

import "github.com/F4tal1t/Goboil/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}