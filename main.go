package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"imp-backend/application/misc"
	_database "imp-backend/infrastructure/persistence/repository"
	token "imp-backend/middleware"

	_ "github.com/swaggo/echo-swagger/example/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "*"},
		ExposeHeaders:    []string{"Accept", "Content-Length", "Content-Type", "Authorization", "Accept:Encoding"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	db := _database.DBInit()

	// initiate logger
	misc.InitLogger(e)

	// Health check
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm well")
	})

	// v1 routes
	v1 := e.Group("")
	
	// middleware
	v1.Use(token.JWTVerify)

	// AuthRoutes
	AuthRoutes(v1, db)
	// UserRoutes
	UserRoutes(v1, db)

	en := godotenv.Load()
	if en != nil {
		fmt.Print(en)
	}
	port := os.Getenv("APP_PORT")

	if port == "" {
		log.Fatalf(fmt.Sprintf("port must be set [%s]", port))
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func AuthRoutes(route *echo.Group, db *gorm.DB) {
	createNewAuthHandler := LoginHandler(db)
	signupHandler := SignupHandler(db)
	auth := route.Group("/auth")
	{
		auth.POST("/login", createNewAuthHandler.Login)
		auth.POST("/signup", signupHandler.Signup)
	}
}

func UserRoutes(route *echo.Group, db *gorm.DB) {
	listUserHandler := ListUserHandler(db)
	user := route.Group("/user")
	{
		user.GET("/userlist", listUserHandler.ListUser)
	}
}
