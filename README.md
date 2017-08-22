# ss-config

`ss-config` is a configuration utility for [shadowsocks-libev].

## Usage

### show

Show current config in prettified JSON format:

```
$ ss-config show -c /path/to/config.json
{
  "server": "192.168.100.1",
  "server_port": 8888,
  "local_port": 1080,
  "password": "test",
  "timeout": 60,
  "method": "bf-cfb"
}
```

If you do not specify a config file (via the `-c` flag), it will default
to `/etc/shadowsocks-libev/config.json`.

### uri

Encode current config to a URI which can be used for QR code generation:

```
$ ss-config uri
ss://YmYtY2ZiOnRlc3RAMTkyLjE2OC4xMDAuMTo4ODg4
```

### new

Generate a config file from command-line flags:

```
$ ss-config new -s ss.example.com -p 8388 -m chacha20-ietf-poly1305
```

You can specify the output file by using the `-c` flag.  If you do not
set a password (via the `-k` flag), one will be generated for you.

You can also generate a config from a URI:

```
$ ss-config new -c /dev/stdout ss://YmYtY2ZiOnRlc3RAMTkyLjE2OC4xMDAuMTo4ODg4
{"server":"192.168.100.1","server_port":8888,"local_port":1080,"password":"test","timeout":60,"method":"bf-cfb"}
```

[shadowsocks-libev]: https://github.com/shadowsocks/shadowsocks-libev
