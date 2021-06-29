package dice

import "testing"

func TestNewDice(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		sides := 6
		d, err := New(sides)
		if err != nil {
			t.Errorf("New failed with error: %s", err)
		}
		if d.sides != sides {
			t.Errorf("Expected d.sides=%d, got %d", sides, d.sides)
		}
	})
	t.Run("Zero error", func(t *testing.T) {
		sides := 0
		_, err := New(sides)
		if err == nil {
			t.Errorf("New should have failed with sides = %d", sides)
		}
	})
	t.Run("Negative error", func(t *testing.T) {
		sides := -1
		_, err := New(sides)
		if err == nil {
			t.Errorf("New should have failed with sides = %d", sides)
		}
	})
}

func TestDice_Roll(t *testing.T) {
	sides := 6
	d, _ := New(sides)
	for i := 0; i < 100; i++ {
		r := d.Roll()
		if r > sides || r < 1 {
			t.Errorf("Roll should be > %d and < 1, got %d", sides, r)
		}
	}
}

func TestDice_RollN(t *testing.T) {
	sides := 6
	d, _ := New(sides)

	t.Run("Correct", func(t *testing.T) {
		n := 3
		rs := d.RollN(n)
		if len(rs) != n {
			t.Errorf("RollN should return %d ints, got %d", n, len(rs))
		}
	})
	t.Run("N is zero", func(t *testing.T) {
		n := 0
		l := 0
		rs := d.RollN(n)
		if len(rs) != l {
			t.Errorf("RollN should return %d ints, got %d", l, len(rs))
		}
	})
	t.Run("N is negative", func(t *testing.T) {
		n := -1
		l := 0
		rs := d.RollN(n)
		if len(rs) != l {
			t.Errorf("RollN should return %d ints, got %d", l, len(rs))
		}
	})
}

func TestDice_SumN(t *testing.T) {
	sides, n := 3, 3
	max, min := sides*n, n
	d, _ := New(sides)
	for i := 0; i < 100; i++ {
		s, _ := d.SumN(n)
		if s > max || s < min {
			t.Errorf("SumN should be > %d and < %d, got %d", min, max, s)
		}
	}
}

func TestDice_SumNK(t *testing.T) {
	sides := 3
	d, _ := New(sides)

	t.Run("Correct", func(t *testing.T) {
		n, k := 3, 2
		max, min := sides*k, k
		for i := 0; i < 100; i++ {
			s, rs := d.SumNK(n, k)
			if s > max || s < min {
				t.Errorf("SumNK should be > %d and < %d, got %d", min, max, s)
			}
			if len(rs) != n {
				t.Errorf("SumNK should return %d ints, got %d", n, len(rs))
			}
		}
	})
	t.Run("Bigger k than n", func(t *testing.T) {
		n, k := 2, 3
		_, rs := d.SumNK(n, k)
		if len(rs) != n {
			t.Errorf("SumNK should return %d ints, got %d", n, len(rs))
		}
	})

	fError := func(n, k, length int, t *testing.T) {
		r, rs := d.SumNK(n, k)
		if r != 0 {
			t.Errorf("SumNK should return 0, got %d", r)
		}
		if len(rs) != length {
			t.Errorf("SumNK should return a list len = %d, got len = %d", length, len(rs))
		}
	}

	t.Run("N is zero", func(t *testing.T) { fError(0, 1, 0, t) })
	t.Run("N is negative", func(t *testing.T) { fError(-1, 1, 0, t) })
	t.Run("K is zero", func(t *testing.T) { fError(1, 0, 1, t) })
	t.Run("K is negative", func(t *testing.T) { fError(1, -1, 1, t) })
}

func BenchmarkDice_Roll(b *testing.B) {
	sides := 6
	d, _ := New(sides)
	for i := 0; i < b.N; i++ {
		_ = d.Roll()
	}
}

func benchmarkDice_RollN(n int, b *testing.B) {
	sides := 6
	d, _ := New(sides)
	for i := 0; i < b.N; i++ {
		_ = d.RollN(n)
	}
}

func BenchmarkDice_RollN1(b *testing.B)   { benchmarkDice_RollN(1, b) }
func BenchmarkDice_RollN10(b *testing.B)  { benchmarkDice_RollN(10, b) }
func BenchmarkDice_RollN100(b *testing.B) { benchmarkDice_RollN(100, b) }

func benchmarkDice_SumNK(n, k int, b *testing.B) {
	sides := 6
	d, _ := New(sides)
	for i := 0; i < b.N; i++ {
		_, _ = d.SumNK(n, k)
	}
}

func BenchmarkDice_SumN1(b *testing.B)   { benchmarkDice_SumNK(1, 1, b) }
func BenchmarkDice_SumN10(b *testing.B)  { benchmarkDice_SumNK(10, 10, b) }
func BenchmarkDice_SumN100(b *testing.B) { benchmarkDice_SumNK(100, 100, b) }

func BenchmarkDice_SumN100K1(b *testing.B)  { benchmarkDice_SumNK(100, 1, b) }
func BenchmarkDice_SumN100K10(b *testing.B) { benchmarkDice_SumNK(100, 10, b) }
func BenchmarkDice_SumN100K50(b *testing.B) { benchmarkDice_SumNK(100, 50, b) }
