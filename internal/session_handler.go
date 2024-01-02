package webtransport

import (
  "context"
	"io"
  "log"
	"time"
  "github.com/quic-go/webtransport-go"
)

const maxBufferLength = 1024

func handleSession(session *webtransport.Session) {
  // Logic for handling a session...
  stream, err := session.AcceptStream(context.Background())
  if err != nil {
    log.Println(err)
		return
  }
  log.Printf("stream accepted, id: %d", stream.StreamID())

	// read no more than maxClientIndicationLength bytes.
	reader := io.LimitReader(stream, maxBufferLength)
	writer := io.Writer(stream)

	buf := make([]byte, maxBufferLength)
	n, err := reader.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Printf("Error :%s", err)
		}
	}

	log.Printf("Read %d bytes", n)
	value := string(buf)

	log.Printf("value :%s", value)

	time.Sleep(2 * time.Second)
	writer.Write(buf)


	time.Sleep(4 * time.Second)
}
