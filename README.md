# Go Redis

## How to start

1. Run application
```
$ go run main.go
```

2. Start InfluxDB and Grafana and go to grafana dashboard at http://localhost:9001
- Add influxdb
- Add dashboard
```
$ docker compose up influxdb grafana
```

3. Run Redis and K6 Load test
```
$ docker compose run --rm k6 run /scripts/test.js
```

