package limiter

type LimitConfig struct {
	RequestsPerSecond int
	BlockDurationSec  int
}
