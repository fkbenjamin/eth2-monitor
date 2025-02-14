package opts

var (
	LogLevel       string
	BeaconNode     string
	MetricsPort    string
	SlackURL       string
	SlackUsername  string
	PushGatewayUrl string
	PushGatewayJob string

	Monitor struct {
		ReplayEpoch         []uint
		SinceEpoch          uint64
		PrintSuccessful     bool
		DistanceTolerance   uint64
		UseAbsoluteDistance bool

		Pubkeys []string
	}

	Slashings struct {
		ShowSlashingReward    bool
		TwitterConsumerKey    string
		TwitterConsumerSecret string
		TwitterAccessToken    string
		TwitterAccessSecret   string
	}

	Maintenance struct {
		Epoch uint64
	}
)
