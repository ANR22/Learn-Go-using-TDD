package structsmethodsinterface

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("calculate area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{width: 12.0, height: 6.0}
		want := 72.0
		checkArea(t, rectangle, want)
	})

	t.Run("calculate area of circle", func(t *testing.T) {
		circle := Circle{radius: 10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{width: 12.0, height: 6.0}, want: 72.0},
		{shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{base: 12.0, height: 6.0}, want: 36.0},
	}
	for _, tc := range areaTest {
		checkArea(t, tc.shape, tc.want)
	}
}
