# E3DC-Prometheus-Exporter
A prometheus exporter for [E3/DC](https://www.e3dc.com/en/) solar power management system

## Usage
You will have to create a ```config.toml``` file next to the binary for configuring the exporters connection to the [E3/DC](https://www.e3dc.com/en/) System.  
The layout is as follows

```toml
[ExporterConfig]
loglevel = "ERROR"

[e3dc]
address = "172.0.0.1"
username = "e3dc_user"
password = "supersafepassword"
key = "Sup3rS4v3K3y"
```

The ```loglevel``` will determine how talkative the exporter is. By default, it logs to the console.
