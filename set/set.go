package set

//region public

type Set interface {
    Add(v interface{})
    Contains(v interface{}) bool
    Iterate() <-chan interface{}
}

func NewSet() Set {
    return &setImpl{ls: map[interface{}]interface{}{}}
}

//endregion

//region private

type setImpl struct {
    ls map[interface{}]interface{}
}

func (s *setImpl) Contains(v interface{}) bool {
    _, ok := s.ls[v]
    return ok
}

func (s *setImpl) Add(v interface{}) {
    if s.Contains(v) {
        return
    }
    s.ls[v] = nil
}

func (s *setImpl) Iterate() <-chan interface{} {
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