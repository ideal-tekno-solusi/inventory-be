package handler

import (
	rest "app/api/inventory"
	"app/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type RestService struct {
	dbr *pgx.Conn
	dbw *pgx.Conn
}

func RestRegister(r *gin.Engine, cfg *bootstrap.Container) {
	rest.Router(r, &RestService{
		dbr: cfg.Dbr(),
		dbw: cfg.Dbw(),
	})
}
