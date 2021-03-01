package types

// TopicRoute ...
type TopicRoute map[string]func(TopicParams, string)

// TopicParams ...
type TopicParams map[string]string
