package entity

type Meeting struct {
	Sponsor       string
	Participators []string
	Startdate     Date
	Enddate       Date
	Title         string
}
