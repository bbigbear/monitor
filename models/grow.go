package models

import (
	"time"
)

//成绩
type Score struct {
	Id      int64
	Sid     string
	Sname   string
	Subject string
	Grade   float64
}

//阅读
type Read struct {
	Id       int64
	Sid      string
	Sname    string
	BookType string
	BookName string
	ReadTime time.Time
}

//科研
type Research struct {
	Id           int64
	Sid          string
	Sname        string
	ProjectTime  time.Time
	ProjectName  string
	ProjectLevel string
}

//上网
type Surfer struct {
	Id         int64
	Sid        string
	Sname      string
	SurferTime time.Time
	Duration   float64
}

//纪律
type Discipline struct {
	Id         int64
	Sid        string
	Sname      string
	Style      string
	RecordTime time.Time
}

//消费
type Consume struct {
	Id          int64
	Sid         string
	Sname       string
	Style       string
	money       float64
	ConsumeTime time.Time
}

//学生
type Student struct {
	Id                int64
	Name              string
	Major             string
	Faculty           string
	BriefIntroduction string
}
