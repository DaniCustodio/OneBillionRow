package measurements

import (
	"reflect"
	"sort"
	"testing"
)

func TestOpenFile(t *testing.T) {
	t.Run("open a valid file", func(t *testing.T) {
		file, err := openFile("measurements_test.txt")

		if err != nil {
			t.Errorf("Expected nil, but got %v", err)
		}

		if file == nil {
			t.Fatalf("Expected a file, but got nil")
		}

		defer file.Close()
	})

	t.Run("open a non-existent file", func(t *testing.T) {
		_, err := openFile("non-existent.txt")
		if err == nil {
			t.Error("Expected an error for a non-existent file, but got nil")
		}
	})
}

func TestMeasurement(t *testing.T) {
	t.Run("get a measurement", func(t *testing.T) {
		expectedLocation := "Rio de Janeiro"
		expectedTemperature := 30.5

		location, temperature, err := getMeasurement("Rio de Janeiro;30.5")

		if err != nil {
			t.Errorf("Expected nil, but got %v", err)
		}

		if location != expectedLocation {
			t.Errorf("Expected %v, but got %v", expectedLocation, location)
		}

		if temperature != expectedTemperature {
			t.Errorf("Expected %v, but got %v", expectedTemperature, temperature)
		}

	})

	t.Run("get a map of measurements", func(t *testing.T) {
		expected := map[string]Measurement{
			"Rio de Janeiro": {
				Min:   19.8,
				Max:   31.2,
				Sum:   76.5,
				Count: 3,
			},
			"São Paulo": {
				Min:   18.7,
				Max:   21.5,
				Sum:   40.2,
				Count: 2,
			},
			"Noruega": {
				Min:   -7.5,
				Max:   -7.5,
				Sum:   -7.5,
				Count: 1,
			},
		}

		data, err := GetMeasurements("measurements_test.txt")

		if err != nil {
			t.Errorf("Expected nil, but got %v", err)
		}

		if !reflect.DeepEqual(data, expected) {
			t.Errorf("Expected %v to be equal %v", data, expected)
		}
	})

	t.Run("get measurements sorted by location", func(t *testing.T) {
		data := map[string]Measurement{
			"Rio de Janeiro": {
				Min:   19.8,
				Max:   31.2,
				Sum:   76.5,
				Count: 3,
			},
			"São Paulo": {
				Min:   18.7,
				Max:   21.5,
				Sum:   40.2,
				Count: 2,
			},
			"Noruega": {
				Min:   -7.5,
				Max:   -7.5,
				Sum:   -7.5,
				Count: 1,
			},
		}

		sortedData := SortMeasurementsByLocation(&data)

		if len(sortedData) <= 0 || !sort.IsSorted(sort.StringSlice(sortedData)) {
			t.Errorf("Expected %v to be sorted", sortedData)
		}
	})
}
