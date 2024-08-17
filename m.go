package main

import (
	"fmt"
	"main/measurements"
	"time"
)

func main() {
	start := time.Now()
	data, err := measurements.GetMeasurements("measurements.txt")

	if err != nil {
		panic(err)
	}

	sortedLocations := measurements.SortMeasurementsByLocation(&data)

	fmt.Printf("{")
	for _, name := range sortedLocations {
		measurement := data[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")

	fmt.Printf("Execution time: %v\n", time.Since(start))
}
