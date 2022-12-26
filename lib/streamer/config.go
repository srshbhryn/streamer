package streamer

type config_t struct {
	topics map[string]struct{}
}

var config config_t

func init() {
	topics := make(map[string]struct{})
	// for _, topic := range strings.Split(os.Getenv("TOPICS"), ",") {
	// 	topics[topic] = struct{}{}
	// }
	topics["a"] = struct{}{}
	topics["b"] = struct{}{}
	topics["c"] = struct{}{}
	config = config_t{
		topics: topics,
	}
}
