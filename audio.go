package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

type Audio interface {
	register(name string, filepath string) error
	play(name string) error
}

type audio struct {
	registry map[string]*beep.Buffer
}

var SAMPLE_RATE beep.SampleRate = 44100

func newAudio() *audio {
	err := speaker.Init(SAMPLE_RATE, SAMPLE_RATE.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	return &audio{
		registry: make(map[string]*beep.Buffer),
	}
}

func (a *audio) register(name string, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	a.registry[name] = buffer

	return nil
}

func (a *audio) play(name string) error {
	buffer, ok := a.registry[name]
	if !ok {
		return errors.New("Not found")
	}

	format := buffer.Format()
	streamer := buffer.Streamer(0, buffer.Len())

	if format.SampleRate == SAMPLE_RATE {
		speaker.Play(streamer)
	} else {
		resampler := beep.Resample(4, format.SampleRate, SAMPLE_RATE, streamer)
		speaker.Play(resampler)
	}

	return nil
}
