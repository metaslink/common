package webix

type Tree struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Open bool `json:"open,omitempty"`
	Attrs map[string]interface{} `json:"attrs,omitempty"`
	Data []*Tree `json:"data,omitempty"`
}

func NewTree(id string, value string, open bool) *Tree {
	data := make([]*Tree, 0)
	return &Tree{Id: id, Value: value, Open: open, Data: data}
}

func (t *Tree) Append(tree *Tree)  {
	t.Data = append(t.Data, tree)
}



