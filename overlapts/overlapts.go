package overlapts

type td struct {
	time     int
	duration int
}

// returns the number of overlapping times
func overlapTimes(tds []td) int {

	if len(tds) == 0 {
		return 0
	}

	if len(tds) == 1 {
		return 1
	}

	startTime := make([]int, 1)
	endTime := make([]int, 1)

	startTime[0] = tds[0].time
	endTime[0] = tds[0].time + tds[0].duration
	t := 0

	for i := 1; i < len(tds); i = i + 1 {
		etime := tds[i].time + tds[i].duration
		switch {
		case tds[i].time >= startTime[t] && etime <= endTime[t]:
			// do nothing
		case tds[i].time < endTime[t] && etime > endTime[t]:
			endTime[t] = etime
		case tds[i].time > endTime[t]:
			startTime = append(startTime, tds[i].time)
			endTime = append(endTime, etime)
		}
	}
	return len(startTime)
}
