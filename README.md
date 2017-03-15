# Zephyr
This repository contains a small tool to upload data from Aerokube products to popular monitoring software.

## Usage
Zephyr is a simple bridge between a set of **readers** and **writers**. A **reader** fetches data from one of Aerokube products. A **writer** saves this data to popular monitoring and visualisation software. To start Zephyr working you need to provide a simple JSON configuration file that determines a list of reader-writer pairs. For each reader and writer you specify its name and some properties. For example for Graphite writer you need to configure host and port to save data to. Each reader is called on schedule - this is why you need to also specify a delay between calls. Each writer is waiting for data from its reader.    
1. Create JSON config file (e.g. in ```/etc/zephyr/zephyr.json```):
```json
[
  {
    "reader": {
      "name": "selenoid",
      "delay": "60s",
      "properties": {
        "selenoid": "my-selenoid-host.example.com:4444"
      }
    },
    "writer": {
      "name": "graphite",
      "properties": {
        "host": "my-graphite-host.example.com",
        "port": "2024"
      }
    }
  }
]
```
2. Run Zephyr container:
```bash
# docker run -d --name zephyr -v /etc/zephyr:/:ro aerokube/zephyr:latest
```

## Building
```
 $ govendor sync
 $ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
 $ docker build -t zephyr:latest .
```