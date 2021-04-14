package statscollector

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/swathins079/chargemywifi/pkg/metric"
	"github.com/swathins079/chargemywifi/pkg/model"
)

type StatsCollector interface {
	Notifier
	Get()
}

type Notifier interface {
	Register(metric.Metric)
	Deregister(metric.Metric)
	Notify()
}

type wifiStats struct {
	stats map[metric.Metric]struct{}
}

func New() StatsCollector {
	return &wifiStats{stats: make(map[metric.Metric]struct{})}
}

func (ws *wifiStats) Get() {
	resp, err := http.Get(model.JIOFIURL)
	checkError(err)
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	htmlSource := string(html)
	for metric := range ws.stats {
		metric.Value(metric.Rgx().FindStringSubmatch(htmlSource)[1])
	}
}

func (ws *wifiStats) Register(data metric.Metric) {
	ws.stats[data] = struct{}{}
}

func (ws *wifiStats) Deregister(data metric.Metric) {
	delete(ws.stats, data)
}

func (ws *wifiStats) Notify() {
	for metric := range ws.stats {
		metric.OnNotify()
	}
}

// checkError is utility function which checks for error.
func checkError(err error) {
	if err != nil {
		log.Fatal("Failure: ", err)
	}
}