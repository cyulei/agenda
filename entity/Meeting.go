package entity

type Meeting struct {
	sponsor       string
	participators []string
	startdate     Date
	enddate       Date
	title         string
}

func GetSponsor(m Meeting) string {
	return m.sponsor
}
func GetParticipators(m Meeting) []string {
	return m.participators
}
func GetStartdate(m Meeting) Date {
	return m.startdate
}
func GetEnddate(m Meeting) Date {
	return m.enddate
}
func GetTitle(m Meeting) string {
	return m.title
}
