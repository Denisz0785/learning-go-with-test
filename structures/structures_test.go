package structures

import "testing"

func TestPeriemetr(t *testing.T) {
	r := Rectangle{
		width: 4.0,
		hight: 6.0,
	}
	want := 20.0
	p := Periemetr(r)

	if want != p {
		t.Errorf("want %.2f got %.2f", want, p)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	}

	t.Run("rectagle area", func(t *testing.T) {
		r := &Rectangle{
			width: 10.0,
			hight: 5.0,
		}
		checkArea(t, r, 50.0)

	})
	t.Run("circle area", func(t *testing.T) {

		want := 314.1592653589793
		c := &Circle{10.0}
		checkArea(t, c, want)
	})
}
