package main

import (
	"reflect"
	"testing"
)

func Test_calc(t *testing.T) {

	update(NAME_FILE, &mvs)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want    ResponseAnaytics
		wantErr bool
	}{

		{
			name: "teste 01",
			want: ResponseAnaytics{
				Summary: Summary{
					TotalMoviesProcessed: 4723,
					GlobalAverageRating:  6.095807749311872,
					TotalRevenue:         392328507332,
					TotalBudget:          138976526519,
					AverageRevenue:       83067649,
					AverageBudget:        29425476,
					AverageProfit:        53642172,
					OverallROI:           9536495,
				},
				TopPerformingYear: TopPerformingYear{
					Year:         2012,
					TotalRevenue: 24141615246,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := calc(mvs)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("calc() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("calc() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calc() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_calcStudios(t *testing.T) {

	update(NAME_FILE, &mvs)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want int
	}{
		{
			name: "deve retornar a quantidade correta de studios",
			want: 1281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcStudios(mvs)
			// TODO: update the condition below to compare got with tt.want.
			if len(got) != tt.want {
				t.Errorf("got %d, want %d", len(got), tt.want)
			}
		})
	}
}

func Test_calcGenre(t *testing.T) {

	update(NAME_FILE, &mvs)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want int
	}{
		{
			name: "deve retornar total de generos",
			want: 21,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcGenre(mvs)
			// TODO: update the condition below to compare got with tt.want.
			if len(got) != tt.want {
				t.Errorf("got %d, want %d", len(got), tt.want)
			}
		})
	}
}
