package loader

type Repository struct {
    packages map[string]map[string]func()interface{}
}

func NewRepository() Repository {
    return Repository{
        map[string]map[string]func()interface{}{},
    }
}

func (r *Repository) Register(module string, name string, value func()interface{}) {
    if r.packages[module] == nil {
        r.packages[module] = map[string]func()interface{}{}
    }
    r.packages[module][name] = value
}

