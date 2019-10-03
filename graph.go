package pp_algo

type Graph interface {
    AddNode(data interface{}) (index uint, ok bool)
    AddEdge(fromIndex uint, toIndex uint) bool

    GetDataForIndex(index uint) (interface{}, bool)

    TopologicalSort() []uint
}

type Node interface {

}

func NewOrientedGraph() OrientedGraph {
    return &orientedGraphImpl{
        nodes: map[uint]node{},
        edges: []edge{},
        lastUsedIndex: 0,
    }
}

type OrientedGraph interface {
    Graph
}

type orientedGraphImpl struct {
    nodes map[uint]node
    edges []edge
    lastUsedIndex uint
}

func (g *orientedGraphImpl) GetDataForIndex(index uint) (interface{}, bool) {
    n, ok := g.nodes[index]
    return n.data, ok
}

type NodeColor int

const (
    ColorWhite = iota
    ColorGrey
    ColorBlack
)

type orderedSet struct {
    ls []uint
}

func (s *orderedSet) Append(v uint) {
    s.ls = append(s.ls, v)
}

func (g *orientedGraphImpl) getAdjacencies(fromIndex uint) []uint {
    var res []uint
    for _, edge := range g.edges {
        if edge.from == fromIndex {
            res = append(res, edge.to)
        }
    }
    return res
}

func (g *orientedGraphImpl) iterTopologicalSort(nodes *map[uint]NodeColor, curNode uint, res *orderedSet) {
    c := (*nodes)[curNode]
    if c == ColorBlack {
        return
    } else if c == ColorGrey {
        panic("cycle") // TODO: normal error handling
    } else {
        (*nodes)[curNode] = ColorGrey
        for _, a := range g.getAdjacencies(curNode) {
            g.iterTopologicalSort(nodes, a, res)
        }
        (*nodes)[curNode] = ColorBlack
        res.Append(curNode)
    }
}

func (g *orientedGraphImpl) TopologicalSort() []uint {
    nodes := map[uint]NodeColor{}
    for index, _ := range g.nodes {
        nodes[index] = ColorWhite
    }
    res := orderedSet{ls: []uint{}}
    for ind, _ := range g.nodes {
        g.iterTopologicalSort(&nodes, ind, &res)
    }
    return res.ls
    //for _, v := range res.ls {
    //   fmt.Println(g.nodes[v].data )
    //}
}

func (g *orientedGraphImpl) AddNode(data interface{}) (index uint, ok bool) {
    index = g.lastUsedIndex + 1
    g.nodes[index] = node{index, data}
    g.lastUsedIndex = index
    return index, true
}

func (g *orientedGraphImpl) AddEdge(fromIndex uint, toIndex uint) bool {
    if _, ok := g.nodes[fromIndex]; !ok {
        return false
    }
    if _, ok := g.nodes[toIndex]; !ok {
        return false
    }
    g.edges = append(g.edges, edge{fromIndex, toIndex})
    return true
}

type edge struct {
    from uint
    to uint
}

type node struct {
    index uint
    data interface{}
}

//func main() {
//    g := orientedGraphImpl{
//        nodes: map[uint]node{},
//        edges: []edge{},
//        lastUsedIndex: 0,
//    }
//    a, _ := g.AddNode("a")
//    b, _ := g.AddNode("b")
//    c, _ := g.AddNode("c")
//    d, _ := g.AddNode("d")
//    e, _ := g.AddNode("e")
//    g.AddEdge(a, b)
//    g.AddEdge(a, c)
//    g.AddEdge(a, d)
//    g.AddEdge(a, e)
//    g.AddEdge(b, d)
//    g.AddEdge(c, d)
//    g.AddEdge(c, e)
//    g.AddEdge(d, e)
//    g.AddEdge(b, c)
//
//    reflect.ValueOf(&g).MethodByName("TopologicalSort").
//    reflect.ValueOf(&g).MethodByName("TopologicalSort").Call([]reflect.Value{reflect.ValueOf(123)})
//
//    g.TopologicalSort(456)
//}