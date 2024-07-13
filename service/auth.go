package service

import (
	"context"
	"database/sql"
	"log/slog"
	"user_service/api/auth"
	pb "user_service/genproto/authentication"
	pbu "user_service/genproto/user"
	// l "user_service/pkg/logger"
	"user_service/storage/postgres"

	"github.com/pkg/errors"
)

type AuthService struct {
	pb.UnimplementedAuthenticationServer
	Repo *postgres.UserRepo
	Log  *slog.Logger
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		Repo: postgres.NewUserRepository(db),
		// Log:  l.NewLogger(),
	}
}

func (S *AuthService) Register(ctx context.Context, req *pb.UserDetails) (*pb.ID, error) {
	S.Log.Info("Register service is working")
	res, err := S.Repo.Register(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "Error on register user.")
		S.Log.Error(err.Error())
	}
	S.Log.Info("Register service completed succesfully")
	return res, err
}

func (S *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	S.Log.Info("Login service is working")
	checker, err := S.Repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		err = errors.Wrap(err, "Error while geting user.")
		S.Log.Error(err.Error())
		return nil, err
	}

	if req.Password != checker.Password {
		return nil, errors.New("username or password error")
	}

	res := &pb.LoginResponse{
		Access: &pb.AccessToken{
			Id:       checker.Id,
			Username: checker.Username,
			Email:    checker.Email,
		},
		Refresh: &pb.RefreshToken{
			Userid: checker.Id,
		},
	}
	err = auth.GeneratedRefreshJWTToken(res)
	S.Log.Info("refresh token is generated")
	if err != nil {
		S.Log.Error("Error while create refresh token")
		return nil, err
	}
	err = auth.GeneratedAccessJWTToken(res)
	S.Log.Info("acces token is generated")
	if err != nil {
		S.Log.Error("Error while create access token")
		return nil, err
	}
	err = S.Repo.StoreRefreshToken(ctx, res)
	if err != nil {
		S.Log.Error("Error while inserting refresh token")
		return nil, err
	}
	S.Log.Info("Login service completed succesfully")
	return res, nil
}

func (S *AuthService) CheckRefreshToken(ctx context.Context, req *pb.CheckRefreshTokenRequest) (*pb.CheckRefreshTokenResponse, error) {
	S.Log.Info("CheckRefreshToken service is working")
	
	_, err := auth.ValidateRefreshToken(req.Token)
	if err != nil {
		S.Log.Error("refresh tokin is invalid", err)
		return &pb.CheckRefreshTokenResponse{Acces: false}, err
	}
	id, err := auth.GetUserIdFromRefreshToken(req.Token)
	if err != nil {
		S.Log.Error("error while reading user id from refresh token", err)
		return nil, err
	}
	id1 := pbu.UserID{
		UserId: id,
	}
	info, err := S.Repo.GetUserByID(ctx, &id1)
	if err != nil {
		S.Log.Error("error while taking user infos", err)
		return nil, err
	}
	res := pb.LoginResponse{
		Access: &pb.AccessToken{
			Id:       info.Id,
			Username: info.Username,
			Email:    info.Email,
		},
	}
	err = auth.GeneratedAccessJWTToken(&res)
	S.Log.Info("acces token is revreated")
	if err != nil {

		return nil, err
	}
	S.Log.Info("CheckRefreshToken service completed succesfully")
	return &pb.CheckRefreshTokenResponse{Acces: true, Accestoken: res.Access.Accesstoken}, nil
}
