package matchers_benchmarks

import (
	"testing"

	"github.com/grobinson-grafana/matchers"
	prom_matchers "github.com/prometheus/alertmanager/pkg/labels"
)

const (
	simpleEquals  = "{foo=\"bar\"}"
	complexEquals = "{foo=\"bar\",bar=\"foo 🙂\",\"baz\"!=qux,qux!=\"baz 🙂\"}"
)

func BenchmarkMatchersSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := matchers.Parse(simpleEquals); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMatchersComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := matchers.Parse(complexEquals); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrometheusSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := prom_matchers.ParseMatchers(simpleEquals); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrometheusComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := prom_matchers.ParseMatchers(complexEquals); err != nil {
			b.Fatal(err)
		}
	}
}
