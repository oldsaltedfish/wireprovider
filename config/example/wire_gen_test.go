package example

import (
	"flag"
	"reflect"
	"testing"
)

func Test_wireContainer(t *testing.T) {

	tests := []struct {
		name string
		want string
		f    func()
	}{
		{
			name: "default",
			want: "default",
		},
		{
			name: "custom",
			want: "custom",
			f: func() {
				flag.Set("conf", "config-custom.yml")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.f != nil {
				tt.f()
			}
			if got := wireConfig(); !reflect.DeepEqual(got.Get("type"), tt.want) {
				t.Errorf("wireContainer() = %v, want %v", got.Get("type"), tt.want)
			}
		})
	}
}
