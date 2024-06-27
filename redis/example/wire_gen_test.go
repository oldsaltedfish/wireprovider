package example

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func Test_wireRedis(t *testing.T) {
	tests := []struct {
		name    string
		want    redis.Cmdable
		wantErr bool
	}{
		{
			name:    "test",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wireRedis()
			if (err != nil) != tt.wantErr {
				t.Errorf("wireRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = got.Set(context.Background(), "testKey", "testValue", 0).Err()
			if err != nil {
				t.Errorf("wireRedis() error = %v", err)
				return
			}
			cmd := got.Get(context.Background(), "testKey")
			if cmd.Err() != nil {
				t.Errorf("wireRedis() error = %v", cmd.Err())
				return
			}
			if cmd.Val() != "testValue" {
				t.Errorf("wireRedis() error = %v", cmd.Val())
				return
			}
			err = got.Del(context.Background(), "testKey").Err()
			if err != nil {
				t.Errorf("wireRedis() error = %v", err)
				return
			}
		})
	}
}
