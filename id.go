package main

import "errors"
import "time"

var CounterOverflow = errors.New("GenerateId: counter overflow")

func GetCounter(u uint64) uint64 {
	return u & ^(^0 << 13)
}

func GenerateId(epoch time.Time, serverId uint64, prev uint64) (uint64, error) {
	duration := time.Since(epoch)
	prevCounter := GetCounter(prev)

	id := uint64(duration.Nanoseconds())
	id = (id >> 23) << 23      // keep first 41 bits
	id = id | (serverId << 13) // or in 10 bits of serverid

	// If there is a collision then we need to OR in the new counter value
	if id == (prev ^ prevCounter) {
		counter := GetCounter(prevCounter + 1)
		if counter == 0 {
			return 0, CounterOverflow
		}
		id |= counter
	}
	return id, nil
}
