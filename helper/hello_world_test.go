package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Udin",
			request: "Udin",
		},
		{
			name:    "Egin",
			request: "Egin",
		},
		{
			name:    "Le Mineral",
			request: "Le Mineral",
		},
		{
			name:    "EnakMakanAbonSapi",
			request: "EnakMakanAbonSapi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Egin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Egin")
		}
	})
	b.Run("Ucok", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ucok")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Egin")
	}
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Egin",
			request:  "Egin",
			expected: "Hello Egin",
		},
		{
			name:     "Udin",
			request:  "Udin",
			expected: "Hello Udin",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Egin", func(t *testing.T) {
		result := HelloWorld("Egin")
		require.Equal(t, "Hello Egin", result, "Result must be 'Hello Egin'")
	})

	t.Run("Ucok", func(t *testing.T) {
		result := HelloWorld("Ucok")
		require.Equal(t, "Hello Ucok", result, "Result must be 'Hello Ucok'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows!")
	}

	result := HelloWorld("Egin")
	require.Equal(t, "Hello Egin", result, "Result must be 'Hello Egin'")
}

//Disarankan
func TestHelloWorldWithAssert(t *testing.T) {
	result := HelloWorld("Egin")
	assert.Equal(t, "Hello Egin", result, "Result must be 'Hello Egin'")
	fmt.Println("TestHelloWorldWithAssert Done")
}

//Disarankan
func TestHelloWorldWithRequire(t *testing.T) {
	result := HelloWorld("Egin")
	require.Equal(t, "Hello Egin", result, "Result must be 'Hello Egin'")
	fmt.Println("TestHelloWorldWithRequire Done")
}

//Tidak Disarankan
func TestHelloWorldEgin(t *testing.T) {
	result := HelloWorld("Egin")
	if result != "Hello Egin" {
		// unit test failed
		//t.Fail()
		t.Error("Result must be 'Hello Egin'")
	}

	fmt.Println("TestHelloWorldEgin Done")
}

//Tidak Disarankan
func TestHelloWorldGin(t *testing.T) {
	result := HelloWorld("Gin")
	if result != "Hello Gin" {
		// unit test failed
		//t.FailNow()
		t.Fatal("Result must be 'Hello Gin'")
	}

	fmt.Println("TestHelloWorldGin Done")
}
