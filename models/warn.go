package models

import (
	"time"
)

type Warn struct {
	Id        int64
	Studentid string
	Sname     string
	WarnName  string
	WarnInfo  string
	WarnTime  time.Time
	Remark    string
	Status    string
}

type WarnStyle struct {
	Id   int64
	Name string
}
