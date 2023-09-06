package split

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestOptionN(t *testing.T){
	cases := []struct {
		name string
		in_1 int
		in_2 []string
		out int
	}{
		{
			name: "ケース１：正常系",
			in_1: 2,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv"},
			out: 0,
		},
		{
			name: "ケース２：正常系",
			in_1: 2,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv","aaa"},
			out: 0,
		},
		{
			name: "ケース３：異常系",
			in_1: 2,
			in_2: []string{},
			out: 1,
		},
		{
			name: "ケース4：異常系",
			in_1: 2,
			in_2: []string{"testdata.csv","aaa"},
			out: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name,func(t *testing.T) {
			want := OptionN(tt.in_1,tt.in_2)
			assert.Equal(t,tt.out,want)
			os.Remove("output_1.txt")
			os.Remove("output_2.txt")
			os.Remove("aaa_output_1.txt")
			os.Remove("aaa_output_2.txt")
		})
	}
}

func TestOptionL(t *testing.T){
	cases := []struct {
		name string
		in_1 int
		in_2 []string
		out int
	}{
		{
			name: "ケース１：正常系",
			in_1: 250,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv"},
			out: 0,
		},
		{
			name: "ケース２：正常系",
			in_1: 250,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv","aaa"},
			out: 0,
		},
		{
			name: "ケース３：異常系",
			in_1: 2,
			in_2: []string{},
			out: 1,
		},
		{
			name: "ケース4：異常系",
			in_1: 2,
			in_2: []string{"testdata.csv","aaa"},
			out: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name,func(t *testing.T) {
			want := OptionL(tt.in_1,tt.in_2)
			assert.Equal(t,tt.out,want)
			os.Remove("output_1.txt")
			os.Remove("output_2.txt")
			os.Remove("aaa_output_1.txt")
			os.Remove("aaa_output_2.txt")
		})
	}
}

func TestOptionB(t *testing.T){
	cases := []struct {
		name string
		in_1 int
		in_2 []string
		out int
	}{
		{
			name: "ケース１：正常系",
			in_1: 9000,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv"},
			out: 0,
		},
		{
			name: "ケース２：正常系",
			in_1: 9000,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv","aaa"},
			out: 0,
		},
		{
			name: "ケース３：異常系",
			in_1: 2,
			in_2: []string{},
			out: 1,
		},
		{
			name: "ケース4：異常系",
			in_1: 2,
			in_2: []string{"testdata.csv","aaa"},
			out: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name,func(t *testing.T) {
			want := OptionB(tt.in_1,tt.in_2)
			assert.Equal(t,tt.out,want)
			os.Remove("output_1.txt")
			os.Remove("output_2.txt")
			os.Remove("output_3.txt")
			os.Remove("aaa_output_1.txt")
			os.Remove("aaa_output_2.txt")
			os.Remove("aaa_output_3.txt")
		})
	}
}

func TestOptionC(t *testing.T){
	cases := []struct {
		name string
		in_1 int
		in_2 []string
		out int
	}{
		{
			name: "ケース１：正常系",
			in_1: 2,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv"},
			out: 0,
		},
		{
			name: "ケース２：正常系",
			in_1: 2,
			in_2: []string{"/Users/taichi/src/EnablementBootcamp_Test/testdata.csv","aaa"},
			out: 0,
		},
		{
			name: "ケース３：異常系",
			in_1: 2,
			in_2: []string{},
			out: 1,
		},
		{
			name: "ケース4：異常系",
			in_1: 2,
			in_2: []string{"testdata.csv","aaa"},
			out: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name,func(t *testing.T) {
			want := OptionC(tt.in_1,tt.in_2)
			assert.Equal(t,tt.out,want)
			os.Remove("output_1.csv")
			os.Remove("output_2.csv")
			os.Remove("aaa_output_1.csv")
			os.Remove("aaa_output_2.csv")
		})
	}
}