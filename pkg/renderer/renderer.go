package renderer

import (
	"github.com/andredurao/prs/pkg/tree"
)

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

type PrRow struct {
	Row     string
	RefName string
	Value   interface{}
}

type (
	renderer struct {
		root      tree.Tree
		renderMap *[]PrRow
	}

	Renderer interface {
		Render()
		RenderMap() *[]PrRow
		renderText(t tree.Tree, text string, spaces []bool, last bool)
		renderItems(t []tree.Tree, spaces []bool)
	}
)

func New(root tree.Tree) Renderer {
	ary := make([]PrRow, 0)
	return &renderer{
		root:      root,
		renderMap: &ary,
	}
}

func (r *renderer) RenderMap() *[]PrRow {
	return r.renderMap
}

func (r *renderer) Render() {
	root := r.root
	row := PrRow{root.Text(), root.Text(), root.Value()}
	*r.renderMap = append(*r.renderMap, row)
	r.renderItems(root.Items(), []bool{})
}

func (r *renderer) renderItems(t []tree.Tree, spaces []bool) {
	for i, f := range t {
		last := i == len(t)-1
		r.renderText(f, f.Text(), spaces, last)
		if len(f.Items()) > 0 {
			spacesChild := append(spaces, last)
			r.renderItems(f.Items(), spacesChild)
		}
	}
}

func (r *renderer) renderText(t tree.Tree, text string, spaces []bool, last bool) {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continueItem
		}
	}

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	rowText := result + indicator + text
	row := PrRow{rowText, t.Text(), t.Value()}
	*r.renderMap = append(*r.renderMap, row)
}
