package main

import (
	"github.com/mrreallyyo/waifu-builder-go/backend"
	"github.com/vugu/vugu"
)

func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	storage := backend.NewStorage()

	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if s, ok := b.(backend.StorageSetter); ok {
			s.SetStorage(&storage)
		}
	})

	ret := &Root{}
	buildEnv.WireComponent(ret)

	return ret
}