[![Build Status](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/actions/workflows/test.yml/badge.svg)](https://github.com/warrensbox/terraform-switcher/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/MatrixCrawler/E3DC-Prometheus-Exporter)](https://goreportcard.com/report/github.com/MatrixCrawler/E3DC-Prometheus-Exporter)
[![CodeQL](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/workflows/CodeQL/badge.svg)](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/actions?query=workflow%3A%22CodeQL%22)
![GitHub Release](https://img.shields.io/github/v/release/MatrixCrawler/E3DC-Prometheus-Exporter)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/MatrixCrawler/E3DC-Prometheus-Exporter/total)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/MatrixCrawler/E3DC-Prometheus-Exporter)

# E3DC-Prometheus-Exporter
A prometheus exporter for [E3/DC](https://www.e3dc.com/en/) solar power management system.

## Usage
You will have to create a ```config.yml``` file next to the binary for configuring the exporters connection to the [E3/DC](https://www.e3dc.com/en/) System.  
The layout is as follows

```yaml
---

exporterconfig:
  port: 10998
  log:
    level: ERROR
    out: stdout  # or file if you want to log to a file
    file: PathToLogFile

e3dc:
  address: 172.0.0.1
  username: e3dc_user
  password: supersafepassword
  key: Sup3rS4v3K3y

```

The ```loglevel``` will determine how talkative the exporter is. By default, it logs to the console.
