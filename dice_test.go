package dice

import "testing"

func TestNewDiceCorrect(t *testing.T) {
	sides := 6
	d, err := NewDice(sides)
	if err != nil {
		t.Errorf("NewDice failed with error: %s", err)
	}
	if d.sides != sides {
		t.Errorf("Expected d.sides=%d, got %d", sides, d.sides)
	}
}

func TestNewDiceZero(t *testing.T) {
	sides := 0
	_, err := NewDice(sides)
	if err == nil {
		t.Errorf("NewDice should have failed with sides = %d", sides)
	}
}

func TestNewDiceNegative(t *testing.T) {
	sides := -1
	_, err := NewDice(sides)
	if err == nil {
		t.Errorf("NewDice should have failed with sides = %d", sides)
	}
}

func TestDice_Roll(t *testing.T) {
	sides := 6
	d, _ := NewDice(sides)
	for i := 0; i < 100; i++ {
		r := d.Roll()
		if r > sides || r < 1 {
			t.Errorf("Roll should be > %d and < 1, got %d", sides, r)
		}
	}
}

func TestDice_RollN(t *testing.T) {
	sides, n := 6, 3
	d, _ := NewDice(sides)
	rs := d.RollN(n)
	if len(rs) != n {
		t.Errorf("RollN should return %d ints, got %d", n, len(rs))
	}
}

func TestDice_SumN(t *testing.T) {
	sides, n := 3, 3
	max, min := sides*n, n
	d, _ := NewDice(sides)
	for i := 0; i < 100; i++ {
		s, _ := d.SumN(n)
		if s > max || s < min {
			t.Errorf("SumN should be > %d and < %d, got %d", min, max, s)
		}
	}
}

func TestDice_SumNK(t *testing.T) {
	sides, n, k := 3, 3, 2
	max, min := sides*k, k
	d, _ := NewDice(sides)
	for i := 0; i < 100; i++ {
		s, rs := d.SumNK(n, k)
		if s > max || s < min {
			t.Errorf("SumNK should be > %d and < %d, got %d", min, max, s)
		}
		if len(rs) != n {
			t.Errorf("SumNK should return %d ints, got %d", n, len(rs))
		}
	}
}

func BenchmarkDice_Roll(b *testing.B) {
	sides := 6
	d, _ := NewDice(sides)
	for i := 0; i < b.N; i++ {
		_ = d.Roll()
	}
}

func benchmarkDice_RollN(n int, b *testing.B) {
	sides := 6
	d, _ := NewDice(sides)
	for i := 0; i < b.N; i++ {
		_ = d.RollN(n)
	}
}

func BenchmarkDice_RollN1(b *testing.B)   { benchmarkDice_RollN(1, b) }
func BenchmarkDice_RollN10(b *testing.B)  { benchmarkDice_RollN(10, b) }
func BenchmarkDice_RollN100(b *testing.B) { benchmarkDice_RollN(100, b) }

func benchmarkDice_SumNK(n, k int, b *testing.B) {
	sides := 6
	d, _ := NewDice(sides)
	for i := 0; i < b.N; i++ {
		_, _ = d.SumNK(n, k)
	}
}

func BenchmarkDice_SumN1(b *testing.B)   { benchmarkDice_SumNK(1, 1, b) }
func BenchmarkDice_SumN10(b *testing.B)  { benchmarkDice_SumNK(10, 10, b) }
func BenchmarkDice_SumN100(b *testing.B) { benchmarkDice_SumNK(100, 100, b) }

func BenchmarkDice_SumN100K1(b *testing.B) { benchmarkDice_SumNK(100, 1, b) }
func BenchmarkDice_SumN100K10(b *testing.B) { benchmarkDice_SumNK(100, 10, b) }
func BenchmarkDice_SumN100K50(b *testing.B) { benchmarkDice_SumNK(100, 50, b) }
