##	Golang Publish/Subscribe Websocket Client

[Go](https://golang.org/) websocket client with [Gorilla](http://www.gorillatoolkit.org/) toolkit.

This client was written for the Raspberry Pi 2. At this time, the function Temp, Memory, Led implemented only as a stub, work continues on them.
With this client you can subscribe, unsubscribe and publish messages.
Implementation of Public/Subscribe server can be found [here](https://github.com/gkiryaziev/go_gorilla_pubsub_websocket_server).

ACTION - `SUBSCRIBE`, `UNSUBSCRIBE`, `PUBLISH`

Message example:
```
{"action" : "ACTION", "topic" : "TOPIC NAME", "data" : "DATA"}
```

![Mind](/mind.png?raw=true "Mind")