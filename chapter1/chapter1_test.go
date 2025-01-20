package chapter1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"testing"
)

func Test_1_1_1(t *testing.T) {
	filePath := "testdata/10.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	s := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("error reading file %s: %v", filePath, err)
	}

	slices.Reverse(s)

	got := s
	want := []string{"9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func f_1_1_2(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)
	s := make([]string, 50)
	i := 0
	for scanner.Scan() {
		s[i%50] = scanner.Text()
		if (i+1)%50 == 0 {
			slices.Reverse(s)
			w.Write([]byte(strings.Join(s, "\n") + "\n"))
			s = make([]string, 50)
		}
		i++
	}
}

func Test_1_1_2(t *testing.T) {
	filePath := "testdata/100.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	f_1_1_2(&buf, file)

	got := buf.String()
	want := `49
48
47
46
45
44
43
42
41
40
39
38
37
36
35
34
33
32
31
30
29
28
27
26
25
24
23
22
21
20
19
18
17
16
15
14
13
12
11
10
9
8
7
6
5
4
3
2
1
0
99
98
97
96
95
94
93
92
91
90
89
88
87
86
85
84
83
82
81
80
79
78
77
76
75
74
73
72
71
70
69
68
67
66
65
64
63
62
61
60
59
58
57
56
55
54
53
52
51
50
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func f_1_1_3(w io.Writer, r io.Reader, n int) {
	rg := NewRingBuffer(n)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if rg.Full() && scanner.Text() == "" {
			w.Write([]byte(rg.First() + "\n"))
		}
		rg.Push(scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func Test_1_1_3(t *testing.T) {
	filePath := "testdata/for_1_1_3.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	f_1_1_3(&buf, file, 42)
	got := buf.String()
	want := `20
21
22
54
55
56
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func solve_1_1_4(w io.Writer, r io.Reader) {
	m := make(map[string]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := m[line]; ok {
			continue
		}
		m[line] = struct{}{}
		w.Write([]byte(line + "\n"))
	}
}

func Test_1_1_4(t *testing.T) {
	filePath := "testdata/for_1_1_4.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	solve_1_1_4(&buf, file)

	got := buf.String()
	want := `0
1
2
3
4
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func solve_1_1_5(w io.Writer, r io.Reader) {
	m := make(map[string]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := m[line]; ok {
			w.Write([]byte(line + "\n"))
			continue
		}
		m[line] = struct{}{}
	}
}

func Test_1_1_5(t *testing.T) {
	filePath := "testdata/for_1_1_4.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	solve_1_1_5(&buf, file)

	got := buf.String()
	want := `1
2
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func solve_1_1_6(w io.Writer, r io.Reader) {
	s := make([]string, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	slices.Sort(s)
	for v := range slices.Values(s) {
		w.Write([]byte(v + "\n"))
	}
}

func Test_1_1_6(t *testing.T) {
	filePath := "testdata/random_10.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	solve_1_1_6(&buf, file)

	got := buf.String()
	want := `0
0
1
2
3
4
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func solve_1_1_7(w io.Writer, r io.Reader) {
	s := make([]string, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	m := make(map[string]int, len(s))
	for v := range slices.Values(s) {
		m[v]++
	}

	slices.Sort(s)
	s = slices.Compact(s)

	newS := make([]string, 0, len(m))
	for v := range slices.Values(s) {
		if cnt := m[v]; cnt == 1 {
			newS = append(newS, v)
		} else {
			newS = append(newS, fmt.Sprintf("%d", cnt))
		}
	}

	for v := range slices.Values(newS) {
		w.Write([]byte(v + "\n"))
	}
}

func Test_1_1_7(t *testing.T) {
	filePath := "testdata/random_10.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	solve_1_1_7(&buf, file)

	got := buf.String()
	want := `2
1
2
3
4
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func solve_1_1_8(w io.Writer, r io.Reader) {
	evens := make([]string, 0)
	odds := make([]string, 0)

	scanner := bufio.NewScanner(r)

	line := 0
	for scanner.Scan() {
		if line%2 == 0 {
			evens = append(evens, scanner.Text())
		} else {
			odds = append(odds, scanner.Text())
		}
		line++
	}

	for v := range slices.Values(evens) {
		w.Write([]byte(v + "\n"))
	}
	for v := range slices.Values(odds) {
		w.Write([]byte(v + "\n"))
	}
}

func Test_1_1_8(t *testing.T) {
	filePath := "testdata/10.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var buf strings.Builder

	solve_1_1_8(&buf, file)

	got := buf.String()
	want := `0
2
4
6
8
1
3
5
7
9
`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
