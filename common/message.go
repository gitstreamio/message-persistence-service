package common

type Message struct {
	Id           string `json:"id"`
	User         string `json:"user"`
	Message      string `json:"message"`
	Organization string `json:"organization"`
	Project      string `json:"project"`
}
