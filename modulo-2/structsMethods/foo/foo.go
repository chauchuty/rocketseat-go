package foo

type Foo struct {
	name string
}

func NewFoo(name string) Foo {
	return Foo{name: name}
}

func (f Foo) GetName() string {
	return f.name
}
