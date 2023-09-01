// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package p503

import (
	"testing"

	"github.com/cloudflare/circl/dh/sidh/internal/common"
)

// Package-level storage for this field element is intended to deter
// compiler optimizations.
var (
	benchmarkFp   common.Fp
	benchmarkFpX2 common.FpX2
	bench_x       = common.Fp{17026702066521327207, 5108203422050077993, 10225396685796065916, 11153620995215874678, 6531160855165088358, 15302925148404145445, 1248821577836769963, 9789766903037985294, 7493111552032041328, 10838999828319306046, 18103257655515297935, 27403304611634}
	bench_y       = common.Fp{4227467157325093378, 10699492810770426363, 13500940151395637365, 12966403950118934952, 16517692605450415877, 13647111148905630666, 14223628886152717087, 7167843152346903316, 15855377759596736571, 4300673881383687338, 6635288001920617779, 30486099554235}
	bench_z       = common.FpX2{1595347748594595712, 10854920567160033970, 16877102267020034574, 12435724995376660096, 3757940912203224231, 8251999420280413600, 3648859773438820227, 17622716832674727914, 11029567000887241528, 11216190007549447055, 17606662790980286987, 4720707159513626555, 12887743598335030915, 14954645239176589309, 14178817688915225254, 1191346797768989683, 12629157932334713723, 6348851952904485603, 16444232588597434895, 7809979927681678066, 14642637672942531613, 3092657597757640067, 10160361564485285723, 240071237}
)

func TestFpCswap(t *testing.T) {
	var one = common.Fp{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var two = common.Fp{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	var x = one
	var y = two

	cswapP503(&x, &y, 0)
	for i := 0; i < FpWords; i++ {
		if (x[i] != one[i]) || (y[i] != two[i]) {
			t.Error("Found", x, "expected", one)
		}
	}

	cswapP503(&x, &y, 1)
	for i := 0; i < FpWords; i++ {
		if (x[i] != two[i]) || (y[i] != one[i]) {
			t.Error("Found", x, "expected", two)
		}
	}
}

func TestFpCmov(t *testing.T) {
	var one = common.Fp{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var two = common.Fp{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	var x = one
	var y = two

	cmovP503(&x, &y, 0)
	for i := 0; i < FpWords; i++ {
		if x[i] != one[i] {
			t.Error("Found", x, "expected", one)
		}
		if y[i] != two[i] {
			t.Error("Found", y, "expected", two)
		}
	}

	cmovP503(&x, &y, 1)
	for i := 0; i < FpWords; i++ {
		if x[i] != two[i] {
			t.Error("Found", x, "expected", two)
		}
		if y[i] != two[i] {
			t.Error("Found", y, "expected", two)
		}
	}
}

// Benchmarking for field arithmetic
func BenchmarkMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mulP503(&benchmarkFpX2, &bench_x, &bench_y)
	}
}

func BenchmarkRdc(b *testing.B) {
	z := bench_z

	// This benchmark actually computes garbage, because
	// rdcP503 mangles its input, but since it's
	// constant-time that shouldn't matter for the benchmarks.
	for n := 0; n < b.N; n++ {
		rdcP503(&benchmarkFp, &z)
	}
}

func BenchmarkAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		addP503(&benchmarkFp, &bench_x, &bench_y)
	}
}

func BenchmarkSub(b *testing.B) {
	for n := 0; n < b.N; n++ {
		subP503(&benchmarkFp, &bench_x, &bench_y)
	}
}

func BenchmarkCswap(b *testing.B) {
	x, y := bench_x, bench_y
	for n := 0; n < b.N; n++ {
		cswapP503(&x, &y, 1)
		cswapP503(&x, &y, 0)
	}
}

func BenchmarkMod(b *testing.B) {
	x := bench_x
	for n := 0; n < b.N; n++ {
		modP503(&x)
	}
}

func BenchmarkX2AddLazy(b *testing.B) {
	x, y, z := bench_z, bench_z, bench_z
	for n := 0; n < b.N; n++ {
		adlP503(&x, &y, &z)
	}
}

func BenchmarkX2SubLazy(b *testing.B) {
	x, y, z := bench_z, bench_z, bench_z
	for n := 0; n < b.N; n++ {
		sulP503(&x, &y, &z)
	}
}
