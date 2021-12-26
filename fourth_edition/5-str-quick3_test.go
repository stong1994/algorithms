package fourth_edition

import (
	"reflect"
	"testing"
)

func TestQuick3_Sort(t *testing.T) {
	tests := []struct {
		name  string
		list  []string
		check func(got []string) bool
	}{
		{
			name: "empty",
			list: nil,
			check: func(got []string) bool {
				return got == nil
			},
		},
		{
			name: "normal",
			list: []string{"edu.pri.cs", "com.apple", "edu.pri.cs", "com.cnn", "com.google", "edu.vua.cs",
				"edu.vua.cs", "edu.vua.cs", "com.adobe", "edu.princeton.ee"},
			check: func(got []string) bool {
				want := []string{"com.adobe", "com.apple", "com.cnn", "com.google", "edu.pri.cs", "edu.pri.cs",
					"edu.princeton.ee", "edu.vua.cs", "edu.vua.cs", "edu.vua.cs"}
				return reflect.DeepEqual(got, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Quick3{}
			q.Sort(tt.list)
			if !tt.check(tt.list) {
				t.Errorf("failed, name: %s, list: %v", tt.name, tt.list)
			}
		})
	}
}
