package tree

type trie struct {
	children map[rune]*trieNode
}

func newTrie() Trie {
	return &trie{
		children: make(map[rune]*trieNode),
	}
}

type trieNode struct {
	letter   rune
	children map[rune]*trieNode
	value    interface{}
	word     string
	hasWord  bool
}

func (node *trieNode) add(new bool, word string, letters []rune, value interface{}) bool {
	if len(letters) == 0 {
		node.hasWord = true
		node.word = word
		node.value = value
		return true
	}

	if node.children == nil {
		node.children = make(map[rune]*trieNode)
	}

	letter := letters[0]

	if new {
		newNode := &trieNode{
			letter: letter,
		}
		node.children[letter] = newNode
		return newNode.add(new, word, letters[1:], value)
	}

	child, ok := node.children[letter]
	new = false
	if !ok {
		new = true
		child = &trieNode{
			letter: letter,
		}
		node.children[letter] = child
	}
	return child.add(new, word, letters[1:], value)
}

func (node *trieNode) iterator(iterate func(node *trieNode)) {
	if node.children != nil {
		for _, child := range node.children {
			iterate(child)
			child.iterator(iterate)
		}
	}
}

func (trie *trie) Add(word string, value interface{}) bool {
	letters := []rune(word)
	length := len(letters)
	if length == 0 {
		return false
	}

	letter := letters[0]
	node, ok := trie.children[letter]
	new := false
	if !ok {
		new = true
		node = &trieNode{
			letter: letter,
		}
		trie.children[letter] = node
	}
	return node.add(new, word, letters[1:], value)
}

func (trie *trie) List() []string {
	words := []string{}
	iterate := func(node *trieNode) {
		if node.hasWord {
			words = append(words, node.word)
		}
	}
	for _, node := range trie.children {
		iterate(node)
		node.iterator(iterate)
	}
	return words
}

func (trie *trie) Clear() {
	trie.children = make(map[rune]*trieNode)
}

func (trie *trie) Contains(word string) bool {
	return false
}
