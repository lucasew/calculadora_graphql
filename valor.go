package main

import (
	"math"
)

type Valor struct {
	Chave string
}

func (e Valor) CHAVE() *string {
	log.Verbose(0, "CHAVE %s", e.Chave)
	return &e.Chave
}

func (e *Valor) NUMERO() *float64 {
	v := valores.E[e.Chave]
	log.Verbose(0, "NUMERO %s = %f", e.Chave, v)
	return &v
}

func (e *Valor) Incrementar() *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v++
	log.Verbose(0, "Incrementar %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Decrementar() *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v--
	log.Verbose(0, "Decrementar %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Multiplicar(n struct{ N *float64 }) *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v *= *n.N
	log.Verbose(0, "Multiplicar %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Dividir(n struct{ N *float64 }) *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v /= *n.N
	log.Verbose(0, "Dividir %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Somar(n struct{ N *float64 }) *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v += *n.N
	log.Verbose(0, "Somar %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Subtrair(n struct{ N *float64 }) *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v -= *n.N
	log.Verbose(0, "Subtrair %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Raiz() *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v = math.Sqrt(v)
	log.Verbose(0, "Raiz %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Potencia(n struct{ N *float64 }) *Valor {
	v, ok := valores.E[e.Chave]
	if !ok {
		v = 0
	}
	v = math.Pow(v, *n.N)
	log.Verbose(0, "Potencia %s -> %f", e.Chave, v)
	valores.E[e.Chave] = v
	return &Valor{Chave: e.Chave}
}

func (e *Valor) Alterar(v struct{ N *float64 }) *Valor {
	valores.E[e.Chave] = *v.N
	log.Verbose(0, "Alterar %s -> %f", e.Chave, *v.N)
	return e
}
