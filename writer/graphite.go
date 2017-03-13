package writer

import (
	"github.com/vania-pooh/selemon/core"
	"github.com/marpaia/graphite-golang"
	"strconv"
)

const (
	graphiteHost = "host"
	graphitePort = "port"
)

type graphiteWriter struct {
	graphiteClient graphite.Graphite
}

func init() {
	core.AddWriter("graphite", &graphiteWriter{})
}

func (gw *graphiteWriter) Write(data *core.Data) error {
	for _, metric := range *data {
		err := gw.graphiteClient.SimpleSend(metric.Key, metric.Value)
		if (err != nil) {
			return err
		}
	}
	return nil
}

func (gw *graphiteWriter) Configure(settings core.Settings) error {
	host, err := settings.GetProperty(graphiteHost)
	if (err != nil) {
		return err
	}
	port, err := settings.GetProperty(graphitePort)
	if (err != nil) {
		return err
	}
	intPort, err := strconv.Atoi(port)
	if (err != nil) {
		return err
	}
	gc, err := graphite.NewGraphite(host, intPort)
	if (err != nil) {
		return err
	}
	gw.graphiteClient = gc
	return nil
}