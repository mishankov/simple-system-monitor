# Simple System Monitor

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/mishankov/simple-system-monitor/ci.yml)](https://github.com/mishankov/simple-system-monitor/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/mishankov/simple-system-monitor/graph/badge.svg?token=SACQUDDX8R)](https://codecov.io/gh/mishankov/simple-system-monitor)
[![GitHub Release](https://img.shields.io/github/v/release/mishankov/simple-system-monitor?display_name=tag&label=latest%20release)](https://github.com/mishankov/simple-system-monitor/releases/latest)
[![GitHub License](https://img.shields.io/github/license/mishankov/simple-system-monitor)](https://github.com/mishankov/simple-system-monitor/blob/main/LICENSE)


![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mishankov/simple-system-monitor)
![GitHub package.json dev/peer/optional dependency version](https://img.shields.io/github/package-json/dependency-version/mishankov/simple-system-monitor/dev/svelte?filename=webapp%2Fpackage.json)

---

See your server resourses on simple web UI

<p align="middle">
  <img src="docs/images/theme-default.png" width="32%"/>
  <img src="docs/images/theme-ayu.png" width="32%"/>
  <img src="docs/images/theme-solarized-light.png" width="32%"/>
</p>

## Setup

### Docker Compose

Simple Docker Compose service looks like this

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
    - ./user-assets/:/app/user-assets # mount user assets to container
  restart: 'unless-stopped'
```

### Binary

Download latest `simple-system-monitor` binary from [latest GitHub release](https://github.com/mishankov/simple-system-monitor/releases/latest), unzip it and run

```shell
wget http://github.com/mishankov/simple-system-monitor/releases/latest/download/simple-server-monitor.zip
unzip simple-server-monitor.zip
./simple-server-monitor
```

## Configuration

Configuration of `simple-system-monitor` is done with environment variables. Available env vars:

- `SSM_PERIOD` - period for updating all monitoring in seconds. Default is `2`
- `SSM_PATH` - path to take system information from. Default is `/proc`
- `SSM_PORT` - port to run web server. Default is `4442`
- `SSM_CPUINFO_PERIOD` - period for updating CPU monitoring in seconds. Default is `SSM_PERIOD` value
- `SSM_CPUINFO_PATH` - path to take CPU information from. Default is `SSM_PATH` value + `/stat`
- `SSM_MEMINFO_PERIOD` - period for updating RAM monitoring in seconds. Default is `SSM_PERIOD` value
- `SSM_MEMINFO_PATH` - path to take RAM information from. Default is `SSM_PATH` value + `/meminfo`
- `SSM_UPTIME_PERIOD` - period for updating uptime monitoring in seconds. Default is `SSM_PERIOD` value
- `SSM_UPTIME_PATH` - path to take uptime information from. Default is `SSM_PATH` value + `/uptime`
- `SSM_USER_ASSETS_PATH` - path to user defined assets like custom CSS. Default is `./user-assets`

## Available pages

- `/` - page with all available metrics
- `/metrics/cpuinfo`, `/metrics/meminfo`, `/metrics/uptime` - dedicated pages for every type of monitoring. Can be used by including in other pages via iframe

## Theming

Theming available by supplying custom CSS file `user.css` at `SSM_USER_ASSETS_PATH`. A few CSS variables are available. See full list [here](./webapp/static/vars.css).

See premade themes [here](./docs/themes.md)
