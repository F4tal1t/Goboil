package service

import (
	"github.com/F4tal1t/Goboil/internal/lib/job"
	"github.com/F4tal1t/Goboil/internal/repository"
	"github.com/F4tal1t/Goboil/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}