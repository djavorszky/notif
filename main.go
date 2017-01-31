// Package notif is a simple library that makes it easy to send JSON messages
// to a specified address:port/endpoint location. The messages are sent as plain
// text.
package notif

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Dest should be a valid URL on the remote host. It should be the full package,
// meaning, it should contain everything from schema till the URI. This is
// used primarily, except if the method specifies otherwise.
var Dest = "http://localhost:80/message"

// Snd JSONifies the message, then sends it as a POST request to the DefaultDest.
func Snd(msg Fireable) error {
	return SndLoc(msg, Dest)
}

// SndLoc JSONifies the message, then sends it as a POST request to the specified destination.
func SndLoc(msg Fireable, dest string) error {
	jMsg, err := jsonify(msg)
	if err != nil {
		return err
	}

	statusCode, err := sendReq(Dest, jMsg)
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return fmt.Errorf("Got non-200 response: %d", statusCode)
	}

	log.Println("Msg sent successfully")
	return nil
}

// Msg is the simplest type of message that can be sent. It has an ID field which
// can be used to identify the message if it belongs to a conversation, a StatusID
// which can correspond to statuses, similar to that of the HTTP response codes,
// and a Message that can contain any text, aimed to have a human readable message.
//
// It also has a seperate Dest field where a custom destination can be specified.
// If empty, DefaultDest will be used.
type Msg struct {
	ID, StatusID int
	Message      string
}

// Fireable is an empty interface. This way, custom structs can also be used. There
// is no restriction on what can be applied here, as long as it's a struct.
type Fireable interface{}

func sendReq(dest string, msg []byte) (int, error) {
	req, err := http.NewRequest(http.MethodPost, dest, bytes.NewBuffer(msg))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	log.Println("Sending msg")
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func jsonify(msg Fireable) ([]byte, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return b, err
}
