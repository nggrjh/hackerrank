package solution

import (
	"sort"
)

var medians []float64

// FIXME: https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem
func activityNotifications(expenditure []int32, trailingDays int32) int32 {
	medians = make([]float64, len(expenditure))

	var notif int32
	for today := int(trailingDays); today < len(expenditure); today++ {
		buff := evaluateMedian(expenditure, today, int(trailingDays)) * 2
		if float64(expenditure[today]) >= buff {
			notif++
		}
	}
	return notif
}

func isOdd(i int) bool {
	return i%2 == 1
}

func evaluateMedian(expenditure []int32, today, trailingDays int) float64 {
	firstTrailingDay := today - trailingDays

	if firstTrailingDay < 1 {
		medians[today] = calculateMedian(expenditure[firstTrailingDay:today])
		return medians[today]
	}

	if expenditure[today] == expenditure[firstTrailingDay-1] { // value in and out are the same, do not re-calculate
		medians[today] = medians[today-1]
	} else if float64(expenditure[today]) == medians[today-1] { // value in and median are the same, do not re-calculate
		medians[today] = medians[today-1]
	} else {
		medians[today] = calculateMedian(expenditure[firstTrailingDay:today])
	}

	return medians[today]
}

func calculateMedian(expenditure []int32) float64 {
	sort.SliceStable(expenditure, func(i, j int) bool { return expenditure[i] < expenditure[j] })

	ln := len(expenditure)
	mid := ln / 2
	if isOdd(ln) {
		return float64(expenditure[mid])
	} else {
		return float64(expenditure[mid]+expenditure[mid-1]) / 2
	}
}
