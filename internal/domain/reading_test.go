package domain

import "testing"

func TestFrameAtRTL(t *testing.T) {
	frame := FrameAt(0, 4, false)
	if frame.Right != 0 || !frame.HasLeft || frame.Left != 1 {
		t.Fatal(frame)
	}
	frame = FrameAt(3, 4, false)
	if frame.Right != 3 || frame.HasLeft {
		t.Fatal(frame)
	}
	frame = FrameAt(1, 4, true)
	if !frame.Single || frame.HasLeft || frame.Right != 1 {
		t.Fatal(frame)
	}
}

func TestStep(t *testing.T) {
	if Step(0, 5, false, true) != 2 {
		t.Fatal("spread next")
	}
	if Step(2, 5, false, false) != 0 {
		t.Fatal("spread prev")
	}
	if Step(4, 5, false, true) != 4 {
		t.Fatal("stay at end")
	}
	if Step(1, 5, true, true) != 2 {
		t.Fatal("single next")
	}
}
