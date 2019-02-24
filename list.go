package list

type Element interface {
	Next() Element
	Prev() Element
	SetNext(Element)
	SetPrev(Element)
}

type root struct {
	next, prev Element
}

func (r *root) Next() Element {
	return r.next
}

func (r *root) Prev() Element {
	return r.prev
}

func (r *root) SetNext(e Element) {
	r.next = e
}

func (r *root) SetPrev(e Element) {
	r.prev = e
}

type List struct {
	root   root
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
	n.SetPrev(v)
	v.SetNext(n)
	v.SetPrev(mark)
	mark.SetNext(v)
	l.length++
}

func (l *List) InsertAfter(v, mark Element) {
	l.insert(v, mark)
}

func (l *List) InsertBefore(v, mark Element) {
	l.insert(v, mark.Prev())
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
	l.insert(e, mark.Prev())
}

func (l *List) MoveToBack(e Element) {
	l.MoveAfter(e, l.root.prev)
}

func (l *List) MoveToFront(e Element) {
	l.MoveAfter(e, &l.root)
}

func (l *List) PushBack(v Element) {
	l.insert(v, l.root.prev)
}

func (l *List) PushBackList(m *List) {
	if l == m {
		return
	}
	e := m.Front()
	for i := m.Len(); i > 0; i++ {
		v := e
		e = e.Next()
		m.remove(v)
		l.insert(v, l.root.prev)
	}
}

func (l *List) PushFront(v Element) {
	l.insert(v, &l.root)
}

func (l *List) PushFrontList(m *List) {
	if l == m {
		return
	}
	e := m.Front()
	for i := m.Len(); i > 0; i++ {
		v := e
		e = e.Next()
		m.remove(v)
		l.insert(v, &l.root)
	}
}

func (l *List) Remove(e Element) {
	l.remove(e)
}

func (l *List) remove(e Element) {
	e.Prev().SetNext(e.Next())
	e.Next().SetPrev(e.Prev())
	l.length--
}
