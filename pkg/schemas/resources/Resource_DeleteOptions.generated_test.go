package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/resources"
)

func TestExpandResourceDeleteOptions(t *testing.T) {
	_default := resources.DeleteOptions{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want resources.DeleteOptions
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"interval": nil,
					"wait":     nil,
					"count":    nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceDeleteOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceDeleteOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDeleteOptionsInto(t *testing.T) {
	_default := map[string]interface{}{
		"interval": nil,
		"wait":     nil,
		"count":    nil,
	}
	type args struct {
		in resources.DeleteOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.DeleteOptions{},
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() resources.DeleteOptions {
					subject := resources.DeleteOptions{}
					subject.Interval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Wait - default",
			args: args{
				in: func() resources.DeleteOptions {
					subject := resources.DeleteOptions{}
					subject.Wait = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Count - default",
			args: args{
				in: func() resources.DeleteOptions {
					subject := resources.DeleteOptions{}
					subject.Count = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceDeleteOptionsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDeleteOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDeleteOptions(t *testing.T) {
	_default := map[string]interface{}{
		"interval": nil,
		"wait":     nil,
		"count":    nil,
	}
	type args struct {
		in resources.DeleteOptions
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: resources.DeleteOptions{},
			},
			want: _default,
		},
		{
			name: "Interval - default",
			args: args{
				in: func() resources.DeleteOptions {
					subject := resources.DeleteOptions{}
					subject.Interval = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Wait - default",
			args: args{
				in: func() resources.DeleteOptions {
					subject := resources.DeleteOptions{}
					subject.Wait = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Count - default",
			args: args{
				in: func() resources.DeleteOptions {
					subject := resources.DeleteOptions{}
					subject.Count = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceDeleteOptions(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDeleteOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
