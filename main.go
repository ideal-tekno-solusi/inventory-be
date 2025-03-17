package main

import (
	"app/api"
	"app/bootstrap"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	// TODO: cek lagi CORS ini
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cfg := bootstrap.InitContainer()

	api.RegisterApi(r, cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", viper.GetString("services.host"), viper.GetString("services.port")),
		Handler: r.Handler(),
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
