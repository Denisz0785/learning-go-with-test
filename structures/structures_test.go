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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: &Rectangle{width: 10.0, hight: 5.0}, want: 50.0},
		{name: "Circle", shape: &Circle{radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{base: 12.0, hight: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.shape.Area()

			if tt.want != got {
				t.Errorf("Shape %#v want %g got %g", tt.shape, tt.want, got)
			}
		})

	}

}
