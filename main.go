package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := flag.String("c", "/etc/shadowsocks-libev/config.json", "path to config file")
	flag.Parse()

	var uri string
	for _, s := range flag.Args() {
		if strings.HasPrefix(s, "ss://") {
			uri = s
		}
	}

	conf := new(config)
	if uri == "" {
		data, err := ioutil.ReadFile(*filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		json.Unmarshal(data, &conf)
		fmt.Println(conf.EncodeURI())
	} else {
		err := conf.DecodeURI(uri)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		data, err := json.Marshal(conf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(*filename, data, 0600)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}
	}
}
