package api

import (
	"context"
	"fmt"
	"fungame/pkg/config"
	"fungame/pkg/db"
	"fungame/pkg/logger"
	"fungame/pkg/service"

	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ctx      context.Context
	database []*gorm.DB
	sevr     *http.Server
)

func Start() error{
	ctx = context.Background()
	postgresConn, err := db.PostgreSQL()
	if err != nil {
		return err
	}
	//err handle
	oracleConn, err := db.OracleConnect()
	// err handle
	database = append(database, postgresConn)
	database = append(database, oracleConn)

	dbObj := db.NewDBservice(postgresConn, oracleConn)
	serObj := service.NewServiceGroup(dbObj)

	StartRouter(serObj)
	return nil

}
func StartRouter(obj service.ServiceLayer) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConfig().GetInt("server.port")),
		Handler: Route(obj, logger.Log()),
	}
	logger.Log().Info("starting routeer")

	if err := srv.ListenAndServe(); err != nil {
		logger.Log().Fatal("Error starting server", zap.Error(err))
	}
	logger.Log().Info("server working")
}
