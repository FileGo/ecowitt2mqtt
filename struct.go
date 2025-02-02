package main

import "time"

type msg struct {
	Passkey         string `form:"PASSKEY"`
	StationType     string `form:"stationtype"`
	Runtime         int64  `form:"runtime"`
	Heap            int64  `form:"heap"`
	Time            time.Time
	TempInF         float64 `form:"tempinf"`
	TempInC         float64
	HumIn           int     `form:"humidityin"`
	BaromRelIn      float64 `form:"baromrelin"`
	BaromRelHpa     float64
	BaromAbsIn      float64 `form:"baromabsin"`
	BaromAbsHpa     float64
	TempOutF        float64 `form:"tempf"`
	TempOutC        float64
	HumOut          int     `form:"humidity"`
	WindDir         int     `form:"winddir"`
	WindSpdMph      float64 `form:"windspeedmph"`
	WindSpdMps      float64
	WindGustMph     float64 `form:"windgustmph"`
	WindGustMps     float64
	MaxDailyGustMph float64 `form:"maxdailygust"`
	MaxDailyGustMps float64
	SolarRadiation  float64 `form:"solarradiation"`
	RainRealTime    float64 `form:"rrain_piezo"`
	RainEvent       float64 `form:"erain_piezo"`
	RainHourly      float64 `form:"hrain_piezo"`
	RainDaily       float64 `form:"drain_piezo"`
	RainWeekly      float64 `form:"wrain_piezo"`
	RainMonthly     float64 `form:"mrain_piezo"`
	RainYearly      float64 `form:"yrain_piezo"`
	RainSeason      float64 `form:"srain_piezo"`
	CapacVoltage    float64 `form:"ws90cap_volt"`
	WH90Version     int     `form:"ws90_ver"`
	WH90Battery     float64 `form:"wh90batt"`
	Frequency       string  `form:"freq"`
	Model           string  `form:"model"`
	UpdateInterval  int     `form:"interval"`
}

func (m *msg) convertUnits() {
	m.TempInC = FtoC(m.TempInF)
	m.TempOutC = FtoC(m.TempOutF)

	m.BaromRelHpa = InToHpa(m.BaromRelIn)
	m.BaromAbsHpa = InToHpa(m.BaromAbsIn)

	m.WindSpdMps = MphToMps(m.WindSpdMph)
	m.WindGustMps = MphToMps(m.WindGustMph)
	m.MaxDailyGustMps = MphToMps(m.MaxDailyGustMph)
}

type config struct {
	MQTTPrefix   string `env:"MQTT_PREFIX,notEmpty"`
	MQTTHost     string `env:"MQTT_HOST,notEmpty"`
	MQTTPort     int    `env:"MQTT_PORT" envDefault:"1883"`
	MQTTUser     string `env:"MQTT_USERNAME,notEmpty"`
	MQTTPass     string `env:"MQTT_PASSWORD"`
	EndpointPath string `env:"ENDPOINT_PATH" envDefault:"/"`
	HTTPPort     int    `env:"HTTP_PORT" envDefault:"55904"`
	RetainValues bool   `env:"RETAIN_VALUES" envDefault:"true"`
}
