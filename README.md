# Ecowitt to MQTT Bridge

Primarily designed to feed Ecowitt WS90 weather station data to an MQTT server, from which it can be used in any other source, e.g. Home Assistant through [MQTT Integration](https://www.home-assistant.io/integrations/mqtt/).

You will need an MQTT broker to use this software, either a standalone one like [Mosquitto](https://mosquitto.org/) or run it as a [Home Assistant Add-on](https://www.home-assistant.io/addons/).

The easiest way to run is via Docker Compose - see [example](examples/docker-compose.yaml). You will need to change relevant environment variables, and the exposed port as well.

Place the compose file in a new folder and run the following commands:

```shell
docker compose build
docker compose up -d
```

You will then need to add sensors in `configuration.yaml` file in Home Assistant. See [example](examples/ha-config.yaml).

## Known issues
1. TLS connections to MQTT broker currently aren't supported.
2. The software was only tested with Ecowitt WS90 station with a GW1100A Wi-Fi bridge.