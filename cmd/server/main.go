// Command server runs the CoffeeChat relay server.
//
// The relay only ever handles ciphertext: it stores users' public keys and
// prekeys, queues encrypted envelopes for offline recipients, and deletes them
// once delivered. It never sees plaintext messages.
package main

import "log"

func main() {
	log.Println("CoffeeChat relay starting")

	// TODO(M5): open the SQLite store and add the admin "invite create" CLI.
	// TODO(M6): start the HTTP server with registration, auth, and the WebSocket hub.
}
