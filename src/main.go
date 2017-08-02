package main

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "huanyu0w0" && password == "3.1415926" {
			return true, nil
		}
		return false, nil
	}))
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Addr:     "119.29.243.98:5432",
	})
	defer db.Close()

	h := handler.NewHandler(db)
	//article api
	e.GET("/v1/articles", h.ListArticle)
	e.GET("/v1/articles/:id", h.GetArticle)
	e.POST("/v1/topics/:topicid/articles", h.CreateArticle)
	e.PATCH("/v1/articles/:id", h.UpdateArticle)
	e.DELETE("/v1/articles/:id", h.DeleteArticle)
	//e.GET("/v1/articles/search", h.SearchArticle)
	e.GET("/v1/articles/:id/like", h.LikeArticle)
	//comment api
	e.GET("/v1/articles/:articleid/comments", h.ListComment)
	e.POST("/v1/articles/:articleid/comments", h.CreateComment)
	e.DELETE("/v1/articles/:articleid/comments/:commentid", h.DeleteComment)
	//topic api
	e.GET("/v1/topics", h.ListTopic)
	e.GET("/v1/topics/:id", h.GetTopic)
	e.POST("/v1/topics", h.CreateTopic)
	e.PUT("/v1/topics/:id", h.UpdateTopic)
	e.DELETE("/v1/topics/:id", h.DeleteTopic)
	logrus.Error(e.Start(":1323"))
}
