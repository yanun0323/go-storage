package internal

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Start(conf Config) error {
	db, err := NewDB(conf)
	if err != nil {
		return fmt.Errorf("new db, err: %v", err)
	}

	_ = db.AutoMigrate(&File{})

	fm := NewFileManager(db)

	e := echo.New()
	p := e.Group("", middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)), TokenAuth(conf))

	e.GET("/file/:filename", fm.GetFile)
	p.POST("/file", fm.PostFile)

	ad := ":" + viper.GetString("server.port")
	e.Logger.Fatal(e.Start(ad))
	return nil
}

func TokenAuth(conf Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if len(token) == 0 || !strings.EqualFold(token, conf.Upload.Token) {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}
