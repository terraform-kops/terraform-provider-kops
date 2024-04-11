package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKarpenterConfig(t *testing.T) {
	_default := kops.KarpenterConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KarpenterConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":        false,
					"log_encoding":   "",
					"log_level":      "",
					"image":          "",
					"memory_limit":   nil,
					"memory_request": nil,
					"cpu_request":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKarpenterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKarpenterConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":        false,
		"log_encoding":   "",
		"log_level":      "",
		"image":          "",
		"memory_limit":   nil,
		"memory_request": nil,
		"cpu_request":    nil,
	}
	type args struct {
		in kops.KarpenterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KarpenterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogEncoding - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.LogEncoding = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.LogLevel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKarpenterConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKarpenterConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":        false,
		"log_encoding":   "",
		"log_level":      "",
		"image":          "",
		"memory_limit":   nil,
		"memory_request": nil,
		"cpu_request":    nil,
	}
	type args struct {
		in kops.KarpenterConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KarpenterConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.Enabled = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogEncoding - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.LogEncoding = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.LogLevel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KarpenterConfig {
					subject := kops.KarpenterConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKarpenterConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKarpenterConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
