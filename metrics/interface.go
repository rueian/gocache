package metrics

import "github.com/eko/gache/codec"

// MetricsInterface represents the metrics interface for all available providers
type MetricsInterface interface {
	Record(store, metric string, value float64)
	RecordFromCodec(codec codec.CodecInterface)
}