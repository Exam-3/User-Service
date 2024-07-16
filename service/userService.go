package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"

	"user_service/pkg/logger"
	"user_service/storage/postgres"

	pbi "user_service/genproto/item"
	pb "user_service/genproto/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Repo       *postgres.UserRepo
	Logger *slog.Logger
	ItemClient pbi.ItemServiceClient
}

func NewUserService(db *sql.DB, item pbi.ItemServiceClient) *UserService {
	return &UserService{
		Repo:       postgres.NewUserRepository(db),
		ItemClient: item,
		Logger: logger.NewLogger(),
	}
}

func (u *UserService) GetUserProfile(ctx context.Context, id *pb.UserID) (*pb.GetUserProfileResponse, error) {
	u.Logger.Info("get user ishlashni boshladi")
	user, err := u.Repo.GetUserByID(ctx, id)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("bazadan ma'lumot olishda xato: %v", err))
		return nil, errors.Wrap(err, "get user profile failed")
	}
	
	u.Logger.Info("get user profile ishlashni tugatdi")
	return user, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateProfileResponse, error) {
	u.Logger.Info("Updating user profile ishlashni boshladi")

	user, err := u.Repo.UpdateUserProfile(ctx, req)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("bazadan ma'lumotni o'zgartirishda xato: %v", err))
		return nil, errors.Wrap(err, "update user profile failed")
	}

	u.Logger.Info("Updating user profile ishlashni tugatdi")
	return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	
	u.Logger.Info("Deleting user ishlashni boshladi")

	id := pb.UserID{
		UserId: req.UserId,
	}
	user, err := u.Repo.DeleteUser(ctx, &id)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("bazadan ma'lumotni o'chirishda xato: %v", err))
		return nil, errors.Wrap(err, "delete user failed")
	}

	u.Logger.Info("Deleting user ishlashni tugatdi")
	return user, nil
}

func (u *UserService) GetEcoPoints(ctx context.Context, req *pb.GetEcoPointsRequest) (*pb.GetEcoPointsResponse, error) {
	
	u.Logger.Info("Get eco points ishlashni boshladi")

	eco, err := u.Repo.GetEcoPoints(ctx, req)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("bazadan eco pointsni olishda xato: %v", err))
		return nil, errors.Wrap(err, "get eco points failed")
	}

	u.Logger.Info("Get eco points ishlashni tugatdi")
	return eco, nil
}

func (u *UserService) AddEcoPoints(ctx context.Context, req *pb.AddEcoPointsRequest) (*pb.AddEcoPointsResponse, error) {
	
	u.Logger.Info("Add eco points ishlashni boshladi")

	eco, err := u.Repo.AddEcoPoints(ctx, req)
    if err!= nil {
		u.Logger.Error(fmt.Sprintf("bazaga eco pointsni qo'shishda xato: %v", err))
		return nil, errors.Wrap(err, "add eco points failed")
    }

	eco.AddedPoints = req.Points
	eco.Reason = "Successful item swap"

	u.Logger.Info("Add eco points ishlashni tugatdi")
    return eco, nil
}

func (u *UserService) GetEcoPointsHistory(ctx context.Context, req *pb.GetEcoPointsHistoryRequest) (*pb.GetEcoPointsHistoryResponse, error) {
	
	u.Logger.Info("Get eco points history ishlashni boshladi")
	history, err := u.Repo.GetEcoPointsHistory(ctx, req)
    if err!= nil {
		u.Logger.Error(fmt.Sprintf("bazadan eco pointsni tarihlardan iborat hisori olishda xato: %v", err))
        return nil, errors.Wrap(err, "get eco points history failed")
    }

	u.Logger.Info("Get eco points history ishlashni tugatdi")
    return history, nil
}


func (u *UserService) ValidateUser(ctx context.Context, id *pb.ValidateUserId) (bool, error) {

	u.Logger.Info("Validate user ishlashni boshladi")
	status, err := u.Repo.ValidateUser(ctx, id)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("bazadan ma'lumotni tasdiqlashda xato: %v", err))
		return false, errors.Wrap(err, "validate user failed")
	}

	u.Logger.Info("Validate user ishlashni tugatdi")
	return status, nil
}
