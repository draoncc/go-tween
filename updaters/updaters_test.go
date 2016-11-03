package updaters_test

import (
	"image/color"
	"time"

	"github.com/draoncc/tween"
	"github.com/draoncc/tween/easing"
	. "github.com/draoncc/tween/updaters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	Describe("Color Tween", func() {
		It("should generate color tween values", func(done Done) {
			start := color.RGBA{255, 0, 0, 255}
			end := color.RGBA{0, 128, 255, 0}

			updater := NewColor(start, end)
			engine := tween.NewEngine(time.Second, easing.Linear, updater)
			engine.Start()

			running := true
			colors := []color.RGBA{}
			for running {
				select {
				case color := <-updater.Updates:
					colors = append(colors, color)
				case <-updater.Done:
					running = false
				}
			}
			Ω(colors).Should(HaveLen(61))
			Ω(colors[len(colors)-1]).Should(Equal(end))
			Ω(colors[0]).Should(Equal(start))
			close(done)
		}, 2)
	})
})
