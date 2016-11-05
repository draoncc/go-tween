package tween

import "time"

// TransitionFunc calculates the percentage of the transition between the start
// and end values based tween (elapsed time) completion status.
// For example, a linear tween simply has a 1:1 ratio between completed
// and transition percentages and returns the completed value unchanged.
// An "ease in" transition may create a logorithmic relationship between
// the completion time and transition value for the first third, then a
// linear relationship for the remainder.
type TransitionFunc func(completed float64) float64

// Updater is the interface for updating the current value as it tweens between
// start and end values of a Tween.
type Updater interface {
	// Start signals the beginning of a tween and is sent before the
	// tweening begins. Start may be used to setup or pre-calculate updates.
	//
	// framerate is the number of frames per second in the tween
	// frames is the total number of frames that will be generated
	// frameTime is the duration for each frame
	// runningTime is the total duration for the entire tween
	Start(framerate, frames int, frameTime, runningTime time.Duration)

	// Update receives information about the current Tween Frame and should
	// be used to update output or state.
	Update(Frame Frame)

	// End signals the end of the tween and is called after all updates.
	// End may be used to clean up resources (e.g. update channels).
	End()
}

// Frame captures information about the current "frame" of a tween transition.
type Frame struct {
	Completed    float64       // Completed is the percentage 0.0 - 1.0 of elapsed time.
	Transitioned float64       // Transitioned is the percentage 0.0 - 1.0 of transition between start and end values of the tween.
	Index        int           // Index is the current frame index
	Elapsed      time.Duration // Elapsed is the current elapsed time in the tween.
}

// Tween runs a tween relying on transitioner and updater.
type Tween struct {
	Duration   time.Duration  // The total duration of the tween.
	Transition TransitionFunc // Transition calculates the transition curve for the tween.
	Updater    Updater        // Updater updates the tween values for each frame.
	Framerate  int            // The number of tween data points per second (defaults to 60 fps - like the real gamers use).

	playhead time.Duration // The playhead position
	reversed bool          // reversed indicates whether playback for this tween has been reversed.
	running  bool          // True if the tween is running
	complete bool          // True if the tween has completed
	stop     chan int      // Internal channel used to terminate the tween early
	pause    chan int      // Internal channel used to pause the tween while running
}

// New creates a basic tween with a framerate of 60fps.
func New(duration time.Duration, transition TransitionFunc, updater Updater) *Tween {
	return &Tween{
		Duration:   duration,
		Transition: transition,
		Updater:    updater,
		Framerate:  60,
	}
}

// Reversed is a getter for Tween.reversed, indicating whether or not this tween
// is playing backwards.
func (e *Tween) Reversed() bool {
	return e.reversed
}

// Running is a getter for Tween.running, indicating whether the tween is
// currently playing.
func (e *Tween) Running() bool {
	return e.running
}

// Complete is a getter for Tween.complete, indicating if the tween has
// completed playback.
func (e *Tween) Complete() bool {
	return e.complete
}

// Play causes the tween to be played forwards from the beginning.
func (e *Tween) Play() {
	e.Pause()
	e.Seek(0)

	// Based on fps we can calculate how long a frame is:
	frameDuration := time.Second / time.Duration(e.Framerate) // The duration in a frame
	frames := int(e.Duration / frameDuration)                 // The number of frames in the duration

	e.Updater.Start(e.Framerate, frames, frameDuration, e.Duration)

	// Initializing values for "forward" playback
	e.reversed = false
	e.playhead = 0

	e.complete = false
	e.stop = make(chan int)

	// Send initial frame
	frame := Frame{}
	e.Updater.Update(frame)

	e.Resume()
}

