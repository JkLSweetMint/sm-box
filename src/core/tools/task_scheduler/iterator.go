package task_scheduler

// iterator - описание итератора для коллекции по хранению задач планировщика.
type iterator interface {
	Index() int
	Value() *Task
	Has() bool
	Len() (l int)
	Next()
	Prev()
	Reset()
	End()
}

// baseIterator - итератор для коллекции по хранению задач планировщика.
type baseIterator struct {
	shelf    *baseShelf
	index    int
	internal int
}

// Index - получение текущего индекса элемента в коллекции.
func (iter *baseIterator) Index() int {
	return iter.index
}

// Value - получение текущего элемента из коллекции.
func (iter *baseIterator) Value() (t *Task) {
	if iter.index < 0 || len(iter.shelf.Tasks) <= iter.index {
		return
	}

	return iter.shelf.Tasks[iter.index]
}

// Has - имеются ли ещё элементы для дальнейшей итерации.
func (iter *baseIterator) Has() (has bool) {
	return iter.internal >= 0 && iter.internal < len(iter.shelf.Tasks)
}

// Len - получение длинны коллекции.
func (iter *baseIterator) Len() (l int) {
	return len(iter.shelf.Tasks)
}

// Next - переход к следующему элементу в коллекции.
func (iter *baseIterator) Next() {
	iter.internal++

	if iter.Has() {
		iter.index++
	}
}

// Prev - переход к предыдущему элементу в коллекции
func (iter *baseIterator) Prev() {
	iter.internal--

	if iter.Has() {
		iter.index--
	}
}

// Reset - сброс итерации.
func (iter *baseIterator) Reset() {
	iter.index = 0
	iter.internal = 0
}

// End - переходит к последнему элементу коллекции.
func (iter *baseIterator) End() {
	iter.index = len(iter.shelf.Tasks) - 1
	iter.internal = iter.index
}
