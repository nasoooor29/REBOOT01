package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	TestDataPath = "data"
)

func checkTestData() error {
	// check the number of files inside dir
	a, err := os.ReadDir(filepath.Join(TestDataPath, "ans"))
	if err != nil {
		return fmt.Errorf("there is no answer folder")
	}
	q, err := os.ReadDir(filepath.Join(TestDataPath, "que"))
	if err != nil {
		return fmt.Errorf("there is no question folder")
	}
	if len(a) != len(q) {
		return fmt.Errorf("Number of questions and answers do not match")
	}
	// match every question with every answer file
	for i := 0; i < len(a); i++ {
		// read the question file
		qn := q[i].Name()
		an := a[i].Name()
		if qn != an {
			return fmt.Errorf("not every question has an answer file")
		}
	}
	return nil
}

func Test_main(t *testing.T) {
	err := checkTestData()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	type args struct {
		in       string
		out      string
		expected string
	}
	tests := []struct {
		name string
		a    args
	}{
		{"unit test 1", args{"data/que/1.txt", "data/results/1.txt", "data/ans/1.txt"}},
		{"unit test 2", args{"data/que/2.txt", "data/results/2.txt", "data/ans/2.txt"}},
		{"unit test 3", args{"data/que/3.txt", "data/results/3.txt", "data/ans/3.txt"}},
		{"unit test 4", args{"data/que/4.txt", "data/results/4.txt", "data/ans/4.txt"}},
		{"unit test 5", args{"data/que/5.txt", "data/results/5.txt", "data/ans/5.txt"}},
		{"unit test 6", args{"data/que/6.txt", "data/results/6.txt", "data/ans/6.txt"}},
		{"unit test 7", args{"data/que/7.txt", "data/results/7.txt", "data/ans/7.txt"}},
		{"unit test 8", args{"data/que/8.txt", "data/results/8.txt", "data/ans/8.txt"}},
		{"unit test 9", args{"data/que/9.txt", "data/results/9.txt", "data/ans/9.txt"}},
		{"unit test 10", args{"data/que/10.txt", "data/results/10.txt", "data/ans/10.txt"}},
		{"unit test 11", args{"data/que/11.txt", "data/results/11.txt", "data/ans/11.txt"}},
		{"unit test 12", args{"data/que/12.txt", "data/results/12.txt", "data/ans/12.txt"}},
		{"unit test 13", args{"data/que/13.txt", "data/results/13.txt", "data/ans/13.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{"cmd", tt.a.in, tt.a.out}
			main()
			data, err := ReadFileContents(tt.a.out)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			expected, err := ReadFileContents(tt.a.expected)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if string(data) != string(expected) {
				t.Errorf("\n\texpected: \t%v\n\tgot: \t\t%v", string(expected), string(data))
			}
		})
	}
}

func TestReloaded(t *testing.T) {

	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{
			"unit test 1",
			"It has been 10 (bin) years",
			"It has been 2 years",
			false,
		},
		{
			"unit test 2",
			"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
			false,
		},
		{
			"unit test 3",
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68.",
			false,
		},
		{
			"unit test 4",
			"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you.",
			false,
		},
		{
			"unit test 5",
			"Punctuation tests are ... kinda boring ,don't you think !?",
			"Punctuation tests are... kinda boring, don't you think!?",
			false,
		},

		{
			"unit test 6",
			"I am exactly how they describe me: ' awesome '",
			"I am exactly how they describe me: 'awesome'",
			false,
		},
		{
			"unit test 7",
			"As Elton John said: ' I am the most well-known homosexual in the world '",
			"As Elton John said: 'I am the most well-known homosexual in the world'",
			false,
		},
		{
			"unit test 8",
			"I am exactly how they describe me: ' awesome ' or whaaaat?",
			"I am exactly how they describe me: 'awesome' or whaaaat?",
			false,
		},
		{
			"unit test 9",
			"it (cap, 5) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
			true,
		},
		{
			"unit test 10",
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
			false,
		},
		{
			"unit test 11",
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
			false,
		},
		{
			"unit test 12",
			"Don not be sad ,because sad backwards is das . And das not good",
			"Don not be sad, because sad backwards is das. And das not good",
			false,
		},
		{
			"unit test 13",
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalReloaded(tt.in)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Reloaded() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("\n\tin:\t %v\n\tgot:\t %v\n\twant:\t %v", tt.in, got, tt.want)
			}
		})
	}
}
