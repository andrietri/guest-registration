package helper

import (
	"net/url"
	"strconv"
	"strings"
)

func PagingDefault(q url.Values) (int, int, string, string, string) {
	pOffset := ExpectedInt(q.Get("offset"))
	limit := ExpectedInt(q.Get("limit"))
	order := strings.TrimSpace(strings.ToLower(q.Get("order")))

	page := ExpectedInt(q.Get("page"))
	if page != 0 {
		pOffset = page
	}

	if order == "" {
		order = "id"
	}

	sort := strings.TrimSpace(strings.ToLower(q.Get("sort")))
	if (sort != "desc" && sort != "asc") || sort == "" {
		sort = "desc"
	}

	like, _ := url.QueryUnescape(strings.TrimSpace(strings.ToUpper(q.Get("like"))))
	if like != "" {
		like = "%" + like + "%"
	}

	switch {
	case limit >= 1000:
		limit = 10
	case limit <= 0:
		limit = 100
	}

	offset := 0
	if pOffset != 0 && ExpectedInt(q.Get("limit")) != 0 {
		offset = (pOffset - 1) * limit
	}

	offsetStr := strconv.Itoa(offset)
	limitStr := strconv.Itoa(limit)
	q.Set("offset", offsetStr)
	q.Set("limit", limitStr)

	return offset, limit, order, sort, like
}

func ExpectedInt(v interface{}) int {
	var result int
	switch v.(type) {
	case int:
		result = v.(int)
	case int64:
		result = int(v.(int64))
	case float64:
		result = int(v.(float64))
	case string:
		result, _ = strconv.Atoi(v.(string))
	}
	return result
}

func ExpectedString(v interface{}) string {
	var result string
	switch v.(type) {
	case int:
		result = strconv.Itoa(v.(int))
	case int64:
		result = strconv.Itoa(int(v.(int64)))
	case float64:
		result = strconv.Itoa(int(v.(float64)))
	case string:
		result, _ = v.(string)
	}
	return result
}
