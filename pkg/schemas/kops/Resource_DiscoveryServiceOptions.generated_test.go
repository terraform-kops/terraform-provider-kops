package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceDiscoveryServiceOptions(t *testing.T) {
	_default := kops.DiscoveryServiceOptions{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DiscoveryServiceOptions
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"url": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceDiscoveryServiceOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceDiscoveryServiceOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDiscoveryServiceOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"url": "",
	}
	type args struct {
		in kops.DiscoveryServiceOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DiscoveryServiceOptions{},
			},
			want: _default,
		},
		{
			name: "URL - default",
			args: args{
				in: func() kops.DiscoveryServiceOptions {
					subject := kops.DiscoveryServiceOptions{}
					subject.URL = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceDiscoveryServiceOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDiscoveryServiceOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDiscoveryServiceOptions(t *testing.T) {
	_default := map[string]interface{}{
		"url": "",
	}
	type args struct {
		in kops.DiscoveryServiceOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DiscoveryServiceOptions{},
			},
			want: _default,
		},
		{
			name: "URL - default",
			args: args{
				in: func() kops.DiscoveryServiceOptions {
					subject := kops.DiscoveryServiceOptions{}
					subject.URL = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceDiscoveryServiceOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDiscoveryServiceOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
