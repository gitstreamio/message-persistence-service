package common

import "time"

type Message struct {
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	Timeline  string    `json:"timeline"`
	Created   time.Time `json:"created,omitempty"`
	Updated   time.Time `json:"updated,omitempty"`
	Permalink string    `json:"permalink,omitempty"`
}
