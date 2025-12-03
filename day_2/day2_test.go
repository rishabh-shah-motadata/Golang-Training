package day2

import (
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name  string
		marks []float64
		want  float64
	}{
		{"No Marks", []float64{}, 0.0},
		{"Single Mark", []float64{90}, 90.0},
		{"Multiple Marks", []float64{80, 90, 100}, 90.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s := student{Marks: tc.marks}
			got := s.average()
			if got != tc.want {
				t.Errorf("expected %.2f got %.2f", tc.want, got)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		name  string
		marks []float64
		want  string
	}{
		{"Grade A", []float64{95, 92, 90}, "A"},
		{"Grade B", []float64{85, 80, 82}, "B"},
		{"Grade C", []float64{70, 75, 72}, "C"},
		{"Grade D", []float64{60, 65, 68}, "D"},
		{"Grade F", []float64{40, 55, 50}, "F"},
		{"No Marks -> F", []float64{}, "F"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s := student{Marks: tc.marks}
			got := s.grade()
			if got != tc.want {
				t.Errorf("expected %q got %q", tc.want, got)
			}
		})
	}
}
