package easing_test

import (
	"fmt"
	"os"

	"github.com/draoncc/tween"
	. "github.com/draoncc/tween/easing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// FuncInfo stores information about a func.
type FuncInfo struct {
	Name string
	Func tween.TransitionFunc
}

var _ = Describe("Basic Curves", func() {
	Describe("Linear", func() {
		It("should generate a linear curve", func() {
			Ω(Linear(.0)).Should(Equal(.0))
			Ω(Linear(.1)).Should(Equal(.1))
			Ω(Linear(.2)).Should(Equal(.2))
			Ω(Linear(.3)).Should(Equal(.3))
			Ω(Linear(.4)).Should(Equal(.4))
			Ω(Linear(.5)).Should(Equal(.5))
			Ω(Linear(.6)).Should(Equal(.6))
			Ω(Linear(.7)).Should(Equal(.7))
			Ω(Linear(.8)).Should(Equal(.8))
			Ω(Linear(.9)).Should(Equal(.9))
			Ω(Linear(1.)).Should(Equal(1.))
		})
	})
	Describe("Swing", func() {
		It("should generate a gentle ease-in-ease-out curve", func() {
			Ω(Swing(.0)).Should(Equal(.0))
			Ω(Swing(.1)).Should(BeNumerically("~", 0.024, .001))
			Ω(Swing(.2)).Should(BeNumerically("~", 0.095, .001))
			Ω(Swing(.3)).Should(BeNumerically("~", 0.206, .001))
			Ω(Swing(.4)).Should(BeNumerically("~", 0.345, .001))
			Ω(Swing(.5)).Should(BeNumerically("~", 0.500, .001))
			Ω(Swing(.6)).Should(BeNumerically("~", 0.654, .001))
			Ω(Swing(.7)).Should(BeNumerically("~", 0.793, .001))
			Ω(Swing(.8)).Should(BeNumerically("~", 0.904, .001))
			Ω(Swing(.9)).Should(BeNumerically("~", 0.975, .001))
			Ω(Swing(1.)).Should(Equal(1.))
		})
	})
	Describe("Ease", func() {
		It("should generate more advanced easing curves", func() {
			funcs := []FuncInfo{
				FuncInfo{"Linear", Linear},
				FuncInfo{"Swing", Swing},
				FuncInfo{"QuadIn", QuadIn},
				FuncInfo{"QuadOut", QuadOut},
				FuncInfo{"QuadInOut", QuadInOut},
				FuncInfo{"CubicIn", CubicIn},
				FuncInfo{"CubicOut", CubicOut},
				FuncInfo{"CubicInOut", CubicInOut},
				FuncInfo{"QuartIn", QuartIn},
				FuncInfo{"QuartOut", QuartOut},
				FuncInfo{"QuartInOut", QuartInOut},
				FuncInfo{"QuintIn", QuintIn},
				FuncInfo{"QuintOut", QuintOut},
				FuncInfo{"QuintInOut", QuintInOut},
				FuncInfo{"ExpoIn", ExpoIn},
				FuncInfo{"ExpoOut", ExpoOut},
				FuncInfo{"ExpoInOut", ExpoInOut},
				FuncInfo{"SineIn", SineIn},
				FuncInfo{"SineOut", SineOut},
				FuncInfo{"SineInOut", SineInOut},
				FuncInfo{"CircIn", CircIn},
				FuncInfo{"CircOut", CircOut},
				FuncInfo{"CircInOut", CircInOut},
				FuncInfo{"LogIn", LogIn},
				FuncInfo{"LogOut", LogOut},
				FuncInfo{"LogInOut", LogInOut},
				FuncInfo{"ElasticIn", ElasticIn},
				FuncInfo{"ElasticOut", ElasticOut},
				FuncInfo{"ElasticInOut", ElasticInOut},
				FuncInfo{"BackIn", BackIn},
				FuncInfo{"BackOut", BackOut},
				FuncInfo{"BackInOut", BackInOut},
				FuncInfo{"BounceIn", BounceIn},
				FuncInfo{"BounceOut", BounceOut},
				FuncInfo{"BounceInOut", BounceInOut},
			}
			html, err := os.Create("easing.html")
			Ω(err).Should(BeNil())
			defer html.Close()
			html.Write([]byte("<html><head><style>body {margin:1em;}\n.box { position: relative; border: 1px solid #ddd; margin: .5em; display: inline-block; height: 300px; width: 300px; }\n.box h1 { position: absolute; top: 0; left: 10px; font-size: small; font-weight: normal; }</style><title>Curves</title></head><body>"))
			for _, curve := range funcs {
				svg, err := os.Create(curve.Name + ".svg")
				Ω(err).Should(BeNil())
				defer svg.Close()
				// Header plus draw axis
				svg.Write([]byte("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><path d=\"M50,250 L250,250 M50,250 L50,50\" style=\"stroke:#000; fill:none;\"/><path d=\"M50,50 L250,50 L250,250\" style=\"stroke:#ccc; fill:none;\"/><path d=\"M50,250"))
				for x := 0.; x <= 1.; x += .005 {
					y := curve.Func(x)
					X := 200*x + 50
					Y := 250 - 200*y
					svg.Write([]byte(fmt.Sprintf(" L%d,%d", int(X), int(Y))))
				}
				svg.Write([]byte("\" style=\"stroke:#660000; fill:none;\"/></svg>"))
				// Add to the html overview
				html.Write([]byte("<div class=\"box\"><h1>" + curve.Name + "</h1><img height=\"100%\" src=\"" + curve.Name + ".svg\"></div>"))
			}
			html.Write([]byte("</body></html>"))
		})
	})
})
