package solution

func hurdleRace(jump int32, heights []int32) int32 {
    var highest int32
	for _, h := range heights {
		if h > highest {
			highest = h
		}
	}

	if jump > highest {
		return 0
	} else {
		return highest - jump
	}
}