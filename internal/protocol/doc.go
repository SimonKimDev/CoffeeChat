// Package protocol defines the wire types shared by the CoffeeChat relay
// server and the desktop client: envelopes, registration and auth messages,
// prekey bundles, delivery acks, and so on.
//
// Everything in here must be safe for the server to see. Message contents are
// carried only as opaque ciphertext produced by the client.
//
// You'll fill this in during milestone M4 (wire protocol design).
package protocol
