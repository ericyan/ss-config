package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flags := flag.NewFlagSet("ss-config", flag.ExitOnError)
	filename := flags.String("c", "/etc/shadowsocks-libev/config.json", "path to config file")
	server := flags.String("s", "0.0.0.0", "server hostname or address")
	serverPort := flags.Int("p", 8388, "server port")
	localPort := flags.Int("l", 1080, "client port")
	password := flags.String("k", "", "pre-shared key")
	method := flags.String("m", "chacha20-ietf-poly1305", "encrypt method")
	timeout := flags.Int("t", 60, "socket timeout in seconds")

	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\tss-config command [-flags] arguments\n")
		os.Exit(2)
	}

	flags.Parse(os.Args[2:])

	if *password == "" {
		buf := make([]byte, 18)
		_, err := rand.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		psk := base64.RawURLEncoding.EncodeToString(buf)
		password = &psk
	}

	var uri string
	for _, s := range flags.Args() {
		if strings.HasPrefix(s, "ss://") {
			uri = s
		}
	}

	conf := &config{
		Server:     *server,
		ServerPort: *serverPort,
		LocalPort:  *localPort,
		Password:   *password,
		Method:     *method,
		Timeout:    *timeout,
	}

	switch os.Args[1] {
	case "show":
		conf, err := readFile(*filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		pp, err := json.MarshalIndent(conf, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		os.Stdout.Write(pp)
	case "uri":
		conf, err := readFile(*filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}

		fmt.Println(conf.EncodeURI())
	case "new":
		if uri != "" {
			err := conf.DecodeURI(uri)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
				os.Exit(1)
			}
		}

		err := conf.WriteFile(*filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ss-config: %s\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("invalid command: %s\n", os.Args[1])
		os.Exit(2)
	}
}
