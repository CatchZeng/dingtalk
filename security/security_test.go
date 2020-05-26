package security

import "testing"

func TestGetDingTalkURL(t *testing.T) {
	timestamp := "1582163555000"

	type args struct {
		accessToken string
		secret      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "without sign",
			args: args{
				accessToken: "1c53e149ba5de6597ca2442f0e901fd86156780b8ac141e4a75afdc44c85ca4f",
			},
			want:    "https://oapi.dingtalk.com/robot/send?access_token=1c53e149ba5de6597ca2442f0e901fd86156780b8ac141e4a75afdc44c85ca4f",
			wantErr: false,
		},
		{
			name: "with sign",
			args: args{
				accessToken: "1c53e149ba5de6597ca2442f0e901fd86156780b8ac141e4a75afdc44c85ca4f",
				secret:      "SECb90923e19e58b466481e9e7b7a5b4f108a4531abde590ad3967fb29f0eae5c68",
			},
			want:    "https://oapi.dingtalk.com/robot/send?access_token=1c53e149ba5de6597ca2442f0e901fd86156780b8ac141e4a75afdc44c85ca4f&sign=BQKsG%2BQOCl%2BbYJOLc6pxDHxjVquzlZPWgvRzeN2J5zY%3D&timestamp=1582163555000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDingTalkURLWithTimestamp(timestamp, tt.args.accessToken, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDingTalkURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDingTalkURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
