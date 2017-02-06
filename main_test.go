package notif

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/djavorszky/ddn/inet"
)

func TestJSONifyMsg(t *testing.T) {
	var msg2 Msg

	msg := Msg{1, http.StatusOK, "TEST MESSAGE"}

	b, err := inet.JSONify(msg)
	if err != nil {
		t.Errorf("Couldn't jsonify Msg: %q", err.Error())
	}

	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&msg2)
	if err != nil {
		t.Errorf("Couldn't unjsonify jsonified Msg: %q", err.Error())
	}

	if msg != msg2 {
		t.Errorf("Seems like the two msgs are not the same:\nOrig\t\t%v\nRe-jsonified\t%v", msg, msg2)
	}
}
