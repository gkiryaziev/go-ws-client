##	Golang Publish/Subscribe Websocket Client

[Go](https://golang.org/) websocket client with [Gorilla](http://www.gorillatoolkit.org/) toolkit.

This client was written for the Raspberry Pi 2. At this time, you can get Cpu temp and memory, and System memory total, used and free. Led implemented only as a stub, work continues on them.
With this client you can subscribe, unsubscribe and publish messages.
Implementation of Public/Subscribe server can be found [here](https://github.com/gkiryaziev/go_gorilla_pubsub_websocket_server).

ACTION - `SUBSCRIBE`, `UNSUBSCRIBE`, `PUBLISH`

Message example:
```
{"action" : "ACTION", "topic" : "TOPIC NAME", "data" : "DATA"}
```

![Mind](/mind.png?raw=true "Mind")

## Installation

#### 1. Install GO
#### 2. Install GB
  `go get -u github.com/constabulary/gb/...`
#### 3. Clone project
  `git clone https://github.com/gkiryaziev/go_gorilla_pubsub_websocket_client.git`
#### 4. Restore vendors
  `cd go_gorilla_pubsub_websocket_client`
  
  `gb vendor restore`
#### 5. Edit configuration
  Copy `config.default.yaml` to `config.yaml` and edit configuration.
#### 6. Build and Run project
  `gb build && bin/rpi.ws.client run`