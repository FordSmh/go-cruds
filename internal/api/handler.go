package api

import (
	"mindlab/internal/api/controller"
)

type Handler struct {
	Login *controller.LoginController
	Post  *controller.PostController
}