// PlayReverse causes the tween to be played backwards from the end.
func (e *Tween) PlayReverse() {
	e.Pause()
	e.Seek(e.Duration)

	// Based on fps we can calculate how long a frame is:
	frameDuration := time.Second / time.Duration(e.Framerate) // The duration in a frame
	frames := int(e.Duration / frameDuration)                 // The number of frames in the duration

	e.Updater.Start(e.Framerate, frames, frameDuration, e.Duration)

	// Initializing values for reversed playback
	e.reversed = true
	e.playhead = e.Duration

	e.complete = false
	e.stop = make(chan int)

	// Send initial frame
	frame := Frame{1, 1, frames, e.Duration}
	e.Updater.Update(frame)

	e.Resume()
}

// Reverse reverses the tweens direction without affecting its playback.
func (e *Tween) Reverse() {
	// Ensuring that playback will not be affected
	if e.running == true {
		e.Pause()
		defer e.Resume()
	}

	e.reversed = !e.reversed
}

// Seek sets the playhead without affecting its playback.
// If the position exceeds the total tween duration, the tween's duration is
// taken instead.
func (e *Tween) Seek(position time.Duration) {
	// Ensuring that playback will not be affected
	if e.running == true {
		e.Pause()
		defer e.Resume()
	}

	if position > e.Duration {
		e.playhead = e.Duration
		return
	}

	e.playhead = position
}

// Resume resumes the tweens playback without affecting the direction.
func (e *Tween) Resume() {
	if e.running || e.complete {
		return
	}

	e.running = true
	e.pause = make(chan int)

	// Threading the Tweens playback to not stop other operations.
	go func() {
		// Based on fps we can calculate how long a frame is:
		frameDuration := time.Second / time.Duration(e.Framerate) // The duration in a frame
		cutoff := e.Duration - frameDuration                      // The cutoff point where elapsed time is considered "stop"

		// Initializing empty frame which will get updated.
		frame := Frame{}

		// set start time of current animation to calculate elapsed time based on playhead
		ticker := time.NewTicker(frameDuration)
		timeChan := ticker.C
		started := time.Now()

	L:
		for e.running {
			select {
			case <-timeChan:
				if e.reversed == false {
					frame.Elapsed = e.playhead + time.Since(started)
				} else {
					frame.Elapsed = e.playhead - time.Since(started)
				}

				// Calculate the frame index - some frames can be skipped so
				// must find correct time slot for this elapsed time
				frame.Index = int(float32(frame.Elapsed)/float32(frameDuration) + 0.5)

				// Calculate the completed percentage of time
				frame.Completed = ((float64(frame.Index) * float64(frameDuration)) / float64(e.Duration))
				if frame.Completed > 1 || frame.Completed <= 0 {
					go e.Stop()
					break
				}

				// Calulate the completed percentage of the transition
				frame.Transitioned = e.Transition(frame.Completed)

				// Update the value
				e.Updater.Update(frame)

				// see if we should keep going
				if frame.Elapsed > cutoff || frame.Elapsed < frameDuration {
					go e.Stop() // terminate ourself
				}
			case <-e.stop:
				e.complete = true
				break L
			case <-e.pause:
				break L
			}
		}

		// cleanup
		ticker.Stop()

		e.playhead = frame.Elapsed
		e.running = false

		// If tween has completed, update with a final frame and send the end signal
		if e.complete {
			if e.reversed == false {
				frame.Elapsed = e.Duration
				frame.Completed = 1
				frame.Transitioned = 1
				frame.Index = int(e.Duration / frameDuration)
			} else {
				frame.Elapsed = 0
				frame.Completed = 0
				frame.Transitioned = 0
				frame.Index = 0
			}
			e.Updater.Update(frame)
			e.Updater.End()
		}
	}()
}

// Pause pauses the tween in place.
func (e *Tween) Pause() {
	if e.running == true {
		close(e.pause)
		// Ensuring the tween has truly paused before continuation
		for e.running {
		}
	}
}

// Stop terminates the tween immediately.
func (e *Tween) Stop() {
	if e.running == true {
		close(e.stop)
		// Ensuring the tween has truly stopped before continuation
		for e.running {
		}
	}
}
