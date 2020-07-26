package backend

import (
	"bytes"
	"math"

	bolt "go.etcd.io/bbolt"
)

// UnsafeForEach must be called holding the lock on the tx

func unsafeForEach(tx *bolt.Tx,
	bucket []byte,
	visitor func(k, v []byte) error) error {
	if b := tx.Bucket(bucket); b != nil {
		return b.ForEach(visitor)
	}
	return nil
}

func unsafeRange(c *bolt.Cursor,
	key, endKey []byte,
	limit int64) (keys [][]byte, vs [][]byte) {
	if limit <= 0 {
		limit = math.MaxInt64
	}

	var isMatch func(b []byte) bool
	if len(endKey) > 0 {
		isMatch = func(b []byte) bool {
			return bytes.Compare(b, endKey) < 0
		}
	} else {
		isMatch = func(b []byte) bool {
			return bytes.Equal(b, key)
		}
		limit = 1
	}

	for ck, cv := c.Seek(key); ck != nil && isMatch(ck); ck, cv = c.Next() {
		vs = append(vs, cv)
		keys = append(keys, ck)
		if limit == int64(len(keys)) {
			break
		}
	}

	return keys, vs
}