package utils

import (
	"time"
	"github.com/lib/pq"
)

func FormatDate(Time time.Time) string {
	return Time.Format("02.01.2006")
}

func FormatDatePqNullTime(Time pq.NullTime) string {
	if Time.Valid {
		return FormatDate(Time.Time)
	} else {
		return ""
	}
}
