package graph

import (
    "github.com/pkg/errors"
    "github.com/wlad031/pp-algo/list"
    "strconv"
)

// TODO: custom types as indexes

//region public

type Graph interface {
    AddNode(data interface{}) (index int, e error)
    AddEdge(fromIndex int, toIndex int) error

    GetDataForIndex(index int) (interface{}, error)

    TopologicalSort() ([]int, error)
}

type OrientedGraph interface {
    Graph
}

func NewOrientedGraph() OrientedGraph {
    return &orientedGraphImpl{
        nodes:            map[int]node{},
        edges:            []edge{},
        lastNotUsedIndex: 0,
    }
}

type NodeColor int

const (
    ColorWhite NodeColor = iota
    ColorGrey
    ColorBlack
)

//endregion

//region private

type edge struct {
    from int
    to int
}

type node struct {
    index int
    data interface{}
}

type orientedGraphImpl struct {
    nodes            map[int]node
    edges            []edge
    lastNotUsedIndex int
}

func (g *orientedGraphImpl) AddNode(data interface{}) (index int, e error) {
    index = g.lastNotUsedIndex
    g.nodes[index] = node{index, data}
    g.lastNotUsedIndex = index + 1
    return index, nil
}

func (g *orientedGraphImpl) AddEdge(fromIndex int, toIndex int) error {
    if _, ok := g.nodes[fromIndex]; !ok {
        return errors.New("Unknown index "+strconv.Itoa(int(fromIndex)))
    }
    if _, ok := g.nodes[toIndex]; !ok {
        return errors.New("Unknown index "+strconv.Itoa(int(toIndex)))
    }
    g.edges = append(g.edges, edge{fromIndex, toIndex})
    return nil
}

func (g *orientedGraphImpl) GetDataForIndex(index int) (interface{}, error) {
    if n, ok := g.nodes[index]; ok {
        return n.data, nil
    } else {
        return nil, errors.New("Unknown index "+strconv.Itoa(int(index)))
    }
}

func (g *orientedGraphImpl) TopologicalSort() ([]int, error) {
    nodes := map[int]NodeColor{}
    for index, _ := range g.nodes {
        nodes[index] = ColorWhite
    }
    ls := list.New()
    for i := 0; i < g.lastNotUsedIndex; i++ {
        if e := g.iterTopologicalSort(&nodes, i, ls); e != nil {
            return nil, e
        }
    }
    var res []int
    for v := range ls.Iterate() {
        res = append(res, v.(int))
    }
    return res, nil
}

func (g *orientedGraphImpl) iterTopologicalSort(
    nodes *map[int]NodeColor,
    curNode int,
    res list.List,
) error {
    color := (*nodes)[curNode]
    if color == ColorBlack {
        return nil
    } else if color == ColorGrey {
        return errors.New("Topological sort is impossible, the graph is not acyclic (index " + strconv.Itoa(curNode) + ")")
    } else {
        (*nodes)[curNode] = ColorGrey
        adjacencies, e := g.getAdjacencies(curNode)
        if e != nil {
            return e
        }
        for _, a := range adjacencies {
            if e = g.iterTopologicalSort(nodes, a, res); e != nil {
                return e
            }
        }
        (*nodes)[curNode] = ColorBlack
        res.Append(curNode)
        return nil
    }
}

func (g *orientedGraphImpl) getAdjacencies(fromIndex int) ([]int, error) {
    var res []int
    if _, ok := g.nodes[fromIndex]; !ok {
        return nil, errors.New("Unknown value "+strconv.Itoa(int(fromIndex)))
    }
    for _, edge := range g.edges {
        if edge.from == fromIndex {
            res = append(res, edge.to)
        }
    }
    return res, nil
}

//endregion
