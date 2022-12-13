package main

import (
	"context"
	"fizzbuzz/config"
	"fizzbuzz/database"
	"fmt"

	"time"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func main() {
	config.InitConfig()

	// Loger
	zaplog, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer zaplog.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ConnectDB
	var conn database.Database
	db := conn.GetConnectionDB()

	fmt.Println(db, ctx)
}
