package user

import (
	"context"
	"os"
	"server/db/hashing"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

type CardgameJWTClaims struct {
	Id 			string 	`json:"id"`
	Username 	string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(ctx context.Context, req *LoginUserRequest)  (*LoginUserResponse, error) {
	timeout_ctx, cancel := context.WithTimeout(ctx,s.timeout)
	defer cancel()

	u, err := s.repository.GetUserByEmail(timeout_ctx,req.Email)
	if err != nil {
		return nil, err
	}

	err = hashing.CheckPassword(req.Password,u.Password)
	if err != nil {
		return nil, err
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CardgameJWTClaims{
		Id:	strconv.Itoa(int(u.Id)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*24)),
		},
	})

	// get the signing key
	signingKey := os.Getenv("JWT_SIGN_KEY")
	sign_str, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	return &LoginUserResponse{
		Id: strconv.Itoa(int(u.Id)),
		Username: u.Username,
		AccessToken: sign_str,
	}, nil
}

func NewService(repository IRepository) IService {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}
