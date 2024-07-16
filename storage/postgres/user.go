package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"
	// "os/user"

	pb "user_service/genproto/user"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetUserByID(ctx context.Context, id *pb.UserID) (*pb.GetUserProfileResponse, error) {

	query := `
        SELECT id, username, email, full_name, eco_points, created_at
        FROM users
        WHERE deleted_at IS NULL AND id = $1
		`

	resp := pb.GetUserProfileResponse{}
	row := u.DB.QueryRowContext(ctx, query, id.UserId)
	err := row.Scan(&resp.Id, &resp.Username, &resp.Email, &resp.FullName, &resp.EcoPoints, &resp.CreatedAt)
	if err != nil {
		log.Println("user not found")
		return nil, err
	}

	return &resp, nil
}

func (u *UserRepo) GetUsers(ctx context.Context, filter *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	query := `
        SELECT id, username, full_name, eco_points
        FROM users
        WHERE deleted_at IS NULL
		limit = $1
		Offset = $2
    `

	rows, err := u.DB.QueryContext(ctx, query, filter.Limit, filter.Page)
	if err != nil {
		log.Println("failed to get users:", err)
		return nil, err
	}
	defer rows.Close()

	users := &pb.GetUsersResponse{}

	for rows.Next() {
		user := &pb.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.FullName, &user.EcoPoints)
		if err != nil {
			log.Println("failed to scan row:", err)
			return nil, err
		}
		users.Users = append(users.Users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("rows iteration error:", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) UpdateUserProfile(ctx context.Context, user *pb.UpdateUserProfileRequest) (*pb.UpdateProfileResponse, error) {
	query := `
        UPDATE users
        SET full_name = $1, bio = $2, updated_at = NOW()
        WHERE id = $3 AND deleted_at IS NULL
        RETURNING id, username, email, full_name, bio, updated_at
    `

	resp := pb.UpdateProfileResponse{}
	row := u.DB.QueryRowContext(ctx, query, user.FullName, user.Bio, user.UserId)
	err := row.Scan(&resp.Id, &resp.Username, &resp.Email, &resp.FullName, &resp.Bio, &resp.UpdatedAt)
	if err != nil {
		log.Println("user not found")
		return nil, err
	}
	return &resp, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id *pb.UserID) (*pb.DeleteUserResponse, error) {
	query := `
        UPDATE users
            SET deleted_at = NOW()
        WHERE id = $1 AND deleted_at IS NULL
    `

	_, err := u.DB.ExecContext(ctx, query, id.UserId)
	if err != nil {
		log.Println("user not found")
		return nil, err
	}

	user := pb.DeleteUserResponse{
		Message: "User deleted successfully",
	}

	return &user, nil
}

func (u *UserRepo) GetEcoPoints(ctx context.Context, eco *pb.GetEcoPointsRequest) (*pb.GetEcoPointsResponse, error) {
	query := `
        SELECT id, eco_points, updated_at
        WHERE id = $1
    `

	resp := pb.GetEcoPointsResponse{}
	row := u.DB.QueryRowContext(ctx, query, eco.UserId)
	err := row.Scan(&resp.UserId, &resp.EcoPoints, &resp.LastUpdated)
	if err != nil {
		log.Println("user not found")
		return nil, err
	}

	return &resp, nil
}

func (u *UserRepo) AddEcoPoints(ctx context.Context, eco *pb.AddEcoPointsRequest) (*pb.AddEcoPointsResponse, error) {
	query := `
        UPDATE users
        SET eco_points = eco_points = $1, points = $2, reason = $3, updated_at = NOW()
        WHERE id = $4 AND deleted_at IS NULL
        RETURNING id, eco_points, updated_at
        `
    user := pb.AddEcoPointsResponse{}
    err := u.DB.QueryRowContext(ctx, query,eco.Points, eco.Points, eco.Reason, eco.UserId).Scan(
		&user.UserId, &user.EcoPoints, &user.Timestamp,
	)

    if err!= nil {
        log.Println("user not found")
        return nil, err
    }
    return &user, nil
}


func (u *UserRepo) GetEcoPointsHistory(ctx context.Context, req *pb.GetEcoPointsHistoryRequest) (*pb.GetEcoPointsHistoryResponse, error) {
	query := `
		SELECT id, points, reason, updated_at
		FROM users 
		WHERE id = $1
		OFFSET $2
		LIMIT $3
		`
	rows, err := u.DB.QueryContext(ctx, query,req.UserId, req.Page, req.Limit)
	if err!= nil {
        log.Println("failed to get eco points history:", err)
        return nil, err
    }
	defer rows.Close()

	history := &pb.GetEcoPointsHistoryResponse{}

	for rows.Next() {
		item := &pb.EcoPointTransaction{}

        err := rows.Scan(&item.Id, &item.Points, &item.Reason, &item.Timestamp)
        if err!= nil {
            log.Println("failed to scan row:", err)
            return nil, err
        }

		if item.Reason == "Successful item swap" {
			item.Type = "earned"
		} else {
			item.Type = "pending"
		}

        history.History = append(history.History, item)
	}

	history.Total = 1
	history.Page = req.Page
	history.Limit = req.Limit
	
	return history, nil

}



func (u *UserRepo) ValidateUser(ctx context.Context, id *pb.ValidateUserId) (bool, error) {
	query := `
    select EXISTS (
		select 1
		from users
		where id = $1 AND deleted_at IS NULL
	)`

	var status bool
	err := u.DB.QueryRowContext(ctx, query, id.Id).Scan(&status)
	if err != nil {
		log.Println("failed to scan user")
		return false, err
	}
	if !status {
		return false, errors.New("not found")
	}

	return status, nil
}
