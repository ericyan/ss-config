package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := flag.String("c", "/etc/shadowsocks-libev/config.json", "path to config file")
	server := flag.String("s", "0.0.0.0", "server hostname or address")
	serverPort := flag.Int("p", 1234, "server port")
	localPort := flag.Int("l", 1080, "client port")
	password := flag.String("k", "", "pre-shared key")
	method := flag.String("m", "rc4-md5", "encrypt method")
	timeout := flag.Int("t", 60, "socket timeout in seconds")
	flag.Parse()

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
	for _, s := range flag.Args() {
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
