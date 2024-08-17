package measurements

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

var ErrCouldNotOpenFile = errors.New("could not open file")

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("%w: %v - %v", ErrCouldNotOpenFile, filename, err)
	}

	return file, nil
}

func getMeasurement(line string) (string, float64, error) {
	semicolonIndex := strings.Index(line, ";")
	location := line[:semicolonIndex]
	temperature := line[semicolonIndex+1:]

	floatTemperature, err := strconv.ParseFloat(temperature, 64)

	if err != nil {
		return "", 0, fmt.Errorf("could not convert temperature to float: %v", err)
	}

	return location, floatTemperature, nil
}

func GetMeasurements(filename string) (map[string]Measurement, error) {
	file, err := openFile(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		location, temperature, err := getMeasurement(line)
		if err != nil {
			return nil, err
		}

		measurement, ok := data[location]

		if !ok {
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count++
		}

		data[location] = measurement
	}

	return data, nil
}

func SortMeasurementsByLocation(data *map[string]Measurement) []string {
	locations := make([]string, 0, len(*data))
	for name := range *data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	return locations
}
