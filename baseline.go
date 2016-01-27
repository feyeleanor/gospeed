package gospeed

var x int
var u uint
var x32 int32
var u32 uint32
var x64 int64
var u64 uint64
var in interface{} = 0
var h map[int]int = map[int]int{10000000: 0, 10000001: 1}
var hs map[string]int = map[string]int{"10000000": 0, "10000001": 1}
var f func() = func() {}
var farg func(x interface{}) = func(x interface{}) {}
var fvarargs func(x ...interface{}) = func(x ...interface{}) {}
var fint func(x int) = func(x int) {}
var fvarints func(x ...int) = func(x ...int) {}

type dummyInterface1 interface {
	m1()
}

type dummyInterface2 interface {
	dummyInterface1
	m1arg(x interface{})
}

type dummyStructure struct {
	i  int
	in interface{}
	s  []int
	h  map[int]int
}

func (d dummyStructure) m1()                         {}
func (d dummyStructure) m1arg(x interface{})         {}
func (d dummyStructure) m1int(x int)                 {}
func (d dummyStructure) m1varargs(x ...interface{})  {}
func (d dummyStructure) m1varints(x ...int)          {}
func (d *dummyStructure) m2()                        {}
func (d *dummyStructure) m2arg(x interface{})        {}
func (d *dummyStructure) m2int(x int)                {}
func (d *dummyStructure) m2varargs(x ...interface{}) {}
func (d *dummyStructure) m2varints(x ...int)         {}
func (d dummyStructure) get(i int) (r int)           { return d.s[i] }
func (d dummyStructure) set(i, x int)                { d.s[i] = x }

var dummy dummyStructure = dummyStructure{}
var di interface{} = dummyStructure{}

type dummyInterface3 interface {
	get(int) int
	set(int, int)
}

var dummyAccessor dummyStructure = dummyStructure{s: []int{0}}
var dai dummyInterface3 = dummyStructure{s: []int{0}}
