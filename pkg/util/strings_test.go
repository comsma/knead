package util

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Snake Case",
			args: args{s: "order_line"},
			want: "OrderLine",
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPascalCase(tt.args.s); got != tt.want {
				t.Errorf("ToPascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
