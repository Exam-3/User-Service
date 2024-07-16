package postgres

import (
	"context"
	// "reflect"
	"testing"
	pb "user_service/genproto/user"
)

func GetUserByID(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	test := pb.UserID{UserId: "65438b09-665b-449e-b45d-92488be222a3"}

	user, err := repo.GetUserByID(context.Background(), &test)
	if err != nil {
		t.Fatal(err)
	}

	exp := pb.GetUserProfileResponse{
		Id:        "65438b09-665b-449e-b45d-92488be222a3",
		Username:  "jane_doe",
		Email:     "jane@example.com",
		FullName:  "Jane Doe ",
		EcoPoints: 100,
		CreatedAt: "2024-07-14 19:54:21.068552",
	}

	if exp.Id != user.Id {
		t.Errorf("Expected %v, got %v", &exp, user)
	}
}

func TestGetUsers(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	users, err := repo.GetUsers(context.Background(), &pb.GetUsersRequest{})
	// exp := &pb.GetUsersResponse{
	// 	Users: []*pb.User{
	// 			Id:        "65438b09-665b-449e-b45d-92488be222a3",
	// 			Username:  "jane_doe",
	// 			Email:     "jane@example.com",
	// 			FullName:  "Jane Doe ",
	// 			EcoPoints: 100,
	// 			CreatedAt: "2024-07-14 19:54:21.068552",
	// 		},
	// 	}
	if err != nil {
		t.Errorf("Expected, got %v", users)
	}
}

func TestUpdateUserPrifile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	user := pb.UpdateUserProfileRequest{
		UserId:   "65438b09-665b-449e-b45d-92488be222a3",
		FullName: "Abbos Ali Updated",
		Bio:      "New Bio",
	}

	updatedUser, err := repo.UpdateUserProfile(context.Background(), &user)
	if err != nil {
		t.Fatal(err)
	}

	exp := &pb.UpdateProfileResponse{
		Id:        "65438b09-665b-449e-b45d-92488be222a3",
		Username:  "jone_doe",
		Email:     "abbos@example.com",
		FullName:  "Abbos Ali Updated",
		Bio:       "New Bio",
		UpdatedAt: "2022-01-01T12:00:00Z",
	}

	if &exp.Id != &updatedUser.Id {
		t.Errorf("Expected %v, got %v", exp, updatedUser)
	}
}

// func TestDeleteUser(t *testing.T) {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()

// 	repo := NewUserRepository(db)

// 	test := pb.UserID{UserId: "4924aba2-9ce5-4be3-a4f7-cf113563f4f6"}

// 	_, err = repo.DeleteUser(context.Background(), &test)
// 	if err != nil {
// 		t.Errorf("Failed to delete user")
// 	}
// }

// func TestEcoPoints(t *testing.T) {
// 	db, err := ConnectDB()
//     if err!= nil {
//         t.Fatal(err)
//     }
//     defer db.Close()

//     repo := NewUserRepository(db)

// 	test := pb.GetEcoPointsRequest{
// 		UserId: "4924aba2-9ce5-4be3-a4f7-cf113563f4f6",
// 	}

//     ecoPoints, err := repo.GetEcoPoints(context.Background(), &test)
//     if err!= nil {
//         t.Fatal(err)
//     }

// 	exp:=pb.GetEcoPointsResponse{
// 		UserId: "4924aba2-9ce5-4be3-a4f7-cf113563f4f6",
// 		EcoPoints: 100,
//         LastUpdated: "2022-01-01T12:00:00Z",
// 	}

//     if reflect.DeepEqual(&exp, ecoPoints) {
//         t.Errorf("Expected %v, got %v", &exp, ecoPoints)
//     }
// }
