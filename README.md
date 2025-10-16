[![Build Status](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/actions/workflows/test.yml/badge.svg)](https://github.com/warrensbox/terraform-switcher/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/MatrixCrawler/E3DC-Prometheus-Exporter)](https://goreportcard.com/report/github.com/MatrixCrawler/E3DC-Prometheus-Exporter)
[![CodeQL](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/workflows/CodeQL/badge.svg)](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/actions?query=workflow%3A%22CodeQL%22)
![GitHub Release](https://img.shields.io/github/v/release/MatrixCrawler/E3DC-Prometheus-Exporter)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/MatrixCrawler/E3DC-Prometheus-Exporter/total)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/MatrixCrawler/E3DC-Prometheus-Exporter)

# E3DC Prometheus Exporter

A Prometheus exporter for [E3/DC](https://www.e3dc.com/en/) solar power management systems. This exporter connects to your E3/DC system via the RSCP protocol and exposes metrics in Prometheus format for monitoring and visualization.

## Features

- 📊 Exports E3/DC system metrics to Prometheus
- 🔐 Secure connection using RSCP protocol
- 📝 Configurable logging (console or file output)
- 🎯 Lightweight and efficient
- ⚙️ Easy configuration via YAML

## Prerequisites

- Go 1.25.0 or higher (for building from source)
- Access to an E3/DC system with valid credentials
- Network connectivity to your E3/DC device

## Installation

### From Release

Download the latest binary for your platform from the [releases page](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/releases).

## Configuration
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

### Configuration Parameters

#### Exporter Configuration
- **port**: The port on which the exporter will listen (default: 10998)
- **log.level**: Logging verbosity level
    - `DEBUG`: Detailed information for debugging
    - `INFO`: General informational messages
    - `WARN`: Warning messages
    - `ERROR`: Error messages only
    - `FATAL`: Only critical errors
- **log.output**: Where to send logs (`stdout` for console, `file` for file output)
- **log.file**: Path to the log file (required when output is `file`)

#### E3/DC Configuration
- **address**: IP address of your E3/DC system
- **username**: Username for E3/DC portal authentication
- **password**: Password for E3/DC portal authentication
- **key**: RSCP encryption key (can be obtained from the E3/DC portal)

## Usage

1. Create your `config.yml` file with the appropriate settings
2. Run the exporter:
3. The exporter will start and listen on the configured port (default: 10998)
4. Access the metrics at: `http://localhost:10998/metrics`


### Prometheus Configuration

Add the following to your `prometheus.yml`:

```yaml
scrape_configs:
  job_name: 'e3dc'
  static_configs:
    - targets: ['localhost:10998']
```

## Endpoints

- `/` - Landing page with link to metrics
- `/metrics` - Prometheus metrics endpoint

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [E3/DC](https://www.e3dc.com/en/) for their solar power management systems
- [spali/go-rscp](https://github.com/spali/go-rscp) for the RSCP protocol implementation

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/MatrixCrawler/E3DC-Prometheus-Exporter/issues) on GitHub.
