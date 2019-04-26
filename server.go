package main

import "math/rand"
import "sync/atomic"
import "flag"
import "time"
import "net/rpc"
import "net"
import "log"

type IdService struct {
	epoch    time.Time
	serverId uint64
	prev     uint64
	//mutex    sync.Mutex
}

func (s *IdService) Next(_ uint64, reply *uint64) error {
	//s.mutex.Lock()
	//defer s.mutex.Unlock()
	//id, err := GenerateId(s.epoch, s.serverId, s.prev)
	//if err != nil {
	//	return err
	//}
	//s.prev = id
	//*reply = id
	//return nil
	for {
		prev := s.prev
		id, err := GenerateId(s.epoch, s.serverId, prev)
		if err != nil {
			return err
		}
		if !atomic.CompareAndSwapUint64(&s.prev, prev, id) {
			time.Sleep(time.Duration(rand.Float64()*1000) * time.Nanosecond)
			continue
		}
		*reply = id
		return nil
	}
}

func main() {
	serverIdVar := flag.Uint64("serverId", 0, "server id")
	epochVar := flag.String("epoch", "Mon Jan 2 15:04:05 MST 2006", "epoch")
	listenVar := flag.String("addr", ":1234", "address to bind to")
	flag.Parse()

	epoch, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", *epochVar)
	if err != nil {
		log.Fatal("cannot parse epoch:", err)
	}

	if time.Since(epoch) < 0 {
		log.Fatal("epoch is in the future!")
	}

	if *serverIdVar >= 1024 {
		log.Fatal("server id out of range!")
	}

	service := new(IdService)
	service.epoch = epoch
	service.serverId = *serverIdVar

	rpc.RegisterName("Flaked", service)
	l, e := net.Listen("tcp", *listenVar)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)
}
