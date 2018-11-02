package entity

import (
	"testing"
)

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
		{
			name: "test1",
			args: args{name: "hh", users: []User{User{Name: "hh", Password: "12345678", Email: "asd@qw.qw", Phone: "12345678977"}}},
			want: true,
		},
		{
			name: "test2",
			args: args{name: "hh1", users: []User{User{Name: "hh1", Password: "12345678", Email: "asd@qw.qw", Phone: "12345678977"}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasUser(tt.args.name, tt.args.users); got != tt.want {
				t.Errorf("HasUser() = %v, want %v", got, tt.want)
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
		{
			name: "test1",
			args: args{name: "hh", participators: []User{User{Name: "hh", Password: "12345678", Email: "asd@qw.qw", Phone: "12345678977"}}},
			want: true,
		},
		{
			name: "test2",
			args: args{name: "hh1", participators: []User{User{Name: "hh1", Password: "12345678", Email: "asd@qw.qw", Phone: "12345678977"}}},
			want: true,
		},
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
		{
			name: "test1",
			args: args{name: "hh", meeting: Meeting{Sponsor: "hh", Participators: []string{"kk"}, Startdate: Date{Year: 2011, Month: 1, Day: 2, Hour: 2, Minute: 30}, Enddate: Date{Year: 2012, Month: 1, Day: 2, Hour: 2, Minute: 30}, Title: "1"}},
			want: true,
		},
		{
			name: "test2",
			args: args{name: "hh1", meeting: Meeting{Sponsor: "hh1", Participators: []string{"sg"}, Startdate: Date{Year: 2011, Month: 1, Day: 2, Hour: 2, Minute: 30}, Enddate: Date{Year: 2012, Month: 1, Day: 2, Hour: 2, Minute: 30}, Title: "1"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParticipatorExistinMeeting(tt.args.name, tt.args.meeting); got != tt.want {
				t.Errorf("IsParticipatorExistinMeeting() = %v, want %v", got, tt.want)
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
		{
			name: "test1",
			args: args{meeting_title: "1", all_meetings: []Meeting{Meeting{Sponsor: "hh", Participators: []string{"kk"}, Startdate: Date{Year: 2011, Month: 1, Day: 2, Hour: 2, Minute: 30}, Enddate: Date{Year: 2012, Month: 1, Day: 2, Hour: 2, Minute: 30}, Title: "1"}}},
			want: false,
		},
		{
			name: "test2",
			args: args{meeting_title: "2", all_meetings: []Meeting{Meeting{Sponsor: "hh1", Participators: []string{"sg"}, Startdate: Date{Year: 2011, Month: 1, Day: 2, Hour: 2, Minute: 30}, Enddate: Date{Year: 2012, Month: 1, Day: 2, Hour: 2, Minute: 30}, Title: "2"}}},
			want: false,
		},
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
		{
			name: "test1",
			args: args{date1: Date{Year: 2011, Month: 1, Day: 2, Hour: 2, Minute: 30}, date2: Date{Year: 2012, Month: 1, Day: 2, Hour: 2, Minute: 30}},
			want: true,
		},
		{
			name: "test2",
			args: args{date1: Date{Year: 2010, Month: 1, Day: 2, Hour: 2, Minute: 30}, date2: Date{Year: 2011, Month: 1, Day: 2, Hour: 2, Minute: 30}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check_date(tt.args.date1, tt.args.date2); got != tt.want {
				t.Errorf("Check_date() = %v, want %v", got, tt.want)
			}
		})
	}
}
