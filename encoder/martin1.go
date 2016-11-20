package encoder

import (
	"image"
	"time"

	"github.com/johnsudaar/sstv/sound"
)

func NewMartin1(sampleRate int, image image.Image) *Encoder {
	return &Encoder{
		VIS:        44,
		Sound:      sound.NewSound(sampleRate),
		Image:      image,
		resX:       320,
		resY:       256,
		PixelCoder: Martin1PixelCoder,
	}
}

func Martin1PixelCoder(i image.Image, s *sound.Sound) error {
	for y := 0; y < 256; y++ {
		// Sync Pulse
		s.AddTone(1200, 4862*time.Microsecond)

		// Sync porch
		s.AddTone(1500, 572*time.Microsecond)

		for c := 0; c < 3; c++ {
			for x := 0; x < 320; x++ {
				r, g, b, a := i.At(x, y).RGBA()
				value := g
				if c == 1 {
					value = b
				} else if c == 2 {
					value = r
				}

				tone := 1500 + (value*800)/a

				s.AddTone(int(tone), 457*time.Microsecond)
			}
			// Separator pulse
			s.AddTone(1500, 572*time.Microsecond)
		}
	}
	return nil
}
