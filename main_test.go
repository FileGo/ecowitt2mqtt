package main

import (
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
)

func TestPublishMqttMessage(t *testing.T) {
	t.Run("fail - not initialised", func(t *testing.T) {
		assert := assert.New(t)

		mClient = nil
		err := publishMqttMessage("ws90", "testvalue")

		if assert.Error(err) {
			assert.Contains(err.Error(), "initialised")
		}
	})

	t.Run("fail - not connected", func(t *testing.T) {
		assert := assert.New(t)

		mClient = mqtt.NewClient(&mqtt.ClientOptions{})
		err := publishMqttMessage("ws90", "testvalue")

		if assert.Error(err) {
			assert.Contains(err.Error(), "connected")
		}
	})
}
