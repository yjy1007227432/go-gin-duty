package cache_service

import (
	"go-gin-duty-master/e"
	"strconv"
	"strings"
)

type CacheRota struct {
	ID       int
	Datetime string
	Month    string
}

func (a *CacheRota) GetRotaKey() string {
	return e.CACHE_ROAT + "_" + strconv.Itoa(a.ID)
}

func (a *CacheRota) GetRotasKeyByDay() string {
	keys := []string{
		e.CACHE_ROAT,
		"LIST",
		"DATETIME",
	}
	if a.Datetime != "" {
		keys = append(keys, a.Datetime)
	}
	return strings.Join(keys, "_")
}

func (a *CacheRota) GetRotasKeyByMonth() string {
	keys := []string{
		e.CACHE_ROAT,
		"LIST",
		"MONTH",
	}
	if a.Month != "" {
		keys = append(keys, a.Month)
	}
	return strings.Join(keys, "_")
}
