package fourth_edition

import (
	"reflect"
	"testing"
)

func TestTrieST_keyWithPrefix(t *testing.T) {
	tests := []struct {
		name string
		pre  string
		want []string
		init func(st *TrieST)
	}{
		{
			name: "empty",
			pre:  "",
			want: nil,
			init: func(st *TrieST) {},
		},
		{
			name: "normal",
			pre:  "a",
			want: []string{"abc", "aef"},
			init: func(st *TrieST) {
				st.put("abc", 1)
				st.put("bcd", 1)
				st.put("aef", 1)
				st.put("jkl", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := NewTrieST()
			tt.init(st)
			if got := st.keyWithPrefix(tt.pre); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyWithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrieST_keysThatMatch(t *testing.T) {
	tests := []struct {
		name string
		pat  string
		want []string
		init func(st *TrieST)
	}{
		{
			name: "empty",
			pat:  "",
			want: nil,
			init: func(st *TrieST) {},
		},
		{
			name: "normal",
			pat:  "abc",
			want: []string{"abc"},
			init: func(st *TrieST) {
				st.put("abc", 1)
				st.put("bcd", 1)
				st.put("aef", 1)
				st.put("jkl", 1)
			},
		},
		{
			name: "match",
			pat:  "a.b",
			want: []string{"aab", "acb"},
			init: func(st *TrieST) {
				st.put("acb", 1)
				st.put("bcd", 1)
				st.put("aab", 1)
				st.put("ab", 1)
				st.put("aaab", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := NewTrieST()
			tt.init(st)
			if got := st.keysThatMatch(tt.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyWithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrieST_longestPrefixOf(t *testing.T) {
	tests := []struct {
		name string
		pre  string
		want string
		init func(st *TrieST)
	}{
		{
			name: "empty",
			pre:  "",
			want: "",
			init: func(st *TrieST) {},
		},
		{
			name: "normal",
			pre:  "abcdefg",
			want: "abcd",
			init: func(st *TrieST) {
				st.put("ab", 1)
				st.put("abc", 1)
				st.put("abcd", 1)
				st.put("bace", 1)
				st.put("baceg", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := NewTrieST()
			tt.init(st)
			if got := st.longestPrefixOf(tt.pre); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestPrefixOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
