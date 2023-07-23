package solution

func angryProfessor(attendance int32, students []int32) string {
	var arrived int32
	for _, s := range students {
		if s < 1 {
			arrived++
		}

		if arrived >= attendance {
			return "NO"
		}
	}

	return "YES"
}
