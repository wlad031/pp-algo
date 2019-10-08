package set

//region public

type OrderedSet interface {
    Set
    Append(v interface{})
}

func NewOrderedSet() OrderedSet {
    return &orderedSetImpl{ls: []interface{}{}}
}

//endregion

//region private

type orderedSetImpl struct {
    ls []interface{}
}

func (s *orderedSetImpl) Contains(v interface{}) bool {
    for _, k := range s.ls {
        if v == k {
            return true
        }
    }
    return false
}

func (s *orderedSetImpl) Add(v interface{}) {
    s.Append(v)
}

func (s *orderedSetImpl) Append(v interface{}) {
    // FIXME: performance of checking duplicates
    if s.Contains(v) {
        return
    }
    s.ls = append(s.ls, v)
}

func (s *orderedSetImpl) Iterate() <-chan interface{} {
    c := make(chan interface{})
    go func() {
        for _, v := range s.ls {
            c <- v
        }
        close(c)
    }()
    return c
}

//endregion