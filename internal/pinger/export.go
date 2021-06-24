package pinger

import (
	"context"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func export(project string, location string, pod string, ip string, value int64) error {
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	dataPoint := &monitoringpb.Point{
		Interval: &monitoringpb.TimeInterval{
			EndTime: &googlepb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
		Value: &monitoringpb.TypedValue{
			Value: &monitoringpb.TypedValue_Int64Value{
				Int64Value: value,
			},
		},
	}
	err = client.CreateTimeSeries(ctx, &monitoringpb.CreateTimeSeriesRequest{
		Name: monitoring.MetricProjectPath(project),
		TimeSeries: []*monitoringpb.TimeSeries{
			{
				Metric: &metricpb.Metric{
					Type: "custom.googleapis.com/pinger",
				},
				Resource: &monitoredrespb.MonitoredResource{
					Type: "generic_node",
					Labels: map[string]string{
						"project_id": project,
						"location":   location,
						"namespace":  pod,
						"node_id":    ip,
					},
				},
				Points: []*monitoringpb.Point{
					dataPoint,
				},
			},
		},
	})
	if err != nil {
		return err
	}
	err = client.Close()
	if err != nil {
		return err
	}
	return nil
}
