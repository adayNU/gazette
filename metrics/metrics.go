package metrics

import "github.com/prometheus/client_golang/prometheus"

// Keys for gazette metrics.
const (
	Fail = "fail"
	Ok   = "ok"
)

// Collectors for Gazette broker & consumer metrics.
var (
	AllocatorAssignmentAddedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_allocator_assignment_added_total",
		Help: "Cumulative number of item / member assignments added by the allocator.",
	})
	AllocatorAssignmentPackedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_allocator_assignment_packed_total",
		Help: "Cumulative number of item / member assignments packed by the allocator.",
	})
	AllocatorAssignmentRemovedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_allocator_assignment_removed_total",
		Help: "Cumulative number of item / member assignments removed by the allocator.",
	})
	AllocatorConvergeTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_allocator_converge_total",
		Help: "Cumulative number of converge iterations.",
	})
	AllocatorMaxFlowRuntimeSeconds = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "gazette_allocator_max_flow_runtime_seconds",
		Help: "Duration required to re-solve for maximum assignment.",
	})
	AllocatorNumItemSlots = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gazette_allocator_desired_replication_slots",
		Help: "Number of desired item replication slots summed across all items.",
	})
	AllocatorNumItems = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gazette_allocator_items",
		Help: "Number of items known to the allocator.",
	})
	AllocatorNumMembers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gazette_allocator_members",
		Help: "Number of members known to the allocator.",
	})
	CommittedBytesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_committed_bytes_total",
		Help: "Cumulative number of bytes committed to journals.",
	})
	CommitsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gazette_commits_total",
		Help: "Cumulative number of commits.",
	}, []string{"status"})
	RecoveryLogRecoveredBytesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_recoverylog_recovered_bytes_total",
		Help: "Cumulative number of bytes recovered.",
	})
	StoreRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gazette_store_requests_total",
		Help: "Cumulative number of fragment store operations.",
	}, []string{"provider", "operation", "status"})
	StorePersistedBytesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gazette_store_persisted_bytes_total",
		Help: "Cumulative number of bytes persisted to fragment stores.",
	}, []string{"provider"})
	JournalServerResponseTimeSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "gazette_journal_server_response_time_seconds",
		Help: "Response time of JournalServer.Append.",
	}, []string{"operation", "status"})
	WriteHead = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gazette_write_head",
		Help: "Current write head.",
	}, []string{"journal"})
)

// GazetteBrokerCollectors lists collectors used by the gazette broker.
func GazetteBrokerCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		AllocatorAssignmentAddedTotal,
		AllocatorAssignmentPackedTotal,
		AllocatorAssignmentRemovedTotal,
		AllocatorConvergeTotal,
		AllocatorMaxFlowRuntimeSeconds,
		AllocatorNumItemSlots,
		AllocatorNumItems,
		AllocatorNumMembers,
		CommitsTotal,
		CommittedBytesTotal,
		JournalServerResponseTimeSeconds,
		StorePersistedBytesTotal,
		StoreRequestTotal,
		WriteHead,
	}
}

// Collectors for gazette.Client and gazette.WriteService metrics.
var (
	GazetteDiscardBytesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_discard_bytes_total",
		Help: "Cumulative number of bytes read and discarded during a fragment seek.",
	})
	GazetteReadBytesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_read_bytes_total",
		Help: "Cumulative number of bytes read.",
	})
	GazetteSequencerQueuedTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gazette_sequencer_queued",
		Help: "Cumulative number of read-uncommitted messages which were sequenced.",
	}, []string{"journal", "flag", "outcome"})
	GazetteSequencerReplayTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gazette_sequencer_replay",
		Help: "Cumulative number of messages re-read from source journal due to insufficient Sequencer ring-buffer size.",
	}, []string{"journal"})
	GazetteWriteBytesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_write_bytes_total",
		Help: "Cumulative number of bytes written.",
	})
	GazetteWriteCountTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_write_count_total",
		Help: "Cumulative number of writes.",
	})
	GazetteWriteDurationTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_write_duration_seconds_total",
		Help: "Cumulative number of seconds spent writing.",
	})
	GazetteWriteFailureTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_write_failure_total",
		Help: "Cumulative number of write errors returned to clients.",
	})
)

// GazetteClientCollectors returns the metrics used by gazette.Client and
// gazette.WriteService.
func GazetteClientCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		GazetteDiscardBytesTotal,
		GazetteReadBytesTotal,
		GazetteSequencerQueuedTotal,
		GazetteSequencerReplayTotal,
		GazetteWriteBytesTotal,
		GazetteWriteCountTotal,
		GazetteWriteDurationTotal,
	}
}

// Collectors for consumer.Runner metrics.
var (
	GazetteConsumerTxCountTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_count_total",
		Help: "Cumulative number of transactions",
	})
	GazetteConsumerTxMessagesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_messages_total",
		Help: "Cumulative number of messages.",
	})
	GazetteConsumerTxSecondsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_seconds_total",
		Help: "Cumulative number of seconds processing transactions.",
	})
	GazetteConsumerTxConsumeSecondsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_consume_seconds_total",
		Help: "Cumulative number of seconds transactions were processing messages.",
	})
	GazetteConsumerTxStalledSecondsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_stalled_seconds_total",
		Help: "Cumulative number of seconds transactions were stalled waiting for Gazette IO.",
	})
	GazetteConsumerTxFlushSecondsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_flush_seconds_total",
		Help: "Cumulative number of seconds transactions were flushing their commit.",
	})
	GazetteConsumerTxSyncSecondsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_tx_sync_seconds_total",
		Help: "Cumulative number of seconds transactions were waiting for their commit to sync.",
	})
	GazetteConsumerBytesConsumedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gazette_consumer_consumed_bytes_total",
		Help: "Cumulative number of bytes consumed.",
	})
	GazetteConsumerReadHead = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gazette_consumer_read_head",
		Help: "Consumer read head",
	}, []string{"journal"})
)

// GazetteConsumerCollectors returns the metrics used by the consumer package.
func GazetteConsumerCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		GazetteConsumerTxCountTotal,
		GazetteConsumerTxMessagesTotal,
		GazetteConsumerTxSecondsTotal,
		GazetteConsumerTxConsumeSecondsTotal,
		GazetteConsumerTxStalledSecondsTotal,
		GazetteConsumerTxFlushSecondsTotal,
		GazetteConsumerBytesConsumedTotal,
		GazetteConsumerReadHead,
	}
}
