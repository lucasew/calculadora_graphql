package main

func (e *Estado) GetAll() (*[]*Valor, error) {
	log.Verbose(0, "Q GetAll")
	ret := make([]*Valor, len(e.E))
	i := 0
	for k := range e.E {
		ret[i] = &Valor{
			Chave: k,
		}
		i++
	}
	return &ret, nil
}

func (e *Estado) GetEstado(p struct{ Key *string }) (*Valor, error) {
	log.Verbose(0, "M GetEstado %s", p.Key)
	_, ok := e.E[*p.Key]
	if !ok {
		return nil, log.Error("key não existe")
	}
	return &Valor{Chave: *p.Key}, nil
}
