package service

// TopicPool map for topics pool
type TopicPool map[string]func(string) []byte
