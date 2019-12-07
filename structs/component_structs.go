package structs

import (
	"fmt"
	"time"
)

type Component struct {
}

type Results interface {
}

func (c *Component) TblLead() string {
	return "lead"
}

func (c *Component) GetMessageSucc() string {
	return "success"
}

func (c *Component) GetMessageErr() string {
	return "sistem error"
}

func (c *Component) GetTimeNow() string {

	t := time.Now()
	return string(t.Format("2006-01-02 15:04:05.999999"))
}

type Marshalers interface {
	MarshalJSON() ([]byte, error)
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
