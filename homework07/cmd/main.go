package main

import (
	"awesomeProject1/homework07/config"
	"awesomeProject1/homework07/internal/cache"
	"awesomeProject1/homework07/internal/middleware"
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/note"
	"awesomeProject1/homework07/internal/user"
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	rdb, err := cache.New(cfg)
	if err != nil {
		slog.Warn("Redis connection failed, continuing without Redis", "error", err)
	} else {
		slog.Info("Redis connected successfully")
	}

	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" +
		cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&models.Note{}, &models.User{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	userHandler := user.NewUserHandler(db, cfg, rdb)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware(cfg))
	{
		users := auth.Group("/users")
		{
			users.PUT("/change-password", userHandler.ModifyPassword)
		}

		noteHandler := note.NewNoteHandler(db, rdb)
		notes := auth.Group("/notes")
		{
			notes.GET("/:id", noteHandler.GetNote)
			notes.POST("", noteHandler.CreateNote)
			notes.PUT("/:id", noteHandler.UpdateNote)
			notes.DELETE("/:id", noteHandler.DeleteNote)
		}
	}

	addr := ":" + cfg.ServerPort
	slog.Info("server starting", "addr", addr)
	r.Run(addr)
}
