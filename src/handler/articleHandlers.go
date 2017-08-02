package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"math/rand"
	"model"
	"net/http"
	"time"
)

func (h *handler) GetArticle(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)
	result["path"] = c.Path()
	result["method"] = c.Request().Method
	result["id"] = c.Param("id")
	result["status_code"] = http.StatusOK

	return c.JSONPretty(http.StatusOK, result, "    ")
}

func (h *handler) ListArticle(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)
	result["articles"] = model.NewArticleSlice()

	//since := c.Param("since")
	//timePoint := getTimePoint(t)
	//topic := c.Param("topic")

	return c.JSONPretty(http.StatusOK, result, "    ")
}

func (h *handler) CreateArticle(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)
	id := int(rand.Int31())
	result["id"] = id
	messageChan := make(chan map[string]interface{})
	errorChan := make(chan map[string]interface{})

	go func(c echo.Context, messageChan, errorChan chan map[string]interface{}) {
		a := model.NewArticle(id, "", "")
		if err := c.Bind(a); err != nil {
			result["status"] = "ERROR"
			result["article"] = nil
			result["error"] = fmt.Sprint(err)
			result["status_code"] = http.StatusInternalServerError
			logrus.Error("CreateArticle.Bind ERROR: ", err)
			errorChan <- result
			return
		}

		//数据库操作

		result["status"] = "CREATED"
		result["article"] = a
		result["status_code"] = http.StatusCreated
		logrus.Infof("Create Article id: [ %d ] Success", id)
		messageChan <- result
	}(c, messageChan, errorChan)

	select {
	case message := <-messageChan:
		return c.JSONPretty(http.StatusCreated, message, "    ")
	case errMessage := <-errorChan:
		return c.JSONPretty(http.StatusInternalServerError, errMessage, "    ")
	case <-time.After(10 * time.Second):
		logrus.Info("CreateArticle timeout")
		result["status"] = "TIMEOUT"
		result["article"] = nil
		result["status_code"] = http.StatusGatewayTimeout
		return c.JSONPretty(http.StatusGatewayTimeout, result, "    ")
	}
}

func (h *handler) UpdateArticle(c echo.Context) error {
	result := make(map[string]interface{})
	t := time.Now()
	result["time"] = t.Format(layout)
	result["method"] = c.Request().Method
	result["path"] = c.Path()
	result["id"] = c.Param("id")

	return c.JSONPretty(http.StatusOK, result, "    ")
}

func (h *handler) DeleteArticle(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)
	result["method"] = c.Request().Method
	result["path"] = c.Path()
	result["id"] = c.Param("id")

	return c.JSONPretty(http.StatusNoContent, result, "    ")
}

func (h *handler) SearchArticle(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)

	return c.JSONPretty(http.StatusOK, result, "    ")
}

func (h *handler) LikeArticle(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)

	return c.JSONPretty(http.StatusOK, result, "    ")
}
