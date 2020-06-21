package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/udit-gr/graphql-go/config"
)

var authorType = graphql.NewObject{
	graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
}

var tutorialType = graphql.NewObject{
	graphql.ObjectConfig{
		Name: "tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"likes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
}

func main() {

	authors := config.populateAuthors()
	tutorials := config.populateTutorials()

	// Schema for the queries
	fields := graphql.Fields{
		"authors": &graphql.Field{
			Type:        graphql.NewList(authorType),
			Description: "Get Authors List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return authors, nil
			},
		},
		"tutorials": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorials List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("[ERR] Failed to create new schema, Err : %+v", err)
	}

	// Query
	query := `
    {
      authors{
        name
      }
    }
  `

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Printf("[ERR] Failed to execute graphql operation, Err : %+v", err)
		return
	}

	result, err := json.Marhsal(r)
	if err != nil {
		log.Printf("[ERR] Error while un-marshalling json data, Err : %+v", err)
		return
	}

	log.Printf("Result : %s\n", string(result))

}
