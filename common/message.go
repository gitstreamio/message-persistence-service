package common

import "time"

/*
Model for our Message
See https://github.com/gitstreamio/gitstream/blob/master/data_model.md
*/
type Message struct {

	//The contents of a message, mostly markdown
	Body string `json:"body"`

	//The git username of the message-author
	Author string `json:"author"`

	//The Timeline where the message was posted (e.g. "someuser/someproject" or just "someuser")
	Timeline string `json:"timeline"`

	//The messages' creation date
	Created time.Time `json:"created,omitempty"`

	//The messages' last update date
	Updated time.Time `json:"updated,omitempty"`

	//The direct link to this message
	Permalink string `json:"permalink,omitempty"`
}
