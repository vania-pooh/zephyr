package main

import (
	"flag"
	"github.com/aerokube/zephyr/core"
	"log"
	"os"
	"os/signal"
	"syscall"
	"math"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config.json", "path to config.json file")
	flag.Parse()
}

func main() {
	config, err := core.LoadConfig(configFile)
	dieOnError(err)
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	for _, transfer := range *config {
		readerSettings := transfer.ReaderSettings
		reader := configureReader(readerSettings)
		writerSettings := transfer.WriterSettings
		writer := configureWriter(writerSettings)
		
		data := make(chan *core.Data, math.MaxInt32)
		
		go func() {
			select {
			case <-stop: return
			default:
			}
			dt, err := reader.Read()
			if (err != nil) {
				log.Printf("Failed to read with %s: %v", readerSettings.Name, err)
			}
			data <- dt
		}()
		go func() {
			for {
				select {
				case dt:= <-data: {
					err := writer.Write(dt)
					if (err != nil) {
						log.Printf("Failed to write with %s: %v", writerSettings.Name, err)
					}
				}
				case <-stop: return
				}
			}
		}()
	}
}

func configureReader(settings core.Settings) core.Reader {
	reader, err := core.GetReader(settings.Name)
	dieOnError(err)
	err = reader.Configure(settings)
	dieOnError(err)
	return reader
}

func configureWriter(settings core.Settings) core.Writer {
	writer, err := core.GetWriter(settings.Name)
	dieOnError(err)
	err = writer.Configure(settings)
	dieOnError(err)
	return writer
}

func dieOnError(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}

func configure() {
}