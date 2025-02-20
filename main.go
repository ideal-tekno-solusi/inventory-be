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

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
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
