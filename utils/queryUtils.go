package utils

import (
	"database/sql"
	"strconv"
	"time"
	"github.com/lib/pq"
)

func NullableFloat (s string) sql.NullFloat64 {
	if s != "" {
		value, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return sql.NullFloat64{value, true}
		}
	}
	return sql.NullFloat64{}
}
func NullableInt (s string) sql.NullInt64 {
	if s != "" {
		value, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return sql.NullInt64{value, true}
		}
	}
	return sql.NullInt64{}
}
func NullableString (s string) sql.NullString {
	return sql.NullString{s, s != ""}
}

func NullableTime (s string) pq.NullTime {
	if s != "" {
		const (layout = "2006-01-02"
		       layout1 = "02.01.2006")

		value, err := time.Parse(layout, s)
		if err == nil {
			return pq.NullTime{value, true}
		}
		value1, err1 := time.Parse(layout1, s)
		if err1 == nil {
			return pq.NullTime{value1, true}
		}
	}
	return  pq.NullTime{}
}
