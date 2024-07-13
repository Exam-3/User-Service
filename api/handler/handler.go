package handler

import (
	"log/slog"
	"user_service/genproto/authentication"
)

type Handler struct {
	Auth authentication.AuthenticationClient
	Log  *slog.Logger
}
