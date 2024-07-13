package service

import (
	"context"
	"database/sql"
	// pbr "user_service/genproto/itemservice"
	pb "user_service/genproto/user"
	// "user_service/pkg/logger"
	"user_service/storage/postgres"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Repo              *postgres.UserRepo
	// ReservationClient pbr.ReservationClient
}

// func NewUserService(db *sql.DB, item pbi.ItemClient) *UserService {
// 	return &UserService{
// 		Repo:              postgres.NewUserRepository(db),
// 		// Res: reser,
// 	}
// }

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
        Repo: postgres.NewUserRepository(db),
    }
}

func (u *UserService) GetUserProfile(ctx context.Context,id *pb.UserID) (*pb.GetUserProfileResponse, error) {
	user, err := u.Repo.GetUserByID(ctx, id)
	if err!= nil {
        return nil, errors.Wrap(err, "get user profile failed")
    }
	return user, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateProfileResponse, error) {
	user, err := u.Repo.UpdateUserProfile(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "update user profile failed")
    }
    return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := pb.UserID{
		UserId: req.UserId,
	}
	user, err := u.Repo.DeleteUser(ctx, &id)
    if err!= nil {
        return nil, errors.Wrap(err, "delete user failed")
    }
	// status, err := u.ReservationClient.DeleteReservationByUserID(ctx, &pbr.ID{Id: req.Id})
	// if err != nil {
	// 	u.Logger.Error(errors.Wrap(err, "failed to delete user reservations").Error())
	// 	return nil, errors.Wrap(err, "failed to delete user reservations")
	// }

	// if !status.Successful {
	// 	u.Logger.Error("deletion of user reservations unsuccessful")
	// 	return nil, errors.New("deletion of user reservations unsuccessful")
	// }
    return user, nil
}

func (u *UserService) GetEcoPoints(ctx context.Context, req *pb.GetEcoPointsRequest) (*pb.GetEcoPointsResponse, error) {
	eco, err := u.Repo.GetEcoPoints(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "get eco points failed")
    }
    return eco, nil
}

func (u *UserService) ValidateUser(ctx context.Context, id string) (bool, error) {
	status, err := u.Repo.ValidateUser(ctx, id)
    if err!= nil {
        return false, errors.Wrap(err, "validate user failed")
    }
	
    return status, nil
}