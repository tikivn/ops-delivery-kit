package util

import (
	"net/url"
	"strconv"
)

func Pager(vals url.Values, defaultLimit, defaultPage int) (page, limit int) {
	if str := vals.Get("limit"); str != "" {
		limit, _ = strconv.Atoi(str)
	}
	if limit <= 0 {
		limit = defaultLimit
	}
	vals.Set("limit", strconv.Itoa(limit))

	if str := vals.Get("page"); str != "" {
		page, _ = strconv.Atoi(str)
	}
	if page <= 0 {
		page = defaultPage
	}
	vals.Set("page", strconv.Itoa(page))

	return page, limit
}
