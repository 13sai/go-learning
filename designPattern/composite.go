package designPattern

// 组织接口，都实现统计人数的功能
type IOrganization interface {
	Count() int
}

// 员工
type Employee struct {
	Name string
}

// 人数统计
func (Employee) Count() int {
	return 1
}

// 部门
type Department struct {
	Name string

	SubOrganizations []IOrganization
}

// 人数统计
func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

// 添加子节点
func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

// NewOrganization 构建组织架构 demo
func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}