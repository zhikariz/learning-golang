package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// initialisasi database
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	// inisialisasi handler
	userHandler := NewUserHandler(db)

	e := echo.New()
	// routing
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// table / entity user
type User struct {
	ID     int64  `json:"id"`
	Nim    string `json:"nim"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

func (User) TableName() string {
	return "users"
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

type UserRequest struct {
	ID     string `param:"id"`
	Nim    string `json:"nim"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

func (h *UserHandler) GetAllUsers(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	users := make([]*User, 0)
	query := h.db.Model(&User{})
	if search != "" {
		query = query.Where("nama LIKE ?", "%"+search+"%")
	}
	if err := query.Find(&users).Error; err != nil { // SELECT * FROM users
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Get All Users"})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Succesfully Get All Users", "data": users, "filter": search})
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input UserRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to Bind Input"})
	}

	user := &User{
		Nim:    input.Nim,
		Nama:   input.Nama,
		Alamat: input.Alamat,
	}

	if err := h.db.Create(user).Error; err != nil { // INSERT INTO users (nim, nama, alamat) VALUES('')
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Create User"})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{"message": "Succesfully Create a User", "data": user})
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	var input UserRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to Bind Input"})
	}

	user := new(User)

	if err := h.db.Where("id =?", input.ID).First(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Get User By ID"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Succesfully Get User By ID : %s", input.ID), "data": user})
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input UserRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to Bind Input"})
	}

	userID, _ := strconv.Atoi(input.ID)

	user := User{
		ID:     int64(userID),
		Nim:    input.Nim,
		Nama:   input.Nama,
		Alamat: input.Alamat,
	}

	query := h.db.Model(&User{}).Where("id = ?", userID)
	if err := query.Updates(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Update User By ID", "error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Succesfully Update User By ID : %s", input.ID), "data": input})
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input UserRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to Bind Input"})
	}

	if err := h.db.Delete(&User{}, input.ID).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to Delete User By ID"})
	}
	return ctx.JSON(http.StatusNoContent, nil)
}

func initDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/db_user?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
