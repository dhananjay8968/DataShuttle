package models

// Message represents the data structure for a message.
type Message struct {
	Topic  string
	Key    string
	Value  string
	Offset int64
}
