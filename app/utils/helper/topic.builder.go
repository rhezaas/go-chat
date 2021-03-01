package helper

import (
	"log"

	"go-chat/app/utils/types"
)

// TopicBuilder ...
func TopicBuilder(topic string, params ...interface{}) string {
	if params != nil {
		switch params[0].(type) {
		case string:
			return ifString(topic, params)
		case types.TopicParams:
			return ifMap(topic, params)
		default:
			log.Fatal("Params can only get string or types.TopicParams")
			return ""
		}
	} else {
		return topic
	}
}

func ifString(topic string, params []interface{}) string {
	newTopic := topic

	for i, param := range params {
		if i == 0 {
			newTopic += "?" + param.(string) + "=" + "{" + param.(string) + "}"
		} else {
			newTopic += "&" + param.(string) + "=" + "{" + param.(string) + "}"
		}
	}

	return newTopic
}

func ifMap(topic string, params []interface{}) string {
	i := 0
	newTopic := topic

	for param, value := range params[0].(types.TopicParams) {
		if i == 0 {
			newTopic += "?" + param + "=" + value
		} else {
			newTopic += "&" + param + "=" + value
		}

		i++
	}

	return newTopic
}
