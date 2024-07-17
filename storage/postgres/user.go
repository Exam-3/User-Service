package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
    var (
        params []interface{}
    )

    query1 := `SELECT id, username, full_name, eco_points FROM users WHERE deleted_at IS NULL`

    count := 0

    if len(filter.Username) > 0 {
        count++
        query1 += fmt.Sprintf(" AND username = $%d", count)
        params = append(params, filter.Username)
    }

    if len(filter.FullName) > 0 {
        count++
        query1 += fmt.Sprintf(" AND full_name = $%d", count)
        params = append(params, filter.FullName)
    }

    if filter.EcoPoints > 0 {
        count++
        query1 += fmt.Sprintf(" AND eco_points = $%d", count)
        params = append(params, filter.EcoPoints)
    }

    if filter.Limit > 0 {
        count++
        query1 += fmt.Sprintf(" LIMIT $%d", count)
        params = append(params, filter.Limit)
    }

    if filter.Page > 0 {
        count++
        query1 += fmt.Sprintf(" OFFSET $%d", count)
        params = append(params, (filter.Page-1)*filter.Limit)
    }

    rows, err := u.DB.QueryContext(ctx, query1, params...)
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

	fmt.Println(user)
	var (
		params []interface{}
	)
	query1 := `
        UPDATE users
        SET updated_at = NOW()`

	query := ""

	ret := `
	RETURNING id, username, email, full_name, bio, updated_at
    `
	count := 0
	if len(user.Username) > 0 {
		params = append(params, user.Username)
		count++
		query = fmt.Sprintf(", username = $%d ", count)
		query1 += query
	}

	if len(user.FullName) > 0 {
		params = append(params, user.FullName)
		count++
		query = fmt.Sprintf(", full_name = $%d ", count)
		query1 += query

	}

	count++
	filter := fmt.Sprintf("WHERE id = $%d AND deleted_at IS NULL", count)
	query1 += filter + ret
	params = append(params, user.UserId)
	fmt.Println(params)
	fmt.Println(query1)

	resp := pb.UpdateProfileResponse{}

	row := u.DB.QueryRowContext(ctx, query1, params...)
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
		FROM users
        WHERE id = $1 AND deleted_at IS NULL
    `

	resp := &pb.GetEcoPointsResponse{} // create pointer to struct
	row := u.DB.QueryRowContext(ctx, query, eco.UserId)
	err := row.Scan(&resp.UserId, &resp.EcoPoints, &resp.LastUpdated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user not found")
			return nil, fmt.Errorf("user with id %s not found", eco.UserId)
		}
		log.Println("failed to scan row:", err)
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) AddEcoPoints(ctx context.Context, eco *pb.AddEcoPointsRequest) (*pb.AddEcoPointsResponse, error) {
	query := `
        UPDATE users
        SET eco_points = eco_points + $1, points = $2, reason = $3, updated_at = NOW()
        WHERE id = $4 AND deleted_at IS NULL
        RETURNING id, eco_points, updated_at
        `
	user := pb.AddEcoPointsResponse{}
	err := u.DB.QueryRowContext(ctx, query, eco.Points, eco.Points, eco.Reason, eco.UserId).Scan(
		&user.UserId, &user.EcoPoints, &user.Timestamp,
	)

	if err != nil {
		log.Println("user not found")
		return nil, err
	}
	return &user, nil
}
func (u *UserRepo) GetEcoPointsHistory(ctx context.Context, filter *pb.GetEcoPointsHistoryRequest) (*pb.GetEcoPointsHistoryResponse, error) {

	var (
		params []interface{}
	)

	query1 := `SELECT id, points, created_at FROM users WHERE deleted_at IS NULL `
	query := ""

	count := 0
	if filter.Points > 0 {
		params = append(params, filter.Points)
		count++
		query = fmt.Sprintf("AND points = $%d ", count)
		query1 += query
	}

	if len(filter.Reason) > 0 {
		params = append(params, filter.Reason)
		count++
		query = fmt.Sprintf("AND reason = $%d ", count)
		query1 += query
	}

	if filter.Limit > 0 {
		count++
		query = fmt.Sprintf("LIMIT $%d ", count)
		query1 += query
		params = append(params, filter.Limit)
	}

	if filter.Page > 0 {
		count++
		query = fmt.Sprintf("OFFSET $%d ", count)
		query1 += query
		params = append(params, (filter.Page-1)*filter.Limit)
	}

	rows, err := u.DB.QueryContext(ctx, query1, params...)
	if err != nil {
		log.Println("failed to get eco points history:", err)
		return nil, err
	}
	defer rows.Close()

	history := &pb.GetEcoPointsHistoryResponse{}

	for rows.Next() {
		item := &pb.EcoPointTransaction{}

		err := rows.Scan(&item.Id, &item.Points, &item.Timestamp)
		if err != nil {
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

	countQuery := `
		SELECT count(id) FROM users WHERE deleted_at IS NULL
	`
	var total int32
	err = u.DB.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		log.Println("failed to get total count:", err)
		return nil, err
	}
	history.Total = total

	history.Page = filter.Page
	history.Limit = filter.Limit

	return history, nil
}


func (u *UserRepo) ValidateUser(ctx context.Context, id *pb.ValidateUserIdRequest) (bool, error) {
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
