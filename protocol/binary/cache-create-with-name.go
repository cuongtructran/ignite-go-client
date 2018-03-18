package binary

import (
	"io"
	"math/rand"

	"github.com/amsokol/ignite-go-client/protocol/binary/internal"

	stderr "errors"
	"github.com/amsokol/go-errors"
)

// CacheCreateWithName Creates a cache with a given name.
// Cache template can be applied if there is a '*' in the cache name.
func CacheCreateWithName(rw io.ReadWriter, name string) (Result, error) {
	var res Result
	id := rand.Int63()

	if err := internal.Write(rw, opCacheCreateWithName, id,
		internal.NewString(name)); err != nil {
		return res, errors.Wrapf(err, "failed to write operation request")
	}

	if err := internal.Read(rw, true, &id, &res.Status, &res.Message); err != nil {
		return res, errors.Wrapf(err, "failed to read operation response")
	}

	if res.Status != StatusSuccess {
		return res, stderr.New(res.Message)
	}

	return res, nil
}
