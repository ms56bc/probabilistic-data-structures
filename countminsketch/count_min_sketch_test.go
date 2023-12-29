package countminsketch

import (
	"testing"
)

func TestCountMinSketch_Increment(t *testing.T) {
	width := 1000
	depth := 5
	cms, _ := NewCountMinSketch(width, depth)

	// Increment counts
	items := []string{"apple", "banana", "orange", "apple", "banana", "apple", "apple"}
	for _, item := range items {
		cms.Increment([]byte(item))
	}

	// Test estimated counts
	testCases := []struct {
		item          string
		expectedCount int
	}{
		{"apple", 4},
		{"banana", 2},
		{"orange", 1},
		{"grape", 0}, // Unseen item should return 0
	}

	for _, tc := range testCases {
		t.Run(tc.item, func(t *testing.T) {
			count := cms.Estimate([]byte(tc.item))
			if count != tc.expectedCount {
				t.Errorf("Expected count for %s: %d, got: %d", tc.item, tc.expectedCount, count)
			}
		})
	}
	cms.clear()
}

func TestCountMinSketch_Clear(t *testing.T) {
	width := 1000
	depth := 5
	cms, _ := NewCountMinSketch(width, depth)

	// Increment counts
	items := []string{"apple", "banana", "orange", "apple", "banana", "apple", "apple"}
	for _, item := range items {
		cms.Increment([]byte(item))
	}

	// Clear the sketch
	cms.clear()

	// Test estimated counts after clearing
	testCases := []struct {
		item          string
		expectedCount int
	}{
		{"apple", 0},
		{"banana", 0},
		{"orange", 0},
		{"grape", 0}, // Unseen item should return 0
	}

	for _, tc := range testCases {
		t.Run(tc.item+" (after clear)", func(t *testing.T) {
			count := cms.Estimate([]byte(tc.item))
			if count != tc.expectedCount {
				t.Errorf("Expected count for %s after clearing: 0, got: %d", tc.item, count)
			}
		})
	}
}

func TestCountMinSketch_Update(t *testing.T) {
	width := 1000
	depth := 5
	cms, _ := NewCountMinSketch(width, depth)

	// Increment counts with the first set of items
	itemsSet1 := []string{"apple", "banana", "orange", "apple", "banana", "apple", "apple"}
	for _, item := range itemsSet1 {
		cms.Increment([]byte(item))
	}

	// Update counts with the second set of items
	itemsSet2 := []string{"apple", "banana", "orange", "grape", "kiwi", "banana", "apple", "kiwi"}
	for _, item := range itemsSet2 {
		cms.Increment([]byte(item))
	}

	// Test estimated counts after updating
	updatedTestCases := []struct {
		item          string
		expectedCount int
	}{
		{"apple", 6},      // 4 from the first set + 2 from the second set
		{"banana", 4},     // 2 from the first set + 2 from the second set
		{"orange", 2},     // 1 from the first set + 1 from the second set
		{"grape", 1},      // 0 from the first set + 1 from the second set
		{"kiwi", 2},       // 0 from the first set + 2 from the second set
		{"strawberry", 0}, // Unseen item should return 0
	}

	for _, tc := range updatedTestCases {
		t.Run(tc.item+" (after update)", func(t *testing.T) {
			count := cms.Estimate([]byte(tc.item))
			if count != tc.expectedCount {
				t.Errorf("Expected count for %s after update: %d, got: %d", tc.item, tc.expectedCount, count)
			}
		})
	}
}

// write a test for invalid creation

func TestCountMinSketch_InvalidCreation(t *testing.T) {
	width := 0
	depth := 5
	cms, _ := NewCountMinSketch(width, depth)
	if cms != nil {
		t.Errorf("Expected to get error")
	}
}
