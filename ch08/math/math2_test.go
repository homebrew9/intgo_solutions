package math
import "testing"

type teststruct struct {
    values []float64
    min float64
    max float64
}

var my_tests = []teststruct {
    { []float64{1,2}, 1, 2},
    { []float64{1,1,1,1,1,1}, 1, 1},
    { []float64{-1,1}, -1, 1},
    { []float64{}, 0, 0},
}

func TestMax(t *testing.T) {
    for _, pair := range my_tests {
        v := Max(pair.values)
        if v != pair.max {
            t.Error(
                "For", pair.values,
                "expected", pair.max,
                "got", v,
            )
        }
    }
}

func TestMin(t *testing.T) {
    for _, pair := range my_tests {
        v := Min(pair.values)
        if v != pair.min {
            t.Error(
                "For", pair.values,
                "expected", pair.min,
                "got", v,
            )
        }
    }
}

