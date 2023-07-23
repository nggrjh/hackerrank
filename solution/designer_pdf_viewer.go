package solution

func designerPdfViewer(heights []int32, word string) int32 {
	var tallest int32
	for _, w := range word {
		if r := rune(w) - 97; heights[r] > tallest {
			tallest = heights[r]
		}
	}
	return tallest * int32(len(word))
}
