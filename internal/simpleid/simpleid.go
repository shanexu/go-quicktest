package simpleid

import (
	"sync/atomic"
	"time"
)

var prevID int64

func NextID() int64 {
	for {
		prev := atomic.LoadInt64(&prevID)
		now := time.Now().UnixMilli()
		for now <= prev {
			prev = atomic.LoadInt64(&prevID)
			now = time.Now().UnixMilli()
		}
		if atomic.CompareAndSwapInt64(&prevID, prev, now) {
			return now
		}
	}
}
