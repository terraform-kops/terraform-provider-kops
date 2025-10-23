package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceFileKeySelector(t *testing.T) {
	_default := core.FileKeySelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.FileKeySelector
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"volume_name": "",
					"path":        "",
					"key":         "",
					"optional":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceFileKeySelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceFileKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceFileKeySelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"volume_name": "",
		"path":        "",
		"key":         "",
		"optional":    nil,
	}
	type args struct {
		in core.FileKeySelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.FileKeySelector{},
			},
			want: _default,
		},
		{
			name: "VolumeName - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.VolumeName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Path - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.Path = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optional - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.Optional = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceFileKeySelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceFileKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceFileKeySelector(t *testing.T) {
	_default := map[string]interface{}{
		"volume_name": "",
		"path":        "",
		"key":         "",
		"optional":    nil,
	}
	type args struct {
		in core.FileKeySelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.FileKeySelector{},
			},
			want: _default,
		},
		{
			name: "VolumeName - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.VolumeName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Path - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.Path = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optional - default",
			args: args{
				in: func() core.FileKeySelector {
					subject := core.FileKeySelector{}
					subject.Optional = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceFileKeySelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceFileKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
