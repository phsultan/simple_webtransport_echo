package webtransport

import (
	"fmt"
	"log"
	"net/http"
	"github.com/quic-go/webtransport-go"
	"github.com/quic-go/quic-go/http3"
)

func StartServer(certFile string, keyFile string) error {
	// create a new webtransport.Server, listening on (UDP) port 443
	s := webtransport.Server{
		H3: http3.Server{Addr: ":4443"},
	}

	s.CheckOrigin = func(r *http.Request) bool {
		// Loop over header names
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println("[CheckOrigin]", name, value)
			}
		}
		return true
	}

	http.HandleFunc("/webtransport", func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Upgrade(w, r)
		if err != nil {
			log.Printf("upgrading failed: %s", err)
			w.WriteHeader(500)
			return
		}

		log.Printf("Connection succeeded from Origin %s", r.Header["Origin"])
		log.Printf("Connection succeeded from address %s", session.RemoteAddr())

		newSession(session)
	})

	err := s.ListenAndServeTLS(certFile, keyFile)
	return err
}

func newSession(session *webtransport.Session) {
	go func() {
		defer func() {
			_ = session.CloseWithError(0, "bye")
			log.Printf("close session: %s", session.RemoteAddr().String())
		}()
		handleSession(session)
	}()
}

