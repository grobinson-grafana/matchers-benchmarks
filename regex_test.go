package matchers_benchmarks

import (
	"testing"

	"github.com/grobinson-grafana/matchers"
	prom_matchers "github.com/prometheus/alertmanager/pkg/labels"
)

const (
	simpleRegex  = "{foo=~\"[a-zA-Z_:][a-zA-Z0-9_:]*\"}"
	complexRegex = "{foo=~\"[a-zA-Z_:][a-zA-Z0-9_:]*\",bar=~\"[a-zA-Z_:]\",baz!~\"[a-zA-Z_:][a-zA-Z0-9_:]*\",qux!~\"[a-zA-Z_:]\"}"
)

func BenchmarkMatchersRegexSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := matchers.Parse(simpleRegex); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMatchersRegexComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := matchers.Parse(complexRegex); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrometheusRegexSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := prom_matchers.ParseMatchers(simpleRegex); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrometheusRegexComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := prom_matchers.ParseMatchers(complexRegex); err != nil {
			b.Fatal(err)
		}
	}
}
