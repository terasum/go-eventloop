package event

import "time"

//Massage is a transfer massage content
type Massage struct {
	MsgType string
	Content string
	Time    time.Time
}
