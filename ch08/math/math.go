package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
    if len(xs) == 0 {
        return 0
    }
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Finds the minimum of a series of numbers
func Min(xs []float64) float64 {
    var min_value float64
    for i, value := range xs {
        if (i == 0) || (value < min_value) {
            min_value = value
        }
    }
    return min_value
}

// Finds the maximum of a series of numbers
func Max(xs []float64) float64 {
    var max_value float64
    for i, value := range xs {
        if (i == 0) || (value > max_value) {
            max_value = value
        }
    }
    return max_value
}

