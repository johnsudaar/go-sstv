package sound

import (
	"math"
	"time"

	"gopkg.in/errgo.v1"

	"github.com/unixpickle/wav"
)

type Sound struct {
	sampleRate     int
	sound          wav.Sound
	singleDuration time.Duration
	lastPoint      float64
}

func NewSound(sampleRate int) *Sound {
	return &Sound{
		sampleRate:     sampleRate,
		lastPoint:      float64(0),
		sound:          wav.NewPCM8Sound(1, sampleRate),
		singleDuration: time.Duration(time.Second.Nanoseconds()/int64(sampleRate)) * time.Nanosecond,
	}
}

func (s *Sound) AddTone(frequency int, length time.Duration) {
	i := int64(0)
	for i = int64(0); i < length.Nanoseconds()/s.singleDuration.Nanoseconds(); i++ {
		s.lastPoint += float64(1) / float64(s.sampleRate) * 2 * math.Pi * float64(frequency)
		value := wav.Sample(0.8 * math.Sin(s.lastPoint))
		s.sound.SetSamples(append(s.sound.Samples(), value))
	}
}

func (s *Sound) WriteFile(path string) error {
	if err := wav.WriteFile(s.sound, path); err != nil {
		return errgo.Mask(err)
	}
	return nil
}
