package tools

type Bag struct {
	Items []string
}

func (b *Bag) Add(item string) {
	b.Items = append(b.Items, item)
}

func (b *Bag) List() string {
	if len(b.Items) == 0 {
		return "Mochila esta vazia"
	}
	result := "Mochila: "
	for _, item := range b.Items {
		result += item + ", "
	}
	return result[:len(result)-2]
}