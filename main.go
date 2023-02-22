package main

import (
	"fmt"
	"github.com/volcengine/datatester-go-sdk/client"
	"github.com/volcengine/datatester-go-sdk/config"
	"os"
	"time"
)

func main() {
	cli := client.NewClient(
		os.Args[1],
		config.WithMetaHost(config.MetaHostSG),
		config.WithTrackHost(config.TrackHostSG),
		config.WithFetchInterval(100*time.Millisecond),
	)
	if cli == nil {
		fmt.Printf("Error with NewClient: client not initialized")
		os.Exit(1)
	}

	key := "datatester-test"
	id := "dummy-id"
	defaultValue := ""
	attributes := map[string]interface{}{}
	variant, err := cli.Activate(key, id, id, defaultValue, attributes)
	if err != nil {
		fmt.Printf("Error with Activate:  %+v", err)
		os.Exit(1)
	}

	if val, ok := variant.(string); ok {
		fmt.Printf("type=string, %s", val)
	} else if val, ok := variant.(map[string]interface{}); ok {
		fmt.Printf("type=map[string]interface{}, %+v", val)
	} else {
		fmt.Printf("Unknown type %+v", variant)
	}
}
