package kw

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-player/cst"
	"go-player/logs"
	"strings"
	"testing"
)

func TestKWSearch(t *testing.T) {
	logs.LogInit()
	k := NewKWApi()
	searchMusic := k.SearchMusic("aliez")
	if searchMusic == nil {
		panic("search null")
	}
	for _, v := range searchMusic.Abslist {
		_, mid, ok := strings.Cut(v.MUSICRID, "_")
		if ok {
			log.Infof("找到mid:%s", mid)
		}
	}
}

func TestDown(t *testing.T) {
	logs.LogInit()
	k := NewKWApi()
	u := k.getTrackUrl3("40900571", cst.HIGHT)
	log.Info(u)
}

func TestCalu(t *testing.T) {
	var v int64 = 2841135770
	var x = 32
	xx := int64(-4294967296)&int64(v<<x) | int64(0xFFFFFFFF)&int64(-3686726880)
	fmt.Println(xx)
}
