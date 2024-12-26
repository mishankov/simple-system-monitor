# Simple System Monitor

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/mishankov/simple-system-monitor/ci.yml)](https://github.com/mishankov/simple-system-monitor/actions/workflows/ci.yml)
[![GitHub Release](https://img.shields.io/github/v/release/mishankov/simple-system-monitor?display_name=tag&label=latest%20release)](https://github.com/mishankov/simple-system-monitor/releases/latest)
[![GitHub License](https://img.shields.io/github/license/mishankov/simple-system-monitor)](https://github.com/mishankov/simple-system-monitor/blob/main/LICENSE)


![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mishankov/simple-system-monitor)
![GitHub package.json dev/peer/optional dependency version](https://img.shields.io/github/package-json/dependency-version/mishankov/simple-system-monitor/dev/svelte?filename=webapp%2Fpackage.json)

---

See your server resourses on simple web UI

![screenshor](images/screenshot.png)

## Setup

### docker compose

Simple docker compose service looks like this

```yaml
simple-system-monitor:
  image: ghcr.io/mishankov/simple-system-monitor:latest
  container_name: simple-system-monitor
  ports:
    - 4442:4442
  environment:
    - SSM_PATH="/hostfs/proc"
  volumes:
    - /:/hostfs:ro
  restart: 'unless-stopped'
```

### Binary

Download latest `simple-system-monitor` binary from [latest GitHub release](https://github.com/mishankov/simple-system-monitor/releases/latest), unzip it and run

```shell
wget http://github.com/mishankov/simple-system-monitor/releases/latest/download/simple-server-monitor.zip
unzip simple-server-monitor.zip
./simple-server-monitor
```



