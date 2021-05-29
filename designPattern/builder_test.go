package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourcePoolConfigBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		builder *designPattern.ResourcePoolConfigBuilder
		want    *designPattern.ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &designPattern.ResourcePoolConfigBuilder{
				Name:     "",
				MaxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &designPattern.ResourcePoolConfigBuilder{
				Name:     "test",
				MaxTotal: 0,
				MaxIdle:  10,
				MinIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &designPattern.ResourcePoolConfigBuilder{
				Name: "test",
			},
			want: &designPattern.ResourcePoolConfig{
				Name:     "test",
				MaxTotal: designPattern.DefaultMaxTotal,
				MaxIdle:  designPattern.DefaultMaxIdle,
				MinIdle:  designPattern.DefaultMinIdle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "Build() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}