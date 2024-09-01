package main

import (
	"carrental/internal/config"
	"fmt"
)

func main() {
	conf := config.NewViper()

	config.ConfigureTracing(
		conf.GetString("log.print_level"),
		conf.GetString("log.stacktrace_level"),
		conf.GetUint("log.max_pc"),
	)

	db := config.NewDatabase(
		conf.GetString("db.host"),
		conf.GetInt("db.port"),
		conf.GetString("db.user"),
		conf.GetString("db.password"),
		conf.GetString("db.name"),
		conf.GetInt("db.pool.idle"),
		conf.GetInt("db.pool.max"),
		conf.GetInt("db.pool.lifetime"),
	)
	router := config.NewGin(conf.GetString("app.mode"))

	config.RegisterCustomValidation(conf.GetString("validator.phone_number"))

	config.Bootstrap(config.BootstrapConfig{
		DB:     db,
		Router: router,
	})

	if err := router.Run(conf.GetString("web.address")); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
