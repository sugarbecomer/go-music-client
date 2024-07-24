package main

import (
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"log"
	"os"
	"testing"
	"time"
)

func TestPlayer(t *testing.T) {
	f, err := os.Open("D:\\CloudMusic\\嫁衣(伴奏)-幸福大街.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
