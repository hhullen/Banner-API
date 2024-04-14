package main

import (
	"fmt"
	"log"
	"os"

	"main/core/model"
	"main/infrastructure/controllers"
	"main/infrastructure/storage"
)

func main() {
	// database := &storage.MOCKDB{}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		GetEnvVarOr("DB_HOST", "localhost"),
		GetEnvVarOr("DB_PORT", "5432"),
		GetEnvVarOr("DB_USER", "postgres"),
		GetEnvVarOr("DB_PASSWORD", "postgres"),
		GetEnvVarOr("DB_NAME", "postgres"),
	)

	database := storage.NewPostgreSQL(dsn)
	database.AddDefaultUser(model.User{
		Role:  model.AdminRole,
		Token: os.Getenv("DEFAULT_ADMIN_TOKEN"),
	})
	database.AddDefaultUser(model.User{
		Role:  model.UserRole,
		Token: os.Getenv("DEFAULT_USER_TOKEN"),
	})

	ctrl := controllers.NewControllerREST(database)
	log.Fatal(ctrl.Serve(":8080"))
}

func GetEnvVarOr(env_var, default_var string) string {
	variable := os.Getenv(env_var)
	if len(variable) == 0 {
		return default_var
	}
	return variable
}
