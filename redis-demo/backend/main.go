package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	rdb *redis.Client
	db  *sql.DB
	ctx = context.Background()
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func login(c echo.Context) error {
	debug("a")
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	debug("b")

	hashedPassword := hashPassword(user.Password)

	debug("c")
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
	fmt.Println(err)
	debug("d")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}
	debug("e")

	if storedPassword != hashedPassword {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	sessionID := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Username+time.Now().String())))
	err = rdb.Set(ctx, sessionID, user.Username, 24*time.Hour).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create session"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful", "session_id": sessionID})
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := c.Request().Header.Get("Authorization")
		if sessionID == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing session ID"})
		}

		username, err := rdb.Get(ctx, sessionID).Result()
		if err == redis.Nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid session ID"})
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking session ID"})
		}

		c.Set("username", username)
		return next(c)
	}
}

func restrictedEndpoint(c echo.Context) error {
	username := c.Get("username").(string)
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome " + username})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var err error
	db, err = sql.Open("mysql", "username:rootpassword@tcp(mysql:3306)/userdb")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	e.POST("/login", login)
	e.GET("/restricted", restrictedEndpoint, authMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}

func debug(msg string) {
	fmt.Println(msg)
}
