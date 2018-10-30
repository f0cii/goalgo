package goalgo

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/facebookgo/inject"

	stdlog "log"

	"github.com/hashicorp/go-plugin"
	"github.com/sumorf/goalgo/log"
)

func Serve(strategy Strategy) {
	flag.StringVar(&id, "id", "", "")
	flag.IntVar(&sid, "sid", 0, "")
	flag.StringVar(&address, "address", "127.0.0.1:9900", "")

	flag.Parse()

	strategy.SetSelf(strategy)

	if id == "^" {
		options := strategy.GetOptions()
		data, err := json.Marshal(options)
		if err != nil {
			panic(err)
		}
		fmt.Printf(string(data))
		os.Exit(0)
		return
	}

	stdlog.Printf("Serve id=%v sid=%v address=%v", id, sid, address)

	var g inject.Graph

	l := &GRPCLog{SID: sid}
	g.Provide(
		&inject.Object{Value: log.GetLogger()},
		&inject.Object{Value: l},
	)
	if err := g.Populate(); err != nil {
		stdlog.Fatal(err)
	}

	plugins := map[string]plugin.Plugin{
		PluginMapStrategyCtlKey: &StrategyPlugin{Impl: strategy},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         plugins,
	})
}
