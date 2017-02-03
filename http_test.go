package ws

import "testing"

type httpVersionCase struct {
	in    []byte
	major int
	minor int
	ok    bool
}

var httpVersionCases = []httpVersionCase{
	{[]byte("HTTP/1.1"), 1, 1, true},
	{[]byte("HTTP/1.0"), 1, 0, true},
	{[]byte("HTTP/1.2"), 1, 2, true},
	{[]byte("HTTP/42.1092"), 42, 1092, true},
}

func TestParseHttpVersion(t *testing.T) {
	for _, c := range httpVersionCases {
		t.Run(string(c.in), func(t *testing.T) {
			major, minor, ok := httpParseVersion(c.in)
			if major != c.major || minor != c.minor || ok != c.ok {
				t.Errorf(
					"parseHttpVersion([]byte(%q)) = %v, %v, %v; want %v, %v, %v",
					string(c.in), major, minor, ok, c.major, c.minor, c.ok,
				)
			}
		})
	}
}

func BenchmarkParseHttpVersion(b *testing.B) {
	for _, c := range httpVersionCases {
		b.Run(string(c.in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = httpParseVersion(c.in)
			}
		})
	}
}
