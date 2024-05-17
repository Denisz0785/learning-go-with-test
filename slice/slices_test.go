package slice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of array", func(t *testing.T) {
		arr := [...]int{1, 2, 3, 4, 5}

		expected := 15
		sum := Sum(arr)

		if sum != expected {
			t.Errorf("expected %d got %d", expected, sum)
		}
	})

	t.Run("sum of slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 6}
		expected := 16

		sum := SumSlice(src)
		if sum != expected {
			t.Errorf("expected %d got %d", expected, sum)
		}

	})
}

func TestSumAll(t *testing.T) {
	src1 := []int{1, 2, 8}
	src2 := []int{3, 4, 9}
	expected := []int{11, 16}
	

	got := SumAll(src1, src2)

	// if !reflect.DeepEqual(expected, got) {
	// 	t.Errorf("expected %d got %d", expected, got)
	// }
	newFunction(expected, got, t)
}

func newFunction(expected []int, got []int, t *testing.T) {
	if !slices.Equal(expected, got) {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestSumTails(t *testing.T) {

	compareSum := func(t testing.TB, want, got []int) {
		t.Helper()
		if !slices.Equal(want, got) {
			t.Errorf("expected %v got %v", want, got)
		}
	}
	t.Run("sum tails", func(t *testing.T) {
		want := []int{13, 11}

		got := SumTails([]int{3, 6, 7}, []int{3, 6, 5})
		compareSum(t, want, got)
	})

	t.Run("sum tails with one empty slice", func(t *testing.T) {
		
		want := []int{0, 11}

		got := SumTails([]int{}, []int{3, 6, 5})
		compareSum(t, want, got)
	})
}

 
