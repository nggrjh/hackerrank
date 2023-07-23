package solution

// FIXME: https://www.hackerrank.com/challenges/save-the-prisoner/problem
func saveThePrisoner(prisoner, candy, start int32) int32 {
	return (start - 1) + (candy - (prisoner * (candy / prisoner)))
}