package domain

type Frame struct {
	Index   int  `json:"index"`
	Right   int  `json:"right"`
	Left    int  `json:"left"`
	HasLeft bool `json:"hasLeft"`
	Single  bool `json:"single"`
}

func FrameAt(index, count int, single bool) Frame {
	frame := Frame{Single: single}
	if count <= 0 {
		return frame
	}
	if index < 0 {
		index = 0
	}
	if index >= count {
		index = count - 1
	}
	frame.Index = index
	frame.Right = index
	if !single && index+1 < count {
		frame.Left = index + 1
		frame.HasLeft = true
	}
	return frame
}

// Step is the same rule as frontend/src/domain/reading.ts.
func Step(index, count int, single, forward bool) int {
	if count <= 0 {
		return 0
	}
	delta := 1
	if !single {
		delta = 2
	}
	if !forward {
		delta = -delta
	}
	next := index + delta
	if next < 0 {
		return 0
	}
	if next >= count {
		return index
	}
	return next
}
