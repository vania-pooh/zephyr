package reader

import (
	"encoding/json"
	"fmt"
	"github.com/aandryashin/selenoid/config"
	"github.com/aerokube/zephyr/core"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Selenoid = "selenoid"
	Hostname = "hostname"

	//Metric names
	Total   = "total"
	Used    = "used"
	Queued  = "queued"
	Pending = "pending"
)

type SelenoidReader struct {
	selenoid string
	hostname string
}

func (sr *SelenoidReader) Read() (*core.Data, error) {
	url := fmt.Sprintf("http://%s/status", Selenoid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Selenoid status code not 200: %d", resp.StatusCode)
	}
	var state config.State
	err = json.NewDecoder(resp.Body).Decode(&state)
	if err != nil {
		return nil, err
	}
	return sr.stateToData(&state), nil
}

func (sr *SelenoidReader) stateToData(state *config.State) *core.Data {
	data := core.Data{}
	data = append(
		data,
		sr.createMetric(Total, strconv.Itoa(state.Total)),
		sr.createMetric(Used, strconv.Itoa(state.Used)),
		sr.createMetric(Queued, strconv.Itoa(state.Queued)),
		sr.createMetric(Pending, strconv.Itoa(state.Pending)),
	)
	for browserName, version := range state.Browsers {
		for versionName, quota := range version {
			for quotaName, used := range quota {
				metricName := strings.Join([]string{browserName, versionName, quotaName}, ".")
				data = append(data, sr.createMetric(metricName, strconv.Itoa(used)))
			}
		}
	}
	return &data
}

func (sr *SelenoidReader) createMetric(name string, value string) core.Metric {
	return core.Metric{Key: sr.fullMetricName(name), Value: value, Timestamp: time.Now()}
}

func (sr *SelenoidReader) fullMetricName(name string) string {
	return fmt.Sprintf("one_min.%s.%s", sr.hostname, name)

}

func (sr *SelenoidReader) Configure(settings core.ReaderSettings) error {
	selenoid, err := settings.GetProperty(Selenoid)
	if err != nil {
		return err
	}
	sr.selenoid = selenoid
	hostname, err := settings.GetProperty(Hostname)
	if err != nil {
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "selenoid-host"
		}
	}
	sr.hostname = hostname
	return nil
}
