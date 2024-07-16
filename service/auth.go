package service

import (
	"context"
	"database/sql"
	"log/slog"
	"github.com/pkg/errors"
	l "user_service/pkg/logger"
	// "user_service/email"
	"user_service/api/auth"
	"user_service/storage/postgres"
	
	pb "user_service/genproto/authentication"
	pbu "user_service/genproto/user"
)

type AuthService struct {
	pb.UnimplementedAuthenticationServer
	Repo *postgres.UserRepo
	Log  *slog.Logger
	// mailer email.Mailer
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		Repo: postgres.NewUserRepository(db),
		Log:  l.NewLogger(),
	}
}

func (S *AuthService) Register(ctx context.Context, req *pb.UserDetails) (*pb.ID, error) {

	S.Log.Info("Register ishlashni boshladi")

	res, err := S.Repo.Register(ctx, req)
	if err != nil {

		err = errors.Wrap(err, "bazaga register qo'shishda xatolik.")
		S.Log.Error(err.Error())
		return nil, err
	}

	S.Log.Info("Register ishini tugatdi")
	return res, err
}

func (S *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	S.Log.Info("login ishini boshladi")

	checker, err := S.Repo.GetUserByUsername(ctx, req.Username)
	if err != nil {

		err = errors.Wrap(err, "bazadan user name topilnadi.")
		S.Log.Error(err.Error())
		return nil, err
	}

	if req.Password != checker.Password {

		S.Log.Error("tekshirilgan password xato")
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

	S.Log.Info("refresh token generatsiya qilishni boshladi")

	if err != nil {
		S.Log.Error("refresh token generatsiya qilishda xatolik ",err)
		return nil, err
	}
	err = auth.GeneratedAccessJWTToken(res)

	S.Log.Info("acces token generatsiya qilishni boshladi")

	if err != nil {
		S.Log.Error("acces token generatsiya qilishda xatolik ",err)
		return nil, err
	}
	err = S.Repo.StoreRefreshToken(ctx, res)

	S.Log.Info("Refresh token kiritishni boshladi")

	if err != nil {
		S.Log.Error("refresh token kiritishda xatolik")
		return nil, err
	}

	S.Log.Info("login service ishini tugatdi")
	return res, nil
}

func (S *AuthService) CheckRefreshToken(ctx context.Context, req *pb.CheckRefreshTokenRequest) (*pb.CheckRefreshTokenResponse, error) {
	S.Log.Info("CheckRefreshToken service ishini boshladi")
	
	_, err := auth.ValidateRefreshToken(req.Token)
	if err != nil {

		S.Log.Error("xato refresh tokin kiritildi", err)
		return &pb.CheckRefreshTokenResponse{Acces: false}, err
	}
	id, err := auth.GetUserIdFromRefreshToken(req.Token)

	S.Log.Info("GetUserIdFromRefreshToken ishini boshladi")
	if err != nil {

		S.Log.Error("get user from refresh tokendan malumot olishda xatolik", err)
		return nil, err
	}
	id1 := pbu.UserID{
		UserId: id,
	}
	info, err := S.Repo.GetUserByID(ctx, &id1)
	S.Log.Info("GetUserByID ishini boshladi")
	if err != nil {

		S.Log.Error("bazadan user topilmadi", err)
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
	S.Log.Info("acces token generatsiya qilishni boshladi")
	if err != nil {
		S.Log.Error("acces token generatsiya qilishda xatolik ", err)
		return nil, err
	}
	S.Log.Info("CheckRefreshToken service ishini tugatdi")
	return &pb.CheckRefreshTokenResponse{Acces: true, Accestoken: res.Access.Accesstoken}, nil
}


func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {

	s.Log.Info("logout service ishini boshladi")
    err := s.Repo.DeleteRefreshToken(ctx, req.UserId)
    if err != nil {
		s.Log.Error("refresh token o'chirishda xatolik", err)
        return nil, err
    }

    return &pb.LogoutResponse{
        Message: "Successfully logged out",
    }, nil
}

// func (s *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
// 	// Check if email exists in database
// 	user, err := s.Repo.GetUserByEmail(ctx, req.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Generate reset password token and save to DB (optional)

// 	// Send reset password email
// 	err = s.mailer.SendResetPasswordEmail(user.Email, user.Username)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ResetPasswordResponse{Message: "Instructions sent"}, nil
// }