package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceNodeSelectorTerm(t *testing.T) {
	_default := core.NodeSelectorTerm{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.NodeSelectorTerm
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"match_expressions": func() []interface{} { return nil }(),
					"match_fields":      func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceNodeSelectorTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceNodeSelectorTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeSelectorTermInto(t *testing.T) {
	_default := map[string]interface{}{
		"match_expressions": func() []interface{} { return nil }(),
		"match_fields":      func() []interface{} { return nil }(),
	}
	type args struct {
		in core.NodeSelectorTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.NodeSelectorTerm{},
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() core.NodeSelectorTerm {
					subject := core.NodeSelectorTerm{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchFields - default",
			args: args{
				in: func() core.NodeSelectorTerm {
					subject := core.NodeSelectorTerm{}
					subject.MatchFields = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceNodeSelectorTermInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeSelectorTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceNodeSelectorTerm(t *testing.T) {
	_default := map[string]interface{}{
		"match_expressions": func() []interface{} { return nil }(),
		"match_fields":      func() []interface{} { return nil }(),
	}
	type args struct {
		in core.NodeSelectorTerm
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.NodeSelectorTerm{},
			},
			want: _default,
		},
		{
			name: "MatchExpressions - default",
			args: args{
				in: func() core.NodeSelectorTerm {
					subject := core.NodeSelectorTerm{}
					subject.MatchExpressions = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MatchFields - default",
			args: args{
				in: func() core.NodeSelectorTerm {
					subject := core.NodeSelectorTerm{}
					subject.MatchFields = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceNodeSelectorTerm(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceNodeSelectorTerm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
