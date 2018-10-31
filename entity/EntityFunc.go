package entity

import "strconv"

func IsParticipatorinList(name string, participators []string) bool {
	for _, j := range participators {
		if name == j {
			return true
		}
	}
	return false
}

func IsParticipatorExist(name string, participators []User) bool {
	for _, j := range participators {
		if name == j.Name {
			return true
		}
	}
	return false
}

func IsParticipatorExistinMeeting(name string, meeting Meeting) bool {
	if name == meeting.Sponsor {
		return true
	}
	for _, j := range meeting.Participators {
		if name == j {
			return true
		}
	}
	return false
}

func IsParticipatorAvailable(name string, all_meetings []Meeting, current_meeting Meeting) bool {
	start_date := current_meeting.Startdate
	end_date := current_meeting.Enddate
	for _, j := range all_meetings {
		if IsParticipatorExistinMeeting(name, j) {
			if Compare(j.Startdate, end_date) >= 0 || Compare(start_date, j.Enddate) >= 0 {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func Convert(date_string []string) (Date, bool) {
	year, flag1 := strconv.Atoi(date_string[0])
	month, flag2 := strconv.Atoi(date_string[1])
	day, flag3 := strconv.Atoi(date_string[2])
	hour, flag4 := strconv.Atoi(date_string[3])
	minute, flag5 := strconv.Atoi(date_string[4])

	var date Date
	if flag1 == nil && flag2 == nil && flag3 == nil && flag4 == nil && flag5 == nil {
		date.Year = year
		date.Month = month
		date.Day = day
		date.Hour = hour
		date.Minute = minute
		return date, true
	} else {
		return date, false
	}

}

func Check_participators(sponsor_name string, participators []string, all_users []User, all_meetings []Meeting, s_date Date, e_date Date) ([]string, bool) {
	var valid_participators []string
	var temp_meeting Meeting
	temp_meeting.Startdate = s_date
	temp_meeting.Enddate = e_date

	for _, j := range participators {
		if IsParticipatorExist(j, all_users) && j != sponsor_name {
			if IsParticipatorAvailable(j, all_meetings, temp_meeting) {
				valid_participators = append(valid_participators, j)
			}
		}
	}
	if len(valid_participators) == 0 {
		return valid_participators, false
	} else {
		return valid_participators, true
	}
}

func Check_title(meeting_title string, all_meetings []Meeting) bool {
	for _, j := range all_meetings {
		if j.Title == meeting_title {
			return false
		}
	}
	return true
}

func Check_date(date1 Date, date2 Date) bool {
	if IsValid(date1) && IsValid(date2) {
		if Compare(date2, date1) >= 0 {
			return true
		}
	}
	return false
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
		} else { //not leap
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
		(second.Year*366*24*60 + second.Month*30*24*60 + second.Day*24*60 + second.Hour*60 + second.Minute)

}
