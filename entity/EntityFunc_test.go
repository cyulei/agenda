package entity

import (
	"reflect"
	"testing"
)

func TestIsParticipatorinList(t *testing.T) {
	type args struct {
		name          string
		participators []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParticipatorinList(tt.args.name, tt.args.participators); got != tt.want {
				t.Errorf("IsParticipatorinList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasUser(t *testing.T) {
	type args struct {
		name  string
		users []User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasUser(tt.args.name, tt.args.users); got != tt.want {
				t.Errorf("HasUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteEmptyMeeting(t *testing.T) {
	type args struct {
		meetings []Meeting
	}
	tests := []struct {
		name string
		args args
		want []Meeting
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteEmptyMeeting(tt.args.meetings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteEmptyMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsParticipatorExist(t *testing.T) {
	type args struct {
		name          string
		participators []User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParticipatorExist(tt.args.name, tt.args.participators); got != tt.want {
				t.Errorf("IsParticipatorExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsParticipatorExistinMeeting(t *testing.T) {
	type args struct {
		name    string
		meeting Meeting
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParticipatorExistinMeeting(tt.args.name, tt.args.meeting); got != tt.want {
				t.Errorf("IsParticipatorExistinMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsParticipatorAvailable(t *testing.T) {
	type args struct {
		name            string
		all_meetings    []Meeting
		current_meeting Meeting
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParticipatorAvailable(tt.args.name, tt.args.all_meetings, tt.args.current_meeting); got != tt.want {
				t.Errorf("IsParticipatorAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		date_string []string
	}
	tests := []struct {
		name  string
		args  args
		want  Date
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Convert(tt.args.date_string)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Convert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCheck_participators(t *testing.T) {
	type args struct {
		sponsor_name  string
		participators []string
		all_users     []User
		all_meetings  []Meeting
		s_date        Date
		e_date        Date
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Check_participators(tt.args.sponsor_name, tt.args.participators, tt.args.all_users, tt.args.all_meetings, tt.args.s_date, tt.args.e_date)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Check_participators() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Check_participators() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCheck_title(t *testing.T) {
	type args struct {
		meeting_title string
		all_meetings  []Meeting
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check_title(tt.args.meeting_title, tt.args.all_meetings); got != tt.want {
				t.Errorf("Check_title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheck_date(t *testing.T) {
	type args struct {
		date1 Date
		date2 Date
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check_date(tt.args.date1, tt.args.date2); got != tt.want {
				t.Errorf("Check_date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	type args struct {
		d Date
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.d); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		first  Date
		second Date
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
