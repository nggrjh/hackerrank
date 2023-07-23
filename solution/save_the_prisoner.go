package solution

// https://www.hackerrank.com/challenges/save-the-prisoner/problem
func saveThePrisoner(prisoner, candy, chair int32) int32 {
	if repeat := candy / prisoner; repeat < 1 {
		return (chair - 1) + (candy - (prisoner * repeat))
	}

	candy -= (prisoner - chair + 1)
	leftover := candy - (prisoner * (candy / prisoner))

	if leftover == 0 {
		return prisoner
	} else {
		return leftover
	}
}
