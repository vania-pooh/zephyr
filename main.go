package main

import (
	"flag"
	"github.com/aerokube/zephyr/core"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
	"sync"
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
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	stop := make(chan bool)
	var wg sync.WaitGroup
	for _, transfer := range *config {
		processTransfer(transfer, stop, &wg)
	}
	go func() {
		<-s
		close(stop)
	}()
	wg.Wait()
}

func processTransfer(transfer core.Transfer, stop chan bool, wg *sync.WaitGroup) {
	readerSettings := transfer.ReaderSettings
	reader, delay := configureReader(readerSettings)
	writerSettings := transfer.WriterSettings
	writer := configureWriter(writerSettings)

	data := make(chan *core.Data, math.MaxInt32)

	ticker := time.NewTicker(delay)
	wg.Add(2)
	log.Printf("Initializing transfer from [%s] to [%s]\n", readerSettings.Name, writerSettings.Name)
	go func() {
		defer wg.Done()
		readerName := readerSettings.Name
		readOnce(readerName, reader, &data)
		for {
			select {
			case <-ticker.C: readOnce(readerName, reader, &data)
			case <-stop:
				{
					log.Printf("Stopping reader [%s]\n", readerSettings.Name)
					ticker.Stop()
					return
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case dt := <-data:
				{
					err := writer.Write(dt)
					if err != nil {
						log.Printf("Failed to write with %s: %v", writerSettings.Name, err)
					}
				}
			case <-stop:
				log.Printf("Stopping writer [%s]\n", writerSettings.Name)
				return
			}
		}
	}()
}

func readOnce(readerName string, reader core.Reader, data * chan *core.Data) {
	dt, err := reader.Read()
	if err != nil {
		log.Printf("Failed to read with %s: %v", readerName, err)
	} else {
		*data <- dt
	}
}

func configureReader(settings core.ReaderSettings) (core.Reader, time.Duration) {
	reader, err := GetReader(settings.Name)
	dieOnError(err)
	err = reader.Configure(settings)
	dieOnError(err)
	delay, err := time.ParseDuration(settings.Delay)
	dieOnError(err)
	return reader, delay
}

func configureWriter(settings core.WriterSettings) core.Writer {
	writer, err := GetWriter(settings.Name)
	dieOnError(err)
	err = writer.Configure(settings)
	dieOnError(err)
	return writer
}

func dieOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
