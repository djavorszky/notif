package notif

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestJSONifyMsg(t *testing.T) {
	var msg2 Msg

	msg := Msg{1, http.StatusOK, "TEST MESSAGE"}

	b, err := jsonify(msg)
	if err != nil {
		t.Errorf("Couldn't jsonify Msg: %q", err.Error())
	}

	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&msg2)
	if err != nil {
		t.Errorf("Couldn't unjsonify jsonified Msg: %q", err.Error())
	}

	if msg2 != msg {
		t.Errorf("Seems like the two msgs are not the same: Orig:%q\nvs\nRe-jsonified%q", msg, msg2)
	}

}
