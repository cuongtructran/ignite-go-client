package ignite

import (
	"encoding/binary"
	"io"
	"math/rand"

	"github.com/amsokol/ignite-go-client/binary/errors"
)

// RequestOperation is struct to store operation request
type RequestOperation struct {
	Code int
	UID  int64

	request
}

// WriteTo is function to write handshake request data to io.Writer.
// Returns written bytes.
func (r *RequestOperation) WriteTo(w io.Writer) (int64, error) {
	// write payload length
	l := int32(r.payload.Len())
	if err := binary.Write(w, binary.LittleEndian, &l); err != nil {
		return 0, errors.Wrapf(err, "failed to write operation request length")
	}

	// write operation code
	code := int16(r.Code)
	if err := binary.Write(w, binary.LittleEndian, &code); err != nil {
		return 0, errors.Wrapf(err, "failed to write operation code")
	}

	// write operation request id
	if err := binary.Write(w, binary.LittleEndian, &r.UID); err != nil {
		return 0, errors.Wrapf(err, "failed to write operation request id")
	}

	// write payload
	n, err := r.payload.WriteTo(w)
	return 4 + 2 + 8 + n, err
}

// NewRequestOperation creates new handshake request object
func NewRequestOperation(code int) *RequestOperation {
	return &RequestOperation{request: newRequest(), Code: code, UID: rand.Int63()}
}