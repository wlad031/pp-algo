package list

//region public

type List interface {
    Append(v interface{})
    Iterate() <-chan interface{}
}

func New() List {
    return &listImpl{ls: []interface{}{}}
}

//endregion

//region private

type listImpl struct {
    ls []interface{}
}

func (s *listImpl) Add(v interface{}) {
    s.Append(v)
}

func (s *listImpl) Append(v interface{}) {
    s.ls = append(s.ls, v)
}

func (s *listImpl) Iterate() <-chan interface{} {
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