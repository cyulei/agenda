package entity

type Date struct {
	Year, Month, Day, Hour, Minute int
}

func IsValid(d Date) bool {
	dayOfMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if d.Year <= 0 || d.Year > 9999 {
		return false
	}
	if d.Month <= 0 || d.Month > 12 {
		return false
	}

	if d.Month == 2 {
		if d.Day <= 0 {
			return false
		}
		if d.Year%4 == 0 && d.Year%100 != 0 || d.Year%400 == 0 {
			if d.Day > 29 {
				return false
			}
		} else {
			if d.Day > 28 {
				return false
			}
		}
	}
	if d.Month <= 0 || d.Month > 12 {
		return false
	}

	if dayOfMonth[d.Month-1] < d.Day || d.Day <= 0 {
		return false
	}
	return true
}
func Compare(first Date, second Date) int {
	return first.Year*366*24*60 + first.Month*30*24*60 + first.Day*24*60 + first.Hour*60 + first.Minute -
		(second.Year*366*24*60 + second.Month*30*24*60 + second.Day*24*60 + second.Hour*60 + first.Minute)

}
