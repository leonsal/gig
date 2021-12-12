package gig

type Group struct {
	Widget
}

func NewGroup() *Group {

	g := new(Group)
	g.Init(g)
	g.SetRender(func() {
		g.RenderChildren()
	})
	return g
}

func (g *Group) Save(pg **Group) *Group {

	*pg = g
	return g
}

func (g *Group) AddChildren(c ...IWidget) *Group {

	g.Widget.AddChildren(c...)
	return g
}

type Func struct {
	Widget
}

func NewFunc(f func()) *Func {

	fn := new(Func)
	fn.Init(fn)
	fn.SetRender(f)
	return fn
}
