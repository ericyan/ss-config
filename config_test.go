package main

import "testing"

func TestConfig(t *testing.T) {
	cases := []struct {
		uri      string
		valid    bool
		method   string
		password string
		hostname string
		port     int
	}{
		{"ss://YmYtY2ZiOnRlc3RAMTkyLjE2OC4xMDAuMTo4ODg4", true, "bf-cfb", "test", "192.168.100.1", 8888},
	}

	for _, c := range cases {
		conf := new(config)
		err := conf.DecodeURI(c.uri)
		if c.valid && err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if !c.valid && err == nil {
			t.Errorf("invalid case undetected: %s", c.uri)
		}

		if c.valid {
			if conf.Method != c.method || conf.Password != c.password || conf.Server != c.hostname || conf.ServerPort != c.port {
				t.Errorf("unexpected decode result: got %s:%s@%s:%d, want %s:%s@%s:%d",
					conf.Method, conf.Password, conf.Server, conf.ServerPort,
					c.method, c.password, c.hostname, c.port,
				)
			}

			if uri := conf.EncodeURI(); uri != c.uri {
				t.Errorf("unexpected encode result: got %s, want %s", uri, c.uri)
			}
		}
	}
}
