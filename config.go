package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type config struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	LocalPort  int    `json:"local_port"`
	Password   string `json:"password"`
	Timeout    int    `json:"timeout"`
	Method     string `json:"method"`
}

func (c *config) EncodeURI() string {
	plain := fmt.Sprintf("%s:%s@%s:%d", c.Method, c.Password, c.Server, c.ServerPort)
	return "ss://" + base64.RawStdEncoding.EncodeToString([]byte(plain))
}

func (c *config) DecodeURI(uri string) error {
	uri = strings.TrimLeft(uri, "ss://")

	var plain string
	if strings.ContainsRune(uri, '@') {
		plain = uri
	} else {
		data, err := base64.RawStdEncoding.DecodeString(uri)
		if err != nil {
			return err
		}

		plain = string(data)
	}

	var args []string
	for _, s := range strings.Split(plain, "@") {
		args = append(args, strings.SplitN(s, ":", 2)...)
	}
	if len(args) != 4 {
		return errors.New("invalid url")
	}
	port, err := strconv.Atoi(args[3])
	if err != nil {
		return errors.New("invalid url")
	}

	c.Method = args[0]
	c.Password = args[1]
	c.Server = args[2]
	c.ServerPort = port
	return nil
}
