package worker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getWorkerArray() *WorkerArray {
	w1 := NewWorker(1, 9000, BonusOptions("1", 25, 10), StorageOptions("1", 10))
	w2 := NewWorker(2, 18000, BonusOptions("2", 25, 10), StorageOptions("2", 10))
	w3 := NewWorker(3, 0, BonusOptions("1", 25, 10), StorageOptions("1", 10))
	return NewWorkerArray(w1, w2, w3)
}

func TestExhausted(t *testing.T) {
	wa := getWorkerArray()
	t.Run("is one exhausted", func(t *testing.T) {
		t.Helper()
		want := true
		got := wa.IsOneExhausted()
		assert.Equal(t, want, got)
	})
	t.Run("not all exhausted", func(t *testing.T) {
		t.Helper()
		want := false
		got := wa.IsAllExhausted()
		assert.Equal(t, want, got)
	})

}

func TestGetBonusStorages(t *testing.T) {
	wa := getWorkerArray()

	t.Run("get bonus with item 123", func(t *testing.T) {
		t.Helper()
		want := 20
		got := wa.GetBonusStorage(123)
		assert.Equal(t, want, got)
	})

	t.Run("test bonus with item 234", func(t *testing.T) {
		t.Helper()
		want := 10
		got := wa.GetBonusStorage(234)
		assert.Equal(t, want, got)
	})

	t.Run("test bonus with item 456", func(t *testing.T) {
		t.Helper()
		want := 0
		got := wa.GetBonusStorage(456)
		assert.Equal(t, want, got)
	})
}

func TestGetBonusProductivities(t *testing.T) {
	wa := getWorkerArray()

	t.Run("get bonus with item 123", func(t *testing.T) {
		t.Helper()
		want := 50
		got := wa.GetBonusProductivity(123)
		assert.Equal(t, want, got)
	})

	t.Run("test bonus with item 234", func(t *testing.T) {
		t.Helper()
		want := 25
		got := wa.GetBonusProductivity(234)
		assert.Equal(t, want, got)
	})

	t.Run("test bonus with item 456", func(t *testing.T) {
		t.Helper()
		want := 0
		got := wa.GetBonusProductivity(456)
		assert.Equal(t, want, got)
	})
}

func TestWAProduce(t *testing.T) {
	tests := map[string]struct {
		item           int
		minutes        int
		wantminutes    int
		wantproduction int
	}{
		"50 minutes, use w1 bonus": {
			item:           123456,
			minutes:        50,
			wantminutes:    50,
			wantproduction: 1350,
		},
		"50 minutes, use w2 bonus": {
			item:           23456,
			minutes:        50,
			wantminutes:    50,
			wantproduction: 1350,
		},
		"100 minutes, use w1 bonus": {
			item:           114514,
			minutes:        100,
			wantminutes:    100,
			wantproduction: 2700,
		},
		"100 minutes, use w2 bonus": {
			item:           23456,
			minutes:        100,
			wantminutes:    100,
			wantproduction: 2690,
		},
		"100 minutes, not use bonus": {
			item:           919,
			minutes:        100,
			wantminutes:    100,
			wantproduction: 190,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			wa := getWorkerArray()
			minutes, production := wa.Produce(test.item, test.minutes)
			assert.Equal(t, test.wantminutes, minutes)
			assert.Equal(t, test.wantproduction, production)
		})
	}
}
