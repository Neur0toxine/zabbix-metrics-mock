# Zabbix Metrics Mock

[![buildx status](https://img.shields.io/github/actions/workflow/status/Neur0toxine/zabbix-metrics-mock/buildx.yml?branch=master&style=for-the-badge)](https://github.com/Neur0toxine/zabbix-metrics-mock/actions?query=workflow%3Abuildx)
[![Docker Pulls](https://img.shields.io/docker/pulls/neur0toxine/zabbix-metrics-mock?style=for-the-badge)](https://hub.docker.com/r/neur0toxine/zabbix-metrics-mock)

This app will print out any Zabbix packet that contains metrics (`request: sender data`). It will bind itself to port 10051 by default. 
You can change the port by using `LISTEN` environment variable like this: `LISTEN=:8080`.

## Usage

1. Run the app.
2. Send the metrics to the `localhost:10051` (or `service_name:10051` in case of Docker).
3. You will see something like this then the packet arrives:
```
packet type `sender data`, timestamp: 2022-12-28 07:27:24 +0000 UTC
 - [test] app.goroutine.count: 9
 - [test] app.mem.alloc: 4850696
 - [test] app.mem.sys: 24099856
end of packet
```

You can use `neur0toxine/zabbix-metrics-mock` Docker image.
