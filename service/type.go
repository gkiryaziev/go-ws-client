package service

// map for topics pool
type TopicPool map[string]func(string) []byte
