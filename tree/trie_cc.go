package tree

import "sync"

func newTrieCC() Trie {
	return &trieNodeCC{
		mtx:      new(sync.RWMutex),
		children: make(map[rune]*trieNodeCC),
		value:    nil,
		hasWord:  false,
		parent:   nil,
	}
}

type trieNodeCC struct {
	mtx      *sync.RWMutex
	letter   rune
	children map[rune]*trieNodeCC
	value    interface{}
	word     string
	hasWord  bool
	parent   *trieNodeCC
}

func (node *trieNodeCC) Word() string {
	node.mtx.RLock()
	word := node.word
	node.mtx.RUnlock()
	return word
}

func (node *trieNodeCC) Value() interface{} {
	node.mtx.RLock()
	value := node.value
	node.mtx.RUnlock()
	return value
}

func (node *trieNodeCC) add(new bool, word string, letters []rune, value interface{}) bool {
	if len(letters) == 0 {
		node.hasWord = true
		node.word = word
		node.value = value
		return true
	}

	if node.children == nil {
		node.children = make(map[rune]*trieNodeCC)
	}

	letter := letters[0]

	if new {
		newNode := &trieNodeCC{
			mtx:    node.mtx,
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
		child = &trieNodeCC{
			mtx:    node.mtx,
			letter: letter,
			parent: node,
		}
		node.children[letter] = child
	}
	return child.add(new, word, letters[1:], value)
}

func (node *trieNodeCC) iterate(iterate func(node *trieNodeCC)) {
	if node.children != nil {
		for _, child := range node.children {
			iterate(child)
			child.iterate(iterate)
		}
	}
}

func (node *trieNodeCC) contains(letters []rune) bool {
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

func (node *trieNodeCC) containsPrefix(letters []rune) bool {
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

func (node *trieNodeCC) lpm(letters []rune) *trieNodeCC {
	found := (*trieNodeCC)(nil)
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

func (node *trieNodeCC) find(letters []rune) *trieNodeCC {
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

func (node *trieNodeCC) delete(letters []rune) {
	length := len(letters)
	letter := letters[length-1]
	if node.children != nil {
		delete(node.children, letter)
	}
	if !node.hasWord && node.parent != nil {
		node.parent.delete(letters[0 : length-1])
	}
}

func (trie *trieNodeCC) Add(word string, value interface{}) bool {
	trie.mtx.Lock()
	defer trie.mtx.Unlock()

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
		node = &trieNodeCC{
			mtx:    trie.mtx,
			letter: letter,
			parent: trie,
		}
		trie.children[letter] = node
	}
	return node.add(new, word, letters[1:], value)
}

func (trie *trieNodeCC) List() []string {
	trie.mtx.RLock()
	defer trie.mtx.RUnlock()

	words := []string{}
	iterate := func(node *trieNodeCC) {
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

func (trie *trieNodeCC) Clear() {
	trie.mtx.Lock()
	defer trie.mtx.Unlock()

	trie.children = make(map[rune]*trieNodeCC)
}

func (trie *trieNodeCC) Contains(word string) bool {
	trie.mtx.RLock()
	defer trie.mtx.RUnlock()

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

func (trie *trieNodeCC) ContainsPrefix(prefix string) bool {
	trie.mtx.RLock()
	defer trie.mtx.RUnlock()

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
func (trie *trieNodeCC) LPM(longerWord string) (TrieNode, bool) {
	trie.mtx.RLock()
	defer trie.mtx.RUnlock()

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

func (trie *trieNodeCC) Delete(word string) bool {
	trie.mtx.Lock()
	defer trie.mtx.Unlock()

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
