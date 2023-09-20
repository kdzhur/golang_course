package averagegrade

func GetAverageGrade(grades []float64) float64 {
	var sum float64

	if len(grades) == 0 {
		return 0.0
	}

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))
}
