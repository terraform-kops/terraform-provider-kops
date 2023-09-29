package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCloudConfiguration(t *testing.T) {
	_default := kops.CloudConfiguration{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CloudConfiguration
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"manage_storage_classes": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCloudConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudConfigurationInto(t *testing.T) {
	_default := map[string]interface{}{
		"manage_storage_classes": nil,
	}
	type args struct {
		in kops.CloudConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudConfiguration{},
			},
			want: _default,
		},
		{
			name: "ManageStorageClasses - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ManageStorageClasses = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCloudConfigurationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCloudConfiguration(t *testing.T) {
	_default := map[string]interface{}{
		"manage_storage_classes": nil,
	}
	type args struct {
		in kops.CloudConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CloudConfiguration{},
			},
			want: _default,
		},
		{
			name: "ManageStorageClasses - default",
			args: args{
				in: func() kops.CloudConfiguration {
					subject := kops.CloudConfiguration{}
					subject.ManageStorageClasses = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCloudConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCloudConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
