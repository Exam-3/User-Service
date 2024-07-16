// package handler

// import (
// 	"database/sql"
// 	"log/slog"
// 	"user_service/service"
// 	"user_service/storage/postgres"
// 	"user_servise/service"
// )

// type Handler struct {
// 	Auth *service.AuthService
// 	Log  *slog.Logger
// }

// func NewHandler(db *sql.DB) *Handler {
// 	return &Handler{
//         Auth: service.NewAuthService(db),
//         // Log:  slog.NewLogger(),
//     }
// }

package handler

import (
	"log/slog"
	"user_service/genproto/authentication"
)

type Handler struct {
	Auth authentication.AuthenticationClient
	Log  *slog.Logger
}