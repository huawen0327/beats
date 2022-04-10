package mongodb

type config struct {
	Hosts      []string `config:"hosts"`
	DB         string   `config:"db"`
	Collection string   `config:"collection"`
}

func defaultConfig() config {
	return config{
		Hosts:      []string{"localhost:27017"},
		DB:         "test",
		Collection: "test",
	}
}
