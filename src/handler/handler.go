package handler

import (
	"github.com/go-pg/pg"
	"time"
)

type (
	handler struct {
		DB *pg.DB
	}
)

const (
	layout = "2006年 01月02日 15:04:05"
)

func NewHandler(db *pg.DB) *handler {
	return &handler{db}
}

func getTimeNode(duration string) int64 {
	switch duration {
	case "day", "latest":
		return time.Now().Add(-24 * time.Hour).Unix()
	case "week":
		return time.Now().Add(-7 * 24 * time.Hour).Unix()
	case "month":
		return time.Now().Add(-30 * 24 * time.Hour).Unix()
	case "year":
		return time.Now().Add(-365 * 24 * time.Hour).Unix()
	default:
		return time.Now().Add(-24 * time.Hour).Unix()
	}
}
