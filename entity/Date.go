package entity

type Date struct {
	Year, Month, Day, Hour, Minute int
}

func IsValid(d Date) bool {
	return true
}
