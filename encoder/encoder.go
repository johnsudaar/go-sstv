package encoder

import (
	"image"
	"time"

	"gopkg.in/errgo.v1"

	"github.com/johnsudaar/sstv/sound"
)

// See http://www.barberdsp.com/files/Dayton%20Paper.pdf
// And http://f1ult.free.fr/DIGIMODES/MULTIPSK/sstv.htm

type Encoder struct {
	Sound      *sound.Sound
	Image      image.Image
	VIS        int // VIS HEADER
	resX       int // Resolution X
	resY       int // Resolution Y
	PixelCoder func(i image.Image, s *sound.Sound) error
}

func (e *Encoder) WriteHeader() {
	// Leader tone
	e.Sound.AddTone(1900, 300*time.Millisecond)

	// Break
	e.Sound.AddTone(1200, 10*time.Millisecond)

	// Leader tone
	e.Sound.AddTone(1900, 300*time.Millisecond)

	// VIS Start bit
	e.Sound.AddTone(1200, 30*time.Millisecond)

	// 7 bits VIS Code
	curVIS := e.VIS
	for i := 1; i <= 7; i++ {
		if curVIS%2 == 0 {
			e.Sound.AddTone(1300, 30*time.Millisecond)
		} else {
			e.Sound.AddTone(1100, 30*time.Millisecond)
		}
		curVIS = curVIS / 2
	}

	// Parity bit (even)
	if e.VIS%2 == 1 {
		e.Sound.AddTone(1300, 30*time.Millisecond)
	} else {
		e.Sound.AddTone(1100, 30*time.Millisecond)
	}

	// VIS Stop bit
	e.Sound.AddTone(1200, 30*time.Millisecond)
}

func (e *Encoder) EncodeImage() error {
	if err := e.PixelCoder(e.Image, e.Sound); err != nil {
		return errgo.Mask(err)
	}
	return nil
}
