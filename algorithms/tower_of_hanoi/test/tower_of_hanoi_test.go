package tower_of_hanoi

import (
	dsa "algorithms/tower_of_hanoi/src"
	"testing"
)

func TestHanoiL2R(t *testing.T) {
	for i := 1; i < 10; i++ {
		res1 := dsa.Hanoi(i, -1, 1)
		res2 := dsa.Hanoi_v2(i, -1, 1)
		if res1 != res2 {
			t.Error("HanoiL2R failed")
			break
		}
	}
	t.Log("HanoiL2R succeed")
}

func TestHanoiL2M(t *testing.T) {
	for i := 1; i < 10; i++ {
		res1 := dsa.Hanoi(i, -1, 0)
		res2 := dsa.Hanoi_v2(i, -1, 0)
		if res1 != res2 {
			t.Error("HanoiL2M failed")
			break
		}
	}
	t.Log("HanoiL2M succeed")
}

func TestHanoiM2R(t *testing.T) {
	for i := 1; i < 10; i++ {
		res1 := dsa.Hanoi(i, 0, 1)
		res2 := dsa.Hanoi_v2(i, 0, 1)
		if res1 != res2 {
			t.Error("HanoiM2R failed")
			break
		}
	}
	t.Log("HanoiM2R succeed")
}

func TestHanoiR2L(t *testing.T) {
	for i := 1; i < 10; i++ {
		res1 := dsa.Hanoi(i, 1, -1)
		res2 := dsa.Hanoi_v2(i, 1, -1)
		if res1 != res2 {
			t.Error("HanoiR2L failed")
			break
		}
	}
	t.Log("HanoiR2L succeed")
}

func TestHanoiR2M(t *testing.T) {
	for i := 1; i < 10; i++ {
		res1 := dsa.Hanoi(i, 0, 1)
		res2 := dsa.Hanoi_v2(i, 0, 1)
		if res1 != res2 {
			t.Error("HanoiR2M failed")
			break
		}
	}
	t.Log("HanoiR2M succeed")
}

func TestHanoiM2L(t *testing.T) {
	for i := 1; i < 10; i++ {
		res1 := dsa.Hanoi(i, 0, -1)
		res2 := dsa.Hanoi_v2(i, 0, -1)
		if res1 != res2 {
			t.Error("HanoiM2L failed")
			break
		}
	}
	t.Log("HanoiM2L succeed")
}
