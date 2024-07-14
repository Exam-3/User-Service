package service

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"user_service/storage/postgres"

	pbi "user_service/genproto/item"
	pb "user_service/genproto/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Repo       *postgres.UserRepo
	ItemClient pbi.ItemServiceClient
}

func NewUserService(db *sql.DB, item pbi.ItemServiceClient) *UserService {
	return &UserService{
		Repo:       postgres.NewUserRepository(db),
		ItemClient: item,
	}
}

func (u *UserService) GetUserProfile(ctx context.Context, id *pb.UserID) (*pb.GetUserProfileResponse, error) {
	user, err := u.Repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get user profile failed")
	}
	return user, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateProfileResponse, error) {
	user, err := u.Repo.UpdateUserProfile(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "update user profile failed")
	}
	return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := pb.UserID{
		UserId: req.UserId,
	}
	user, err := u.Repo.DeleteUser(ctx, &id)
	if err != nil {
		return nil, errors.Wrap(err, "delete user failed")
	}

	status, err := u.ItemClient.DeleteItem(ctx, &pbi.DeleteItemRequest{UserId: req.UserId, ItemId: ""})
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete user items")
	}

	message := "User deleted successfully"
	if status.Message != message {
		return nil, errors.New("deletion of user items unsuccessful")
	}
	return user, nil
}

func (u *UserService) GetEcoPoints(ctx context.Context, req *pb.GetEcoPointsRequest) (*pb.GetEcoPointsResponse, error) {
	eco, err := u.Repo.GetEcoPoints(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "get eco points failed")
	}
	return eco, nil
}

func (u *UserService) AddEcoPoints(ctx context.Context, req *pb.AddEcoPointsRequest) (*pb.AddEcoPointsResponse, error) {
	eco, err := u.Repo.AddEcoPoints(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "add eco points failed")
    }

	eco.AddedPoints = req.Points
	eco.Reason = "Successful item swap"

    return eco, nil
}


func (u *UserService) ValidateUser(ctx context.Context, id string) (bool, error) {
	status, err := u.Repo.ValidateUser(ctx, id)
	if err != nil {
		return false, errors.Wrap(err, "validate user failed")
	}

	return status, nil
}
