package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"math/rand"
	"model"
	"net/http"
	"strconv"
	"time"
)

func (h *handler) ListComment(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)

	return c.JSONPretty(http.StatusOK, result, "    ")
}

func (h *handler) CreateComment(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)
	comment_id := int(rand.Int31())
	article_id, err := strconv.Atoi(c.Param("articleid"))
	if err != nil {
		result["comment_id"] = c.Param("commentid")
		result["article_id"] = c.Param("articleid")
		result["status"] = "ERROR"
		result["error"] = fmt.Sprint(err)
		result["comment"] = nil
		result["status_code"] = http.StatusInternalServerError
		return c.JSONPretty(http.StatusInternalServerError, result, "    ")
	}
	result["comment_id"] = comment_id
	result["article_id"] = article_id

	messageChan := make(chan map[string]interface{})
	errorChan := make(chan map[string]interface{})
	go func(c echo.Context, messageChan, errorChan chan map[string]interface{}) {
		comment := model.NewComment(comment_id, article_id, "", "", "")
		if err = c.Bind(comment); err != nil {
			result["status"] = "ERROR"
			result["comment"] = nil
			result["error"] = fmt.Sprint(err)
			result["status_code"] = http.StatusInternalServerError
			logrus.Error("CreateComment.Bind ERROR: ", err)
			errorChan <- result
			return
		}

		//数据库操作

		result["status"] = "CREATED"
		result["comment"] = comment
		result["status_code"] = http.StatusCreated
		logrus.Infof("Create Comment id: [ %d ] in Article( id: [ %d ] ) Success", comment_id, article_id)
		messageChan <- result
	}(c, messageChan, errorChan)

	select {
	case message := <-messageChan:
		return c.JSONPretty(http.StatusCreated, message, "    ")
	case errMessage := <-errorChan:
		return c.JSONPretty(http.StatusInternalServerError, errMessage, "    ")
	case <-time.After(10 * time.Second):
		logrus.Info("CreateComment timeout")
		result["status"] = "TIMEOUT"
		result["comment"] = nil
		result["status_code"] = http.StatusGatewayTimeout
		return c.JSONPretty(http.StatusGatewayTimeout, result, "    ")
	}
}

func (h *handler) DeleteComment(c echo.Context) error {
	result := make(map[string]interface{})
	result["time"] = time.Now().Format(layout)

	return c.JSONPretty(http.StatusNoContent, result, "    ")
}
