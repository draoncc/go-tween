package easing

import "math"

// Auto-generated file - do not edit directly! See source in easing/gen/gen.go

// QuadIn eases in a Quad transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuadIn(completed float64) float64 {
	return math.Pow(completed, 2)
}

// QuadOut eases out a Quad transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuadOut(completed float64) float64 {
	return 1 - QuadIn(1-completed)
}

// QuadInOut eases in and out a Quad transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuadInOut(completed float64) float64 {
	if completed < 0.5 {
		return QuadIn(completed*2) / 2
	}
	return 1 - QuadIn((completed*-2)+2)/2
}

// CubicIn eases in a Cubic transition.
// See http://jqueryui.com/easing/ for curve in action.
func CubicIn(completed float64) float64 {
	return math.Pow(completed, 3)
}

// CubicOut eases out a Cubic transition.
// See http://jqueryui.com/easing/ for curve in action.
func CubicOut(completed float64) float64 {
	return 1 - CubicIn(1-completed)
}

// CubicInOut eases in and out a Cubic transition.
// See http://jqueryui.com/easing/ for curve in action.
func CubicInOut(completed float64) float64 {
	if completed < 0.5 {
		return CubicIn(completed*2) / 2
	}
	return 1 - CubicIn((completed*-2)+2)/2
}

// QuartIn eases in a Quart transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuartIn(completed float64) float64 {
	return math.Pow(completed, 4)
}

// QuartOut eases out a Quart transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuartOut(completed float64) float64 {
	return 1 - QuartIn(1-completed)
}

// QuartInOut eases in and out a Quart transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuartInOut(completed float64) float64 {
	if completed < 0.5 {
		return QuartIn(completed*2) / 2
	}
	return 1 - QuartIn((completed*-2)+2)/2
}

// QuintIn eases in a Quint transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuintIn(completed float64) float64 {
	return math.Pow(completed, 5)
}

// QuintOut eases out a Quint transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuintOut(completed float64) float64 {
	return 1 - QuintIn(1-completed)
}

// QuintInOut eases in and out a Quint transition.
// See http://jqueryui.com/easing/ for curve in action.
func QuintInOut(completed float64) float64 {
	if completed < 0.5 {
		return QuintIn(completed*2) / 2
	}
	return 1 - QuintIn((completed*-2)+2)/2
}

// ExpoIn eases in a Expo transition.
// See http://jqueryui.com/easing/ for curve in action.
func ExpoIn(completed float64) float64 {
	return math.Pow(completed, 6)
}

// ExpoOut eases out a Expo transition.
// See http://jqueryui.com/easing/ for curve in action.
func ExpoOut(completed float64) float64 {
	return 1 - ExpoIn(1-completed)
}

// ExpoInOut eases in and out a Expo transition.
// See http://jqueryui.com/easing/ for curve in action.
func ExpoInOut(completed float64) float64 {
	if completed < 0.5 {
		return ExpoIn(completed*2) / 2
	}
	return 1 - ExpoIn((completed*-2)+2)/2
}

// SineIn eases in a Sine transition.
// See http://jqueryui.com/easing/ for curve in action.
func SineIn(completed float64) float64 {
	return 1 - math.Cos(completed*math.Pi/2)
}

// SineOut eases out a Sine transition.
// See http://jqueryui.com/easing/ for curve in action.
func SineOut(completed float64) float64 {
	return 1 - SineIn(1-completed)
}

// SineInOut eases in and out a Sine transition.
// See http://jqueryui.com/easing/ for curve in action.
func SineInOut(completed float64) float64 {
	if completed < 0.5 {
		return SineIn(completed*2) / 2
	}
	return 1 - SineIn((completed*-2)+2)/2
}

// CircIn eases in a Circ transition.
// See http://jqueryui.com/easing/ for curve in action.
func CircIn(completed float64) float64 {
	return 1 - math.Sqrt(1-completed*completed)
}

// CircOut eases out a Circ transition.
// See http://jqueryui.com/easing/ for curve in action.
func CircOut(completed float64) float64 {
	return 1 - CircIn(1-completed)
}

// CircInOut eases in and out a Circ transition.
// See http://jqueryui.com/easing/ for curve in action.
func CircInOut(completed float64) float64 {
	if completed < 0.5 {
		return CircIn(completed*2) / 2
	}
	return 1 - CircIn((completed*-2)+2)/2
}

// LogIn eases in a Log transition.
// See http://jqueryui.com/easing/ for curve in action.
func LogIn(completed float64) float64 {
	return 1 - math.Log((1-completed)*(math.E-1)+1)
}

// LogOut eases out a Log transition.
// See http://jqueryui.com/easing/ for curve in action.
func LogOut(completed float64) float64 {
	return 1 - LogIn(1-completed)
}

// LogInOut eases in and out a Log transition.
// See http://jqueryui.com/easing/ for curve in action.
func LogInOut(completed float64) float64 {
	if completed < 0.5 {
		return LogIn(completed*2) / 2
	}
	return 1 - LogIn((completed*-2)+2)/2
}

// ElasticIn eases in a Elastic transition.
// See http://jqueryui.com/easing/ for curve in action.
func ElasticIn(completed float64) float64 {
	if completed == 0 || completed == 1 {
		return completed
	}
	return -math.Pow(2, 8*(completed-1)) * math.Sin(((completed-1)*80-7.5)*math.Pi/15)
}

// ElasticOut eases out a Elastic transition.
// See http://jqueryui.com/easing/ for curve in action.
func ElasticOut(completed float64) float64 {
	return 1 - ElasticIn(1-completed)
}

// ElasticInOut eases in and out a Elastic transition.
// See http://jqueryui.com/easing/ for curve in action.
func ElasticInOut(completed float64) float64 {
	if completed < 0.5 {
		return ElasticIn(completed*2) / 2
	}
	return 1 - ElasticIn((completed*-2)+2)/2
}

// BackIn eases in a Back transition.
// See http://jqueryui.com/easing/ for curve in action.
func BackIn(completed float64) float64 {
	return completed * completed * (3*completed - 2)
}

// BackOut eases out a Back transition.
// See http://jqueryui.com/easing/ for curve in action.
func BackOut(completed float64) float64 {
	return 1 - BackIn(1-completed)
}

// BackInOut eases in and out a Back transition.
// See http://jqueryui.com/easing/ for curve in action.
func BackInOut(completed float64) float64 {
	if completed < 0.5 {
		return BackIn(completed*2) / 2
	}
	return 1 - BackIn((completed*-2)+2)/2
}

// BounceIn eases in a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func BounceIn(completed float64) float64 {

	bounce := float64(3)
	var pow2 float64
	for pow2 = math.Pow(2, bounce); completed < ((pow2 - 1) / 11); pow2 = math.Pow(2, bounce) {
		bounce--
	}
	return 1/math.Pow(4, 3-bounce) - 7.5625*math.Pow((pow2*3-2)/22-completed, 2)
}

// BounceOut eases out a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func BounceOut(completed float64) float64 {
	return 1 - BounceIn(1-completed)
}

// BounceInOut eases in and out a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func BounceInOut(completed float64) float64 {
	if completed < 0.5 {
		return BounceIn(completed*2) / 2
	}
	return 1 - BounceIn((completed*-2)+2)/2
}
