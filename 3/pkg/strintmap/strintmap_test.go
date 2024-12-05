package strintmap

import "testing"

func TestStringIntMap(t *testing.T) {
	impl := New()

	tests := []struct {
		name     string
		addKey   string
		addValue int
		expected map[string]int
	}{
		{
			name:     "Add one element",
			addKey:   "one",
			addValue: 1,
			expected: map[string]int{
				"one": 1,
			},
		},
		{
			name:     "Add more elements",
			addKey:   "two",
			addValue: 2,
			expected: map[string]int{
				"one": 1,
				"two": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl.Add(tt.addKey, tt.addValue)

			for key, expectedValue := range tt.expected {
				actualValue, exists := impl.Get(key)
				if !exists {
					t.Errorf("key %s does not exist", key)
				}
				if actualValue != expectedValue {
					t.Errorf("for key %s, expected %d, got %d", key, expectedValue, actualValue)
				}
			}
		})
	}

	t.Run("Remove element", func(t *testing.T) {
		impl.Remove("two")

		if _, exists := impl.Get("two"); exists {
			t.Error("expected 'two' to be removed")
		}

		if val, exists := impl.Get("one"); !exists || val != 1 {
			t.Errorf("expected 'one' to exist with value 1, got %d", val)
		}
	})

	t.Run("Check existence", func(t *testing.T) {
		if !impl.Exists("one") {
			t.Error("expected 'one' to exist")
		}
		if impl.Exists("two") {
			t.Error("expected 'two' to be removed")
		}
	})

	t.Run("Copy map", func(t *testing.T) {
		copiedMap := impl.Copy()

		if len(copiedMap) != len(impl.sim) {
			t.Errorf("expected copied map length %d, got %d", len(impl.sim), len(copiedMap))
		}

		for key, value := range impl.sim {
			if copiedValue, exists := copiedMap[key]; !exists || copiedValue != value {
				t.Errorf("for key %s, expected %d in copied map, got %d", key, value, copiedValue)
			}
		}

		impl.Add("three", 3)
		if copiedMap["three"] == 3 {
			t.Error("copied map should not reflect changes in original map")
		}
	})

	t.Run("Get value", func(t *testing.T) {
		val, exists := impl.Get("one")
		if !exists || val != 1 {
			t.Errorf("expected 'one' to have value 1, got %d", val)
		}

		_, exists = impl.Get("four")
		if exists {
			t.Error("expected 'four' to not exist")
		}
	})
}
