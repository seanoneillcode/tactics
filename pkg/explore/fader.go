package explore

type Fader struct {
	isFadeOut bool
	isFadeIn  bool
	// the current amount of fade time
	fadeTimer int64
	// the total amount of fade time
	totalFadeTime int64

	// the resulting amount of scaled alpha value
	fadeLevel float64
}

func NewFader() *Fader {
	return &Fader{
		fadeLevel: 1,
	}
}

func (f *Fader) Update(delta int64, state *State) {
	if f.isFadeOut {
		f.fadeTimer = f.fadeTimer - delta
		if f.fadeTimer < 0 {
			f.isFadeOut = false
			f.isFadeIn = true
			f.fadeTimer = f.totalFadeTime
		}
	}
	if f.isFadeIn {
		f.fadeTimer = f.fadeTimer - delta
		if f.fadeTimer < 0 {
			f.isFadeIn = false
		}
	}

	if f.fadeTimer > 0 {
		if f.isFadeOut {
			f.fadeLevel = float64(f.fadeTimer) / float64(f.totalFadeTime)
		}
		if f.isFadeIn {
			f.fadeLevel = float64(f.totalFadeTime-f.fadeTimer) / float64(f.totalFadeTime)
		}
	} else {
		f.fadeLevel = 1
	}
}

func (f *Fader) FadeOutAndIn(timeAmount int64) {
	f.isFadeOut = true
	f.fadeTimer = timeAmount
	f.totalFadeTime = timeAmount
}
