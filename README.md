[![Build Status](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/actions/workflows/test.yml/badge.svg)](https://github.com/warrensbox/terraform-switcher/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/MatrixCrawler/E3DC-Prometheus-Exporter)](https://goreportcard.com/report/github.com/MatrixCrawler/E3DC-Prometheus-Exporter)
![GitHub Release](https://img.shields.io/github/v/release/MatrixCrawler/E3DC-Prometheus-Exporter)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/MatrixCrawler/E3DC-Prometheus-Exporter)

# E3DC-Prometheus-Exporter
A prometheus exporter for [E3/DC](https://www.e3dc.com/en/) solar power management system

## Usage
You will have to create a ```config.yml``` file next to the binary for configuring the exporters connection to the [E3/DC](https://www.e3dc.com/en/) System.  
The layout is as follows

```yaml
---

exporterconfig:
  loglevel: DEBUG

e3dc:
  address: 172.0.0.1
  username: e3dc_user
  password: supersafepassword
  key: Sup3rS4v3K3y

```

The ```loglevel``` will determine how talkative the exporter is. By default, it logs to the console.
