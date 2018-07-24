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
	Money       float64
	ConsumeTime time.Time
}

//学生
type Student struct {
	Id                int64
	Sid               string
	Name              string
	Major             string
	Faculty           string
	Position          string
	Level             string
	Identity          string
	BriefIntroduction string
	PicUrl            string
	Year              int
}

//监控
type WarnSetting struct {
	Id         int64
	Style      string
	Tzfs       string
	XyyjZcj    float64
	XyyjDkcj   float64
	YktDbxf    float64
	YktDrxf    float64
	YktDyxf    float64
	SwyjCs     int
	SwyjZsc    float64
	SwyjDcswsc float64
	TsjySl     int
	TsjySj     float64
	PksYxf     float64
	TsgbmRs    int
	Gkcs       int
}
