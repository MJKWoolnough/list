package list

type Element interface {
	Next() Element
	Prev() Element
	SetNext(Element)
	SetPrev(Element)
}

type NoDouble struct{}

func (NoDouble) Next() Element   { return nil }
func (NoDouble) Prev() Element   { return nil }
func (NoDouble) SetNext(Element) {}
func (NoDouble) SetPrev(Element) {}

type Simple struct {
	next, prev Element
}

func (s *Simple) Next() Element {
	return s.next
}

func (s *Simple) Prev() Element {
	return s.prev
}

func (s *Simple) SetNext(e Element) {
	s.next = e
}

func (s *Simple) SetPrev(e Element) {
	s.prev = e
}

type List struct {
	root   Simple
	length int
}

func New() *List {
	return &List{}
}

func (l *List) Back() Element {
	return l.root.prev
}

func (l *List) Front() Element {
	return l.root.next
}

func (l *List) insert(v, mark Element) {
	n := mark.Next()
	if n != nil {
		n.SetPrev(v)
	} else {
		l.root.prev = v
	}
	v.SetNext(n)
	if mark != &l.root {
		v.SetPrev(mark)
	}
	mark.SetNext(v)
	l.length++
}

func (l *List) InsertAfter(v, mark Element) {
	l.insert(v, mark)
}

func (l *List) InsertBefore(v, mark Element) {
	l.insert(v, l.prev(mark))
}

func (l *List) Len() int {
	return l.length
}

func (l *List) MoveAfter(e, mark Element) {
	l.remove(e)
	l.insert(e, mark)
}

func (l *List) MoveBefore(e, mark Element) {
	l.remove(e)
	l.insert(e, l.prev(mark))
}

func (l *List) MoveToBack(e Element) {
	l.MoveAfter(e, l.prev(&l.root))
}

func (l *List) MoveToFront(e Element) {
	l.MoveAfter(e, &l.root)
}

func (l *List) prev(e Element) Element {
	p := e.Prev()
	if p != nil {
		return p
	}
	var c Element = &l.root
	for i := 0; i < l.length; i++ {
		n := c.Next()
		if n == e {
			break
		}
		c = n
	}
	return c
}

func (l *List) PushBack(v Element) {
	l.insert(v, l.prev(&l.root))
}

func (l *List) PushBackList(m *List) {
	if l == m {
		return
	}
	for i := m.Len(); i > 0; i++ {
		v := m.Front()
		m.remove(v)
		l.insert(v, l.prev(&l.root))
	}
}

func (l *List) PushFront(v Element) {
	l.insert(v, &l.root)
}

func (l *List) PushFrontList(m *List) {
	if l == m {
		return
	}
	for i := 0; i < m.length; i++ {
		v := m.Back()
		m.remove(v)
		l.insert(v, &l.root)
	}
}

func (l *List) Remove(e Element) {
	l.remove(e)
}

func (l *List) remove(e Element) {
	p := l.prev(e)
	n := e.Next()
	if p != nil {
		p.SetNext(n)
	}
	if n != nil {
		n.SetPrev(p)
	}
	e.SetNext(nil)
	e.SetPrev(nil)
	l.length--
}
