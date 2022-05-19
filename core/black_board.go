package core

type TreeData struct {
	NodeMemory     *Memory
	OpenNodes      []IBaseNode
	TraversalDepth int
	TraversalCycle int
}

func NewTreeData() *TreeData {
	return &TreeData{
		NodeMemory:     NewMemory(),
		OpenNodes:      make([]IBaseNode, 0),
		TraversalDepth: 0,
		TraversalCycle: 0,
	}
}

type Memory struct {
	memory map[string]interface{}
}

func NewMemory() *Memory {
	return &Memory{make(map[string]interface{}, 0)}
}

func (m *Memory) Get(key string) interface{} {
	if value, ok := m.memory[key]; ok {
		return value
	}
	return nil
}

func (m *Memory) Set(key string, value interface{}) {
	m.memory[key] = value
}

type TreeMemory struct {
	memory     *Memory
	treeData   *TreeData
	nodeMemory map[string]*Memory
}

func NewTreeMemory() *TreeMemory {
	return &TreeMemory{
		memory:     NewMemory(),
		treeData:   NewTreeData(),
		nodeMemory: make(map[string]*Memory),
	}
}

type BlackBoard struct {
	baseMemory *Memory
	treeMemory map[string]*TreeMemory
}

func NewBlackBoard() *BlackBoard {
	b := &BlackBoard{}
	b.Initialize()
	return b
}

func (b *BlackBoard) Initialize() {
	b.baseMemory = NewMemory()
	b.treeMemory = make(map[string]*TreeMemory)
}

func (b *BlackBoard) getTreeMemory(tScope string) *TreeMemory {
	if _, ok := b.treeMemory[tScope]; !ok {
		b.treeMemory[tScope] = NewTreeMemory()
	}
	return b.treeMemory[tScope]
}

func (b *BlackBoard) getNodeMemory(treeMemory *TreeMemory, nScope string) *Memory {
	memory := treeMemory.nodeMemory
	if _, ok := memory[nScope]; !ok {
		memory[nScope] = NewMemory()
	}
	return memory[nScope]
}

func (b *BlackBoard) getMemory(tScope, nScope string) *Memory {
	memory := b.baseMemory
	if len(tScope) > 0 {
		treeMem := b.getTreeMemory(tScope)
		memory = treeMem.memory
		if len(nScope) > 0 {
			memory = b.getNodeMemory(treeMem, nScope)
		}
	}
	return memory
}

func (b *BlackBoard) getTreeData(tScope string) *TreeData {
	treeMem := b.getTreeMemory(tScope)
	return treeMem.treeData
}

func (b *BlackBoard) SetTree(key string, value interface{}, tScope string) {
	memory := b.getMemory(tScope, "")
	memory.Set(key, value)
}

func (b *BlackBoard) Set(key string, value interface{}, tScope, nScope string) {
	memory := b.getMemory(tScope, nScope)
	memory.Set(key, value)
}

func (b *BlackBoard) Get(key, tScope, nScope string) interface{} {
	memory := b.getMemory(tScope, nScope)
	return memory.Get(key)
}

func (b *BlackBoard) GetBool(key, tScope, nScope string) bool {
	if r := b.Get(key, tScope, nScope); r != nil {
		return r.(bool)
	}
	return false
}

func (b *BlackBoard) GetInt(key, tScope, nScope string) int {
	if r := b.Get(key, tScope, nScope); r != nil {
		return r.(int)
	}
	return 0
}

func (b *BlackBoard) GetInt64(key, tScope, nScope string) int64 {
	if r := b.Get(key, tScope, nScope); r != nil {
		return r.(int64)
	}
	return 0
}

func (b *BlackBoard) GetInt32(key, tScope, nScope string) int32 {
	if r := b.Get(key, tScope, nScope); r != nil {
		return r.(int32)
	}
	return 0
}
