package main

import "testing"

func TestReadArguments(t *testing.T) {
	t.Run("verify with correct parameters", func(t *testing.T) {
		want := int64(12)
		args := []string{"12", "12"}
		var load, percentage = readArguments(args)
		if load != want || percentage != want {
			t.Errorf("should have recieved a value back")
		}
	})

	t.Run("verify with incorrect parameters", func(t *testing.T) {
		args := []string{"not-a-number", "not-a-number"}
		var load, percentage = readArguments(args)
		if load != 0 && percentage != 0 {
			t.Errorf("should have recieved zero")
		}
	})
}

func TestCaluatePercentage(t *testing.T) {
	t.Run("calculates the percentage based on input", func(t *testing.T) {
		percentageLoad := int64(78)
		want := float64(0.78)
		var percentage = calcuatePercentage(percentageLoad)
		if want != percentage {
			t.Errorf("should have recieved a value back")
		}
	})
}

func TestCalculateCPUCount(t *testing.T) {
	t.Run("calculates the percentage based on input", func(t *testing.T) {
		percentageLoad := float64(0.78)
		CPU := 4
		want := float64(3.12)
		var percentage = calcuateCPUCount(percentageLoad, CPU)
		print(percentage)
		if want != percentage {
			t.Errorf("should have recieved a value back")
		}
	})
}
