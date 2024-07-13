package postgres

import (
	"context"
	pb "user_service/genproto/authentication"
)

func (u *UserRepo) Register(ctx context.Context, user *pb.UserDetails) (*pb.ID, error) {
	var id pb.ID
	query := `
	insert into users (
		username, email, password
	)
	values (
		$1, $2, $3
	)
	returning id`

	err := u.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&id.Id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (u *UserRepo) GetUserByUsername(ctx context.Context, username string) (*pb.UserInfo, error) {
	user := pb.UserInfo{Username: username}
	query := `
	select
		id, email, password
	from
		users
	where
		deleted_at is null and username = $1
	`

	err := u.DB.QueryRowContext(ctx, query, username).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) StoreRefreshToken(ctx context.Context, token *pb.LoginResponse) error {
	query := `
	insert into refresh_tokens (
		user_id, token
	)
	values (
		$1, $2
	)`

	_, err := u.DB.ExecContext(ctx, query, token.Refresh.Userid, token.Refresh.Refreshtoken)
	return err
}

func (u *UserRepo) DeleteRefreshToken(ctx context.Context, userID string) error {
	query := `
	delete from
		refresh_tokens
	where
		user_id = $1`

	_, err := u.DB.ExecContext(ctx, query, userID)
	return err
}
