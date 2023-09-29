package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceInstanceRootVolumeSpec(t *testing.T) {
	_default := kops.InstanceRootVolumeSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.InstanceRootVolumeSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"size":           nil,
					"type":           nil,
					"iops":           nil,
					"throughput":     nil,
					"optimization":   nil,
					"encryption":     nil,
					"encryption_key": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceInstanceRootVolumeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceInstanceRootVolumeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceRootVolumeSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"size":           nil,
		"type":           nil,
		"iops":           nil,
		"throughput":     nil,
		"optimization":   nil,
		"encryption":     nil,
		"encryption_key": nil,
	}
	type args struct {
		in kops.InstanceRootVolumeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.InstanceRootVolumeSpec{},
			},
			want: _default,
		},
		{
			name: "Size - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Size = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Type = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IOPS - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.IOPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Throughput - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optimization - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Optimization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Encryption - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Encryption = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionKey - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.EncryptionKey = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceInstanceRootVolumeSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceRootVolumeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceInstanceRootVolumeSpec(t *testing.T) {
	_default := map[string]interface{}{
		"size":           nil,
		"type":           nil,
		"iops":           nil,
		"throughput":     nil,
		"optimization":   nil,
		"encryption":     nil,
		"encryption_key": nil,
	}
	type args struct {
		in kops.InstanceRootVolumeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.InstanceRootVolumeSpec{},
			},
			want: _default,
		},
		{
			name: "Size - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Size = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Type - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Type = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IOPS - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.IOPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Throughput - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Throughput = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optimization - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Optimization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Encryption - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.Encryption = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionKey - default",
			args: args{
				in: func() kops.InstanceRootVolumeSpec {
					subject := kops.InstanceRootVolumeSpec{}
					subject.EncryptionKey = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceInstanceRootVolumeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceInstanceRootVolumeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
