package handlers

import (
	"io"
	"net/http"
	"strconv"
)

func QueueHandler(w http.ResponseWriter, r *http.Request) {
	var str string
	w.Header().Set("Content-Type", "text/plain")
	Listener.TestMaker()
	for _, msg := range Listener.Messages {
		str += "\nRecieved: " + strconv.Itoa(int(msg.OPCode))
	}
	io.WriteString(w, str)
}

func (q *Queue) TestMaker() {
	for i := 0; i <= 7; i++ {
		new := Message{
			OPCode: 1 << i,
		}

		q.Messages = append(q.Messages, new)
	}
}
