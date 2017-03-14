package writer

import (
	"github.com/aerokube/zephyr/core"
	"github.com/marpaia/graphite-golang"
	"strconv"
	"log"
)

const (
	GraphiteHost = "host"
	GraphitePort = "port"
)

type GraphiteWriter struct {
	graphiteClient *graphite.Graphite
}

func (gw *GraphiteWriter) Write(data *core.Data) error {
	for _, metric := range *data {
		log.Printf("Sending metric [%s = %s] to Graphite at [%s:%s]\n", metric.Key, metric.Value, gw.graphiteClient.Host, gw.graphiteClient.Port)
		err := gw.graphiteClient.SimpleSend(metric.Key, metric.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (gw *GraphiteWriter) Configure(settings core.WriterSettings) error {
	host, err := settings.GetProperty(GraphiteHost)
	if err != nil {
		return err
	}
	port, err := settings.GetProperty(GraphitePort)
	if err != nil {
		return err
	}
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	gc, err := graphite.NewGraphite(host, intPort)
	if err != nil {
		return err
	}
	gw.graphiteClient = gc
	return nil
}
