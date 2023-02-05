package main

import (
	"time"

	"github.com/yuanyu90221/blockchain-project/network"
)

func main() {
	trlLocal := network.NewLocalTransport("LOCAL")
	trlRemote := network.NewLocalTransport("REMOTE")

	trlLocal.Connect(trlRemote)
	trlRemote.Connect(trlLocal)
	go func() {
		for {
			trlRemote.SendMessage(trlLocal.Addr(), []byte("Hello World"))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trlLocal},
	}
	s := network.NewServer(opts)
	s.Start()
}
