package worker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getWorker() *Worker {
	return NewWorker(1, 9000, BonusOptions("1", 25, 10), StorageOptions("1", 10))
}

func TestGetBonusStorage(t *testing.T) {
	w := getWorker()

	t.Run("get bonus", func(t *testing.T) {
		t.Helper()
		var want int32 = 10
		got := w.GetBonusStorage(123456)
		assert.Equal(t, want, got)
	})

	t.Run("can't get bonus", func(t *testing.T) {
		t.Helper()
		var want int32 = 0
		got := w.GetBonusStorage(234561)
		assert.Equal(t, want, got)
	})
}

func TestGetBonusProductivity(t *testing.T) {
	w := getWorker()

	t.Run("get bonus", func(t *testing.T) {
		t.Helper()
		var want int32 = 25
		got := w.GetBonusProductivity(123456)
		assert.Equal(t, want, got)
	})

	t.Run("can't get bonus", func(t *testing.T) {
		t.Helper()
		var want int32 = 0
		got := w.GetBonusProductivity(234561)
		assert.Equal(t, want, got)
	})
}

func TestProduce(t *testing.T) {
	tests := map[string]struct {
		item           int32
		minutes        int32
		wantminutes    int32
		wantproduction int32
	}{
		"50 minutes, use bonus": {
			item:           123456,
			minutes:        50,
			wantminutes:    50,
			wantproduction: 1300,
		},
		"50 minutes, not use bonus": {
			item:           23456,
			minutes:        50,
			wantminutes:    50,
			wantproduction: 50,
		},
		"100 minutes, use bonus": {
			item:           114514,
			minutes:        100,
			wantminutes:    100,
			wantproduction: 2600,
		},
		"100 minutes, not use bonus": {
			item:           919,
			minutes:        100,
			wantminutes:    90,
			wantproduction: 90,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			w := getWorker()
			minutes, production := w.Produce(test.item, test.minutes)
			assert.Equal(t, test.wantminutes, minutes)
			assert.Equal(t, test.wantproduction, production)
		})
	}
}
