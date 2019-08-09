package main

func (e *Estado) SetEstado(p struct {
	Key   *string
	Valor *float64
}) (*Valor, error) {
	log.Verbose(0, "M SetEstado %s = %.f", *p.Key, *p.Valor)
	e.E[*p.Key] = *p.Valor
	return &Valor{Chave: *p.Key}, nil
}
