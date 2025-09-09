# Cardano Node Monitor

Hey there! This is a monitoring tool I built for Cardano blockchain nodes. If you're running a Cardano node and want to keep an eye on its health and performance, this tool is for you.

## What does it do?

Think of this as a health checkup system for your Cardano node. It watches your node 24/7 and tells you:

- Is my node healthy and syncing properly?
- What's the current block height?
- How fast is my node responding?
- Are there any errors I should worry about?
- How long has my node been running?

## Why I built this

I created this project to demonstrate my Go skills and understanding of blockchain infrastructure. It shows how to build production-ready monitoring systems that companies like Blink Labs use to keep their Cardano infrastructure running smoothly.

## Quick start

### Running with Docker (easiest way)

```bash
git clone https://github.com/yourusername/cardano-node-monitor
cd cardano-node-monitor
docker-compose up -d
```

Then open http://localhost:8080 in your browser.

### Running directly

If you have Go installed:

```bash
go build -o cardano-monitor cmd/main.go
./cardano-monitor --help
```

## What you can check

Once it's running, you can visit these URLs:

- **http://localhost:8080** - Main API info
- **http://localhost:8080/api/v1/health** - Quick health check
- **http://localhost:8080/api/v1/status** - Detailed node status
- **http://localhost:8080/metrics** - Prometheus metrics (for advanced monitoring)

## Configuration

You can customize how it runs:

```bash
./cardano-monitor --node-url http://your-cardano-node:12798 --port 8080 --interval 30s
```

Or set environment variables:
```bash
export CARDANO_MONITOR_NODE_URL=http://localhost:12798
export CARDANO_MONITOR_PORT=8080
export CARDANO_MONITOR_INTERVAL=30s
```

## Monitoring dashboard

I've included a Grafana dashboard (`grafana-dashboard.json`) that creates beautiful charts from the metrics. Just import it into Grafana and you'll get:

- Health status indicators
- Sync progress gauges  
- Block height charts
- Response time graphs
- Error tracking

## What's inside

The code is organized like this:

```
├── cmd/main.go              # Main application
├── internal/
│   ├── api/                 # REST API handlers
│   ├── config/              # Configuration management
│   └── monitor/             # Core monitoring logic
├── pkg/
│   ├── client/              # Cardano node client
│   └── metrics/             # Prometheus metrics
├── docker-compose.yml       # Easy deployment
└── grafana-dashboard.json   # Ready-to-use dashboard
```

## Technical details

Built with:
- **Go 1.21** - Main language
- **Gin** - Web framework for the API
- **Prometheus** - Metrics collection
- **Docker** - Containerization
- **Grafana** - Visualization (dashboard included)

The app uses goroutines for concurrent monitoring and provides a clean REST API for integration with other tools.

## Why this matters

This project demonstrates several key skills:

1. **Go development** - Concurrent applications, proper error handling, clean code structure
2. **Blockchain infrastructure** - Understanding of how to monitor node health and performance
3. **API design** - RESTful endpoints with proper JSON responses
4. **Monitoring** - Prometheus metrics, Grafana dashboards, production observability
5. **DevOps** - Docker deployment, CI/CD pipelines, configuration management

## Contributing

Feel free to open issues or submit pull requests. I'm always looking to improve the code and add new features.

## Contact

Built by Hari Babu Tirumani - tirumaniharibabu575@gmail.com

This project showcases my skills for blockchain infrastructure roles, particularly at companies like Blink Labs that focus on Cardano ecosystem tooling.