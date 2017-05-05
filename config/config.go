// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Brokers []string `config:"brokers"`
	Topics  []string `config:"topics"`
	Group   string   `config:"group"`
}

var DefaultConfig = Config{
	Brokers: []string{"localhost:9092"},
	Topics:  []string{"test1", "test2"},
	Group:   "kafkabeat",
}
