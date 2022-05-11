package tree

func newTrie() Trie {
	return &trieNode{
		children: make(map[rune]*trieNode),
		value:    nil,
		hasWord:  false,
		parent:   nil,
	}
}

type trieNode struct {
	letter   rune
	children map[rune]*trieNode
	value    interface{}
	word     string
	hasWord  bool
	parent   *trieNode
}

func (node *trieNode) Word() string {
	return node.word
}

func (node *trieNode) Value() interface{} {
	return node.value
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
			parent: node,
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
			parent: node,
		}
		node.children[letter] = child
	}
	return child.add(new, word, letters[1:], value)
}

func (node *trieNode) iterate(iterate func(node *trieNode)) {
	if node.children != nil {
		for _, child := range node.children {
			iterate(child)
			child.iterate(iterate)
		}
	}
}

func (node *trieNode) contains(letters []rune) bool {
	if len(letters) == 0 {
		if node.hasWord {
			return true
		}
		return false
	}

	if node.children == nil {
		return false
	}

	letter := letters[0]
	child, ok := node.children[letter]
	if !ok {
		return false
	}
	return child.contains(letters[1:])
}

func (node *trieNode) containsPrefix(letters []rune) bool {
	if len(letters) == 0 {
		return true
	}

	if node.children == nil {
		return false
	}

	letter := letters[0]
	child, ok := node.children[letter]
	if !ok {
		return false
	}
	return child.contains(letters[1:])
}

func (node *trieNode) lpm(letters []rune) *trieNode {
	found := (*trieNode)(nil)
	if node.hasWord {
		found = node
	}
	if len(letters) == 0 || node.children == nil {
		return found
	}

	letter := letters[0]
	child, ok := node.children[letter]
	if !ok {
		return found
	}
	new := child.lpm(letters[1:])
	if new != nil {
		found = new
	}
	return found
}

func (node *trieNode) find(letters []rune) *trieNode {
	if len(letters) == 0 {
		if node.hasWord {
			return node
		}
		return nil
	}

	if node.children == nil {
		return nil
	}

	letter := letters[0]
	child, ok := node.children[letter]
	if !ok {
		return nil
	}
	return child.find(letters[1:])
}

func (node *trieNode) delete(letters []rune) {
	length := len(letters)
	letter := letters[length-1]
	if node.children != nil {
		delete(node.children, letter)
	}
	if !node.hasWord && node.parent != nil {
		node.parent.delete(letters[0 : length-1])
	}
}

func (trie *trieNode) Add(word string, value interface{}) bool {
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
			parent: trie,
		}
		trie.children[letter] = node
	}
	return node.add(new, word, letters[1:], value)
}

func (trie *trieNode) List() []string {
	words := []string{}
	iterate := func(node *trieNode) {
		if node.hasWord {
			words = append(words, node.word)
		}
	}
	for _, node := range trie.children {
		iterate(node)
		node.iterate(iterate)
	}
	return words
}

func (trie *trieNode) Clear() {
	trie.children = make(map[rune]*trieNode)
}

func (trie *trieNode) Contains(word string) bool {
	letters := []rune(word)
	length := len(letters)
	if length == 0 {
		return false
	}
	letter := letters[0]
	node, ok := trie.children[letter]
	if !ok {
		return false
	}
	return node.contains(letters[1:])
}

func (trie *trieNode) ContainsPrefix(prefix string) bool {
	letters := []rune(prefix)
	length := len(letters)
	if length == 0 {
		return false
	}
	letter := letters[0]
	node, ok := trie.children[letter]
	if !ok {
		return false
	}
	return node.containsPrefix(letters[1:])
}

// longest prefix matching
func (trie *trieNode) LPM(longerWord string) (TrieNode, bool) {
	letters := []rune(longerWord)
	length := len(letters)
	if length == 0 {
		return nil, false
	}
	letter := letters[0]
	node, ok := trie.children[letter]
	if !ok {
		return nil, false
	}
	found := node.lpm(letters[1:])
	return found, found != nil
}

func (trie *trieNode) Delete(word string) bool {
	letters := []rune(word)
	length := len(letters)
	if length == 0 {
		return false
	}
	letter := letters[0]
	node, ok := trie.children[letter]
	if !ok {
		return false
	}
	found := node.find(letters[1:])
	if found != nil {
		if found.children == nil {
			found.parent.delete(letters)
		} else {
			found.hasWord = false
		}
		return true
	}
	return false
}
