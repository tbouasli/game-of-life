package game_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/tbouasli/game-of-life/game"
)

const SAMPLES_DIR = "samples"
const GLIDER_DIR = SAMPLES_DIR + "/glider"
const LARGE_DATA_SET_DIR = SAMPLES_DIR + "/large_data_set"

var GliderSamples = []string{
	"1.life",
	"2.life",
	"3.life",
	"4.life",
	"5.life",
}

var LargeDataSetSamples = []string{
	"1.life",
	"10.life",
	"100.life",
}

func TestGlider(t *testing.T) {
	for i, sample := range GliderSamples {
		if i+1 >= len(GliderSamples) {
			continue
		}

		t.Run(fmt.Sprintf("GliderSample %s", sample), func(t *testing.T) {
			currentFile, err := os.Open(GLIDER_DIR + "/" + sample)
			if err != nil {
				t.Fatalf("Failed to open sample file: %v", err)
			}
			defer currentFile.Close()

			grid, err := game.FromLife106Format(bufio.NewScanner(currentFile))
			if err != nil {
				t.Fatalf("Failed to parse sample file: %v", err)
			}

			nextFile := GliderSamples[i+1]
			expected, err := os.Open(GLIDER_DIR + "/" + nextFile)
			if err != nil {
				t.Fatalf("Failed to open expected file: %v", err)
			}
			defer expected.Close()

			expectedGrid, err := game.FromLife106Format(bufio.NewScanner(expected))
			if err != nil {
				t.Fatalf("Failed to parse expected file: %v", err)
			}

			grid = grid.NextGeneration()

			actualOutput := game.ToLife106Format(grid)
			expectedOutput := game.ToLife106Format(expectedGrid)

			if actualOutput != expectedOutput {
				t.Errorf("Sample %s does not match expected %s.\nActual:\n%s\nExpected:\n%s",
					sample, nextFile, actualOutput, expectedOutput)
			}
		})
	}
}

func TestLargeDataSet(t *testing.T) {
	startFile, err := os.Open(LARGE_DATA_SET_DIR + "/" + LargeDataSetSamples[0])
	if err != nil {
		t.Fatalf("Failed to open sample file: %v", err)
	}
	defer startFile.Close()

	grid, err := game.FromLife106Format(bufio.NewScanner(startFile))
	if err != nil {
		t.Fatalf("Failed to parse sample file: %v", err)
	}

	testCases := []struct {
		generations     int
		expectedFileIdx int
		description     string
	}{
		{10, 1, "10 generations"},
		{100, 2, "100 generations"},
	}

	currentGeneration := 0

	for _, cases := range testCases {
		t.Run(cases.description, func(t *testing.T) {
			generationsToRun := cases.generations - currentGeneration

			for range generationsToRun {
				grid = grid.NextGeneration()
			}
			currentGeneration = cases.generations

			expectedFile, err := os.Open(LARGE_DATA_SET_DIR + "/" + LargeDataSetSamples[cases.expectedFileIdx])
			if err != nil {
				t.Fatalf("Failed to open expected file for %s: %v", cases.description, err)
			}

			expectedGrid, err := game.FromLife106Format(bufio.NewScanner(expectedFile))
			expectedFile.Close()

			if err != nil {
				t.Fatalf("Failed to parse expected file for %s: %v", cases.description, err)
			}

			actualOutput := game.ToLife106Format(grid)
			expectedOutput := game.ToLife106Format(expectedGrid)

			if actualOutput != expectedOutput {
				t.Errorf("After %s: actual output does not match expected output from %s",
					cases.description, LargeDataSetSamples[cases.expectedFileIdx])
			}
		})
	}
}
