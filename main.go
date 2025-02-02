package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	cfg     config
	mClient mqtt.Client
	wg      sync.WaitGroup
)

func publishMqttMessage(param string, msg interface{}) error {
	if mClient == nil {
		log.Println("mqtt client not initialised")
		return fmt.Errorf("mqtt client not initialised")
	}

	if !mClient.IsConnected() {
		log.Println("mqtt client not connected")
		return fmt.Errorf("mqtt client not connected")
	}

	token := mClient.Publish(fmt.Sprintf("%s/%s", cfg.MQTTPrefix, param), 0, cfg.RetainValues, msg)

	token.WaitTimeout(5 * time.Second)
	return token.Error()
}

func main() {
	time.LoadLocation("UTC") // WS90 reports time in UTC

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environmental variables: %s", err.Error())
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.MQTTHost, cfg.MQTTPort))
	opts.SetUsername(cfg.MQTTUser)
	opts.SetPassword(cfg.MQTTPass)
	mClient = mqtt.NewClient(opts)

	if token := mClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	r := gin.Default()
	r.POST(cfg.EndpointPath, func(c *gin.Context) {
		wg.Add(1)
		var m msg
		var err error

		c.Bind(&m)

		// Date
		m.Time, err = time.Parse("2006-01-02 15:04:05", c.PostForm("dateutc"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		m.convertUnits()

		go func() {
			publishMqttMessage("passkey", m.Passkey)
			publishMqttMessage("stationType", m.StationType)
			publishMqttMessage("runtime", fmt.Sprintf("%d", m.Runtime))
			publishMqttMessage("heap", fmt.Sprintf("%d", m.Heap))
			publishMqttMessage("time", m.Time.Format("2006-01-02 15:04:05"))
			publishMqttMessage("tempInC", fmt.Sprintf("%2.1f", m.TempInC))
			publishMqttMessage("humidityIn", fmt.Sprintf("%d", m.HumIn))
			publishMqttMessage("baromRelHpa", fmt.Sprintf("%4.1f", m.BaromRelHpa))
			publishMqttMessage("baromAbsHpa", fmt.Sprintf("%4.1f", m.BaromAbsHpa))
			publishMqttMessage("tempOutC", fmt.Sprintf("%2.1f", m.TempOutC))
			publishMqttMessage("humidityOut", fmt.Sprintf("%d", m.HumOut))
			publishMqttMessage("windDir", fmt.Sprintf("%d", m.WindDir))
			publishMqttMessage("windSpdMps", fmt.Sprintf("%2.1f", m.WindSpdMps))
			publishMqttMessage("windGustMps", fmt.Sprintf("%2.1f", m.WindGustMps))
			publishMqttMessage("maxDailyGustMps", fmt.Sprintf("%2.1f", m.MaxDailyGustMps))
			publishMqttMessage("solarRadiation", fmt.Sprintf("%3.2f", m.SolarRadiation))
			publishMqttMessage("rainRealTime", fmt.Sprintf("%3.3f", m.RainRealTime))
			publishMqttMessage("rainEvent", fmt.Sprintf("%3.3f", m.RainEvent))
			publishMqttMessage("rainHourly", fmt.Sprintf("%3.3f", m.RainHourly))
			publishMqttMessage("rainDaily", fmt.Sprintf("%3.3f", m.RainDaily))
			publishMqttMessage("rainWeekly", fmt.Sprintf("%3.3f", m.RainWeekly))
			publishMqttMessage("rainMonthly", fmt.Sprintf("%3.3f", m.RainMonthly))
			publishMqttMessage("rainYearly", fmt.Sprintf("%3.3f", m.RainYearly))
			publishMqttMessage("rainSeason", fmt.Sprintf("%3.3f", m.RainSeason))
			publishMqttMessage("capacVolt", fmt.Sprintf("%1.2f", m.CapacVoltage))
			publishMqttMessage("wh90Version", fmt.Sprintf("%d", m.WH90Version))
			publishMqttMessage("wh90Battery", fmt.Sprintf("%1.2f", m.WH90Battery))
			publishMqttMessage("frequency", m.Frequency)
			publishMqttMessage("model", m.Model)
			publishMqttMessage("updateInterval", fmt.Sprintf("%d", m.UpdateInterval))
		}()
	})

	r.Run(fmt.Sprintf(":%d", cfg.HTTPPort))

	<-ctx.Done()
	wg.Wait()

	mClient.Disconnect(500)

	os.Exit(0)
}
