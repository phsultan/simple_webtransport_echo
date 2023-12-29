package webtransport

import (
  "context"
  "log"
  "github.com/quic-go/webtransport-go"
)

func handleSession(session *webtransport.Session) {
  // Logic for handling a session...
  stream, err := session.AcceptUniStream(context.Background())
  if err != nil {
    log.Println(err)
      return
  }
  log.Printf("uni stream accepted, id: %d", stream.StreamID())

  indication, err := receiveClientIndication(stream)
  if err != nil {
    log.Println(err)
      return
  }
  log.Printf("client indication: %+v", indication)

  if err := validateClientIndication(indication); err != nil {
    log.Println(err)
      return
  }

  // this method blocks.
  if err := communicate(session); err != nil {
    log.Println(err)
  }
}


