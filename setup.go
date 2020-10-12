package main

import (
	"github.com/mrreallyyo/waifu-builder-go/backend"
	"github.com/mrreallyyo/waifu-builder-go/components"
	"github.com/vugu/vugu"
)

func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	storage := backend.NewStorage()

	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if s, ok := b.(backend.StorageSetter); ok {
			s.SetStorage(&storage)
		}
	})

	ret := &components.App{}
	buildEnv.WireComponent(ret)

	return ret
}