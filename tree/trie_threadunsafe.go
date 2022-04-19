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
	node, ok := node.children[letter]
	new := false
	if !ok {
		node = &trieNode{
			letter: letter,
		}
		node.children[letter] = node
		new = true
	}
	return node.add(new, word, letters[1:], value)
}

func (trie *trie) Add(word string, value interface{}) bool {
	if len(word) == 0 {
		return false
	}
	letters := []rune(word)
	letter := letters[0]
	node, ok := trie.children[letter]
	new := false
	if !ok {
		node = &trieNode{
			letter: letter,
		}
		trie.children[letter] = node
		new = true
	}
	return node.add(new, word, letters[1:], value)
}

func (trie *trie) Clear() {
	trie.children = make(map[rune]*trieNode)
}

func (trie *trie) Contains(word string) bool {
}
