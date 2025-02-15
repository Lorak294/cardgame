package user

import (
	"context"
	"server/db/hashing"
	"strconv"
	"time"
)

type service struct {
	repository IRepository
	timeout    time.Duration	
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse,error) {
	timeout_ctx, cancel := context.WithTimeout(ctx,s.timeout)
	defer cancel()

	hashedPassword, err := hashing.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r, err := s.repository.CreateUser(timeout_ctx,u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		Id: strconv.Itoa(int(r.Id)),
		Username: r.Username,
		Email: r.Username,
	}

	return res, nil
}

func NewService(repository IRepository) IService {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}
