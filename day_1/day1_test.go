package day1

import "testing"

func TestTimeGreeting(t *testing.T) {
	tests := []struct {
		hour int
		want string
	}{
		{0, "Good morning"},
		{8, "Good morning"},
		{12, "Good afternoon"},
		{16, "Good afternoon"},
		{18, "Good evening"},
		{20, "Good evening"},
		{22, "Good night"},
		{23, "Good night"},
	}

	for _, tc := range tests {
		result := getTimeGreeting(tc.hour)
		if result != tc.want {
			t.Errorf("For hour %d, expected %s but got %s\n", tc.hour, tc.want, result)
		}
	}
}
