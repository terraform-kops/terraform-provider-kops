package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceAssetsSpec(t *testing.T) {
	_default := kops.AssetsSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AssetsSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"container_registry": nil,
					"file_repository":    nil,
					"container_proxy":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceAssetsSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceAssetsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAssetsSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"container_registry": nil,
		"file_repository":    nil,
		"container_proxy":    nil,
	}
	type args struct {
		in kops.AssetsSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AssetsSpec{},
			},
			want: _default,
		},
		{
			name: "ContainerRegistry - default",
			args: args{
				in: func() kops.AssetsSpec {
					subject := kops.AssetsSpec{}
					subject.ContainerRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileRepository - default",
			args: args{
				in: func() kops.AssetsSpec {
					subject := kops.AssetsSpec{}
					subject.FileRepository = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerProxy - default",
			args: args{
				in: func() kops.AssetsSpec {
					subject := kops.AssetsSpec{}
					subject.ContainerProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceAssetsSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAssetsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceAssetsSpec(t *testing.T) {
	_default := map[string]interface{}{
		"container_registry": nil,
		"file_repository":    nil,
		"container_proxy":    nil,
	}
	type args struct {
		in kops.AssetsSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AssetsSpec{},
			},
			want: _default,
		},
		{
			name: "ContainerRegistry - default",
			args: args{
				in: func() kops.AssetsSpec {
					subject := kops.AssetsSpec{}
					subject.ContainerRegistry = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FileRepository - default",
			args: args{
				in: func() kops.AssetsSpec {
					subject := kops.AssetsSpec{}
					subject.FileRepository = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ContainerProxy - default",
			args: args{
				in: func() kops.AssetsSpec {
					subject := kops.AssetsSpec{}
					subject.ContainerProxy = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceAssetsSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceAssetsSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
