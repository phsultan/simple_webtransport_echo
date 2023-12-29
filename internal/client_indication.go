package webtransport

import (
	"log"
	"io"
	"encoding/binary"
	"errors"

	"github.com/quic-go/webtransport-go"
)

const maxClientIndicationLength = 65535

// define known client indication keys.
type clientIndicationKey int16

const (
	clientIndicationKeyOrigin clientIndicationKey = 0
	clientIndicationKeyPath                       = 1
)

// ClientIndication container.
type ClientIndication struct {
	// Origin client indication value.
	Origin string
	// Path client indication value.
	Path string
}

func receiveClientIndication(stream webtransport.ReceiveStream) (ClientIndication, error) {
	var clientIndication ClientIndication

	// read no more than maxClientIndicationLength bytes.
	reader := io.LimitReader(stream, maxClientIndicationLength)

	done := false

	for {
		if done {
			break
		}
		var key int16
		err := binary.Read(reader, binary.BigEndian, &key)
		if err != nil {
			log.Printf("err :%s", err)
			if err == io.EOF {
				done = true
			} else {
				return clientIndication, err
			}
		}

		log.Printf("key (hex): %x", key)

		var valueLength int16
		err = binary.Read(reader, binary.BigEndian, &valueLength)
		if err != nil {
			return clientIndication, err
		}

		log.Printf("valueLength (hex): %x", valueLength)

		buf := make([]byte, valueLength)
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				// still need to process indication value.
				done = true
			} else {
				return clientIndication, err
			}
		}
		if int16(n) != valueLength {
			return clientIndication, errors.New("read less than expected")
		}
		value := string(buf)

		switch clientIndicationKey(key) {
		case clientIndicationKeyOrigin:
			clientIndication.Origin = value
		case clientIndicationKeyPath:
			clientIndication.Path = value
		default:
			log.Printf("skip unknown client indication key: %d: %s", key, value)
		}
	}
	return clientIndication, nil
}

func validateClientIndication(indication ClientIndication) error {
	return errors.New("not implemented yet")
}

func communicate(session *webtransport.Session) error {
	return errors.New("not implemented yet")
}
