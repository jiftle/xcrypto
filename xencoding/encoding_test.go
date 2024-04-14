package xencoding

import "testing"

func TestHexStr2Utf8Str(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "test_01",
			args: args{
				data: Utf8Str2HexStr("张三"),
			},
			wantOut: "张三",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := HexStr2Utf8Str(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexStr2Utf8Str() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("HexStr2Utf8Str() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
