package entity

type Date struct {
	year, month, day, hour, minute int
}

func IsValid(d Date) bool {

}

func GetYear(d Date) int {
	return d.year
}
func GetMonth(d Date) int {
	return d.month
}
func GetDay(d Date) int {
	return d.day
}
func GetHour(d Date) int {
	return d.hour
}
func GetMinute(d Date) int {
	return d.minute
}
