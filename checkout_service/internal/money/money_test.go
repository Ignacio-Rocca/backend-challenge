package money

import "testing"

func TestFormat(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "01 - format success",
			args: args{n: "10000"},
			want: 100.00,
		},
		{
			name: "02 - format success",
			args: args{n: "10050"},
			want: 100.50,
		},
		{
			name: "03 - format success",
			args: args{n: "1005"},
			want: 10.05,
		},
		{
			name: "03 - format invalid string return 0",
			args: args{n: "sadsdasd"},
			want: 0.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFromIntStringToFloat(tt.args.n); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatFromFloatToInt(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01 - success",
			args: args{
				amount: 10.00,
			},
			want: "1000",
		},
		{
			name: "02 - success",
			args: args{
				amount: 0.00,
			},
			want: "0",
		},
		{
			name: "03 - success",
			args: args{
				amount: 100.40,
			},
			want: "10040",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFromFloatToInt(tt.args.amount); got != tt.want {
				t.Errorf("FormatFromFloatToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
