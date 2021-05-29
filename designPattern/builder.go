package designPattern

import "fmt"

const (
	DefaultMaxTotal = 10
	DefaultMaxIdle  = 9
	DefaultMinIdle  = 1
)

// ResourcePoolConfig resource pool
type ResourcePoolConfig struct {
	Name     string
	MaxTotal int
	MaxIdle  int
	MinIdle  int
}

// 用于构建 ResourcePoolConfig
type ResourcePoolConfigBuilder struct {
	Name     string
	MaxTotal int
	MaxIdle  int
	MinIdle  int
}

// SetName SetName
func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	b.Name = name
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", minIdle)
	}
	b.MinIdle = minIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", maxIdle)
	}
	b.MaxIdle = maxIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max tatal cannot <= 0, input: %d", maxTotal)
	}
	b.MaxTotal = maxTotal
	return nil
}

func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.Name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	// 设置默认值
	if b.MinIdle == 0 {
		b.MinIdle = DefaultMinIdle
	}

	if b.MaxIdle == 0 {
		b.MaxIdle = DefaultMaxIdle
	}

	if b.MaxTotal == 0 {
		b.MaxTotal = DefaultMaxTotal
	}

	if b.MaxTotal < b.MaxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.MaxTotal, b.MaxIdle)
	}

	if b.MinIdle > b.MaxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.MaxIdle, b.MinIdle)
	}

	return &ResourcePoolConfig{
		Name:     b.Name,
		MaxTotal: b.MaxTotal,
		MaxIdle:  b.MaxIdle,
		MinIdle:  b.MinIdle,
	}, nil
}