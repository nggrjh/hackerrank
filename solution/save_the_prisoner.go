package solution

// https://www.hackerrank.com/challenges/save-the-prisoner/problem
func saveThePrisoner(prisoner, candy, chair int32) int32 {
	firstRoundCandy := candy - (prisoner - chair + 1)
	if firstRoundCandy == 0 {
		return prisoner
	} else if firstRoundCandy < 0 {
		return chair - 1 + candy
	}

	leftover := firstRoundCandy - (prisoner * (firstRoundCandy / prisoner))

	if leftover == 0 {
		return prisoner
	} else {
		return leftover
	}
}
