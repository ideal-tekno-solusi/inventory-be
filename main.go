package main

import (
	"app/api"
	"app/bootstrap"
	"app/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	cfg := bootstrap.InitContainer()
	env := viper.GetString("config.env")

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	// TODO: cek lagi CORS ini
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	csrfSecret := viper.GetString("config.csrf.secret")
	csrfDomain := viper.GetString("config.csrf.domain")
	csrfPath := viper.GetString("config.csrf.path")
	csrfAge := viper.GetInt("config.csrf.age")

	CSRF := csrf.Protect(
		[]byte(csrfSecret),
		csrf.Domain(csrfDomain),
		csrf.Path(csrfPath),
		csrf.MaxAge(csrfAge),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			errorMessage := r.Context().Value("gorilla.csrf.Error").(error)
			utils.SendProblemDetailJsonHttp(w, http.StatusForbidden, errorMessage.Error(), r.URL.Path, uuid.NewString())
		})),
	)

	api.RegisterApi(r, cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", viper.GetString("services.host"), viper.GetString("services.port")),
		Handler: CSRF(r.Handler()),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("failed to listen with error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Warn("shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg.StopDb(ctx)

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Warnf("shutdown server with error: %v", err)
	}

	select {
	case <-ctx.Done():
		logrus.Warn("timeout of 5 seconds.")
	}

	logrus.Warn("Server exiting")
}
