package webtransport

import (
  "context"
	"io"
  "log"
  "github.com/quic-go/webtransport-go"
)

const maxBufferLength = 1024

func handleSession(session *webtransport.Session) {
	log.Printf("[handleSession] new session")
  // Logic for handling a session...
  stream, err := session.AcceptStream(context.Background())
  if err != nil {
    log.Println(err)
		return
  }
  log.Printf("stream accepted, id: %d", stream.StreamID())

	// read no more than maxClientIndicationLength bytes.
	// reader := io.Reader(stream, maxBufferLength)
	// writer := io.Writer(stream)

	buf := make([]byte, maxBufferLength)

	for {
		n, err := stream.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error :%s", err)
			}
			break
		}
		log.Printf("Read %d bytes", n)

		if (n < maxBufferLength) {
			buf = buf[:n]
		}

		n, err = stream.Write(buf)
		log.Printf("Wrote %d bytes", n)
	}
}
