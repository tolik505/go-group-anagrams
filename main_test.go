package main

import (
	"reflect"
	"testing"
)

func Test_anagrams(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "success",
			args: args{
				s: []string{"abb", "tere", "bba", "terere", "abb", "eret", "bba"},
			},
			want: [][]string{
				{"abb", "bba", "abb", "bba"},
				{"tere", "eret"},
				{"terere"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagrams(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s1: "abba",
				s2: "baba",
			},
			want: true,
		},
		{
			name: "not",
			args: args{
				s1: "abbal",
				s2: "baba",
			},
			want: false,
		},
		{
			name: "not",
			args: args{
				s1: "baba",
				s2: "abbal",
			},
			want: false,
		},
		{
			name: "not",
			args: args{
				s1: "aba",
				s2: "abb",
			},
			want: false,
		},
		{
			name: "not2",
			args: args{
				s1: "abab",
				s2: "abb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
