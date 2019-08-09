package main

import (
	"fmt"
	"github.com/friendsofgo/graphiql"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/lucasew/golog"
	"io/ioutil"
	"net/http"
)

var valores = Estado{
	E: map[string]float64{},
}

type Estado struct {
	E map[string]float64
}

const porta = 8080

var log = golog.Default

func main() {
	ui, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", ui)
	f, err := ioutil.ReadFile("./schema.graphql")
	if err != nil {
		log.Panic(err.Error())
	}
	schema := gql.MustParseSchema(string(f), &valores)
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Info("Iniciando servidor na porta %d...", porta)
	http.ListenAndServe(fmt.Sprintf(":%d", porta), nil)
}
