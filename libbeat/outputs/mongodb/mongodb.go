package mongodb

import (
	"fmt"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/outputs"
)

func init() {
	outputs.RegisterType("mongodb", makeMongodb)
}

func makeMongodb(
	_ outputs.IndexManager,
	beat beat.Info,
	observer outputs.Observer,
	cfg *common.Config,
) (outputs.Group, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}
	fmt.Println("mongodb config:", config)

	return outputs.Group{}, nil
}
