// Based on https://github.com/DiSiqueira/GoTree
package tree

type (
	tree struct {
		text  string
		items []Tree
		value interface{}
	}

	Tree interface {
		Add(text string, value interface{}) Tree
		AddTree(tree Tree)
		Items() []Tree
		Text() string
		Value() interface{}
	}
)

func New(text string, value interface{}) Tree {
	return &tree{
		text:  text,
		items: []Tree{},
		value: value,
	}
}

func (t *tree) Add(text string, value interface{}) Tree {
	n := New(text, value)
	t.items = append(t.items, n)
	return n
}

func (t *tree) AddTree(tree Tree) {
	t.items = append(t.items, tree)
}

func (t *tree) Text() string {
	return t.text
}

func (t *tree) Value() interface{} {
	return t.value
}

func (t *tree) Items() []Tree {
	return t.items
}
