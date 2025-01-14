# trucktrack

projet de lecture bluetooth

## Start

To start the code, you need to execute :

`make build`
`make run`

If you want tot check that a MAC address is authorized:

`curl http://localhost:9090/author?ID=01:01:01:01:01:01`
Feel free to edit the MAC address to see the different result we have

If you want to send a message to the server;

`curl -X POST http://localhost:9090/activity -d "Your message"`

The serveur should return you this:

`{Your message}`
