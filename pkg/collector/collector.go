package collector

import (
	"fmt"
	
	"github.com/nirmata/kube-netc/pkg/tracker"
	"github.com/nirmata/kube-netc/pkg/cluster"
	"github.com/prometheus/client_golang/prometheus"
)

func StartCollector(tr *tracker.Tracker, ci *cluster.ClusterInfo) {
	for {
		select {
		case update := <-tr.NodeUpdateChan:
			ActiveConnections.Set(float64(update.NumConnections))

		case update := <-tr.ConnUpdateChan:
			id := prometheus.Labels{"id": tracker.FormatCID(update.Connection)}
			BytesSent.With(id).Set(float64(update.Data.BytesSent))
			BytesRecv.With(id).Set(float64(update.Data.BytesRecv))
			BytesSentPerSecond.With(id).Set(float64(update.Data.BytesSentPerSecond))
			BytesRecvPerSecond.With(id).Set(float64(update.Data.BytesRecvPerSecond))
			
		case update := <-ci.MapUpdateChan:
			
			if update.Info != nil {
				fmt.Printf("Pod Update: %s -> %s\n", update.IP, update.Info.Name)
			} else {
				fmt.Printf("Pod Deleted: %s\n", update.IP)
			}
			
		}
	}
}
