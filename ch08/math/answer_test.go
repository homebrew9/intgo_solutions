package math
import "testing"

func TestMath(t *testing.T) {
    cases := []struct {
        xs       []float64
        max, min float64
    }{
         {
             xs: []float64{3, 5, 2, 1, 7, 9},
             max: 9,
             min: 1,
         },
         {
             xs: []float64{},
             max: 0,
             min: 0,
         },
    }
    for _, c := range cases {
        max := Max(c.xs)
        if max != c.max {
            t.Errorf("expected %f got %f", c.max, max)
        }
        min := Min(c.xs)
        if max != c.max {
            t.Errorf("expected %f got %f", c.min, min)
        }
    }
}

