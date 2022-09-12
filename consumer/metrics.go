package consumer

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Measures for the stats quickstart.
var (
	// The latency in milliseconds
	mLatencyMs = stats.Float64("consumer/latency", "The latency in milliseconds per message", stats.UnitMilliseconds)

	// Counts the number of lines read in from standard input
	mMessagesIn = stats.Int64("consumer/messages_in", "The number of messages consumed", stats.UnitNone)

	// Encounters the number of non EOF(end-of-file) errors.
	mErrors = stats.Int64("consumer/errors", "The number of errors encountered", stats.UnitNone)

	// Counts/groups the lengths of lines read in.
	mMessageSizes = stats.Int64("consumer/messsage_sizes", "The distribution of message sizes", stats.UnitBytes)
)

// TagKeys for the stats quickstart.
var (
	keyProcessor, _ = tag.NewKey("processor")
	keyTopic, _     = tag.NewKey("topic")
)

// Views for the stats quickstart.
var (
	latencyView = &view.View{
		Name:        "consumer/latency",
		Measure:     mLatencyMs,
		Description: "The distribution of the latencies",

		// Latency in buckets:
		// [>=0ms, >=25ms, >=50ms, >=75ms, >=100ms, >=200ms, >=400ms, >=600ms, >=800ms, >=1s, >=2s, >=4s, >=6s]
		Aggregation: view.Distribution(25, 50, 75, 100, 200, 400, 600, 800, 1000, 2000, 4000, 6000),
		TagKeys: []tag.Key{
			keyProcessor,
			keyTopic,
		},
	}

	messageCountView = &view.View{
		Name:        "consumer/messages_in",
		Measure:     mMessagesIn,
		Description: "The number of lines from standard input",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			keyProcessor,
			keyTopic,
		},
	}

	errorCountView = &view.View{
		Name:        "consumer/errors",
		Measure:     mErrors,
		Description: "The number of errors encountered",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			keyProcessor,
			keyTopic,
		},
	}

	messageSizeView = &view.View{
		Name:        "consumer/message_sizes",
		Description: "Groups the lengths of keys in buckets",
		Measure:     mMessageSizes,
		// Lengths: [>=0B, >=5B, >=10B, >=15B, >=20B, >=40B, >=60B, >=80, >=100B, >=200B, >=400, >=600, >=800, >=1000]
		Aggregation: view.Distribution(5, 10, 15, 20, 40, 60, 80, 100, 200, 400, 600, 800, 1000),
		TagKeys: []tag.Key{
			keyProcessor,
			keyTopic,
		},
	}

	OpenCensusViews = []*view.View{
		latencyView,
		messageCountView,
		errorCountView,
		messageSizeView,
	}
)

func recordMetrics(ctx context.Context, data []byte) func(err error) {
	startTime := time.Now()
	return func(err error) {
		ms := float64(time.Since(startTime).Nanoseconds()) / 1e6
		stats.Record(ctx,
			mLatencyMs.M(ms),
			mMessagesIn.M(1),
			mMessageSizes.M(int64(len(data))),
		)
		if err != nil {
			stats.Record(ctx, mErrors.M(1))
		}
	}
}
