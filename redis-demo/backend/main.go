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

func signup(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	hashedPassword := hashPassword(user.Password)

	_, err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", user.Username, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}

func login(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	hashedPassword := hashPassword(user.Password)

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE name = ?", user.Username).Scan(&storedPassword)
	debug("a")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	debug("b")
	debug(storedPassword)
	debug(hashedPassword)
	if storedPassword != hashedPassword {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	debug("c")

	sessionID := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Username+time.Now().String())))
	err = rdb.Set(ctx, sessionID, user.Username, 24*time.Hour).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create session"})
	}

	// CookieにセッションIDを設定
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
}

func logout(c echo.Context) error {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "No session ID found"})
	}

	sessionID := cookie.Value

	// RedisからセッションIDを削除
	err = rdb.Del(ctx, sessionID).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not delete session"})
	}

	// クッキーを無効化
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing session ID"})
		}

		sessionID := cookie.Value
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
	db, err = sql.Open("mysql", "mysqluser:mysqlpassword@tcp(mysql:3306)/redisdemo")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	e.POST("/signup", signup)
	e.POST("/login", login)
	e.POST("/logout", logout)
	e.GET("/restricted", restrictedEndpoint, authMiddleware)

	e.Logger.Fatal(e.Start(":8888"))
}

func debug(msg string) {
	fmt.Println(msg)
}
