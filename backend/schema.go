package main

import (
	"context" // Import the context package
	firebase "firebase.google.com/go/v4"
	"github.com/graphql-go/graphql"
)

// Quote represents a quote structure
type Quote struct {
	PostContent string `firestore:"content"`
	Author      string `firestore:"author"`
	Title       string `firestore:"title"`
}

// Create a GraphQL ObjectType for Quote
var quoteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Quote",
	Fields: graphql.Fields{
		"postContent": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Define the input type for the CreateQuote mutation
var createQuoteInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateQuoteInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"postContent": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"author": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

// Define the root Query for GraphQL
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"quote": &graphql.Field{
			Type: quoteType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Return a static quote for demonstration purposes
				return &Quote{
					PostContent: "To love is to live",
					Author:      "Dhruv Puri",
					Title:       "Love & Life",
				}, nil
			},
		},
	},
})

// Define the root Mutation for GraphQL
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createQuote": &graphql.Field{
			Type: quoteType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(createQuoteInputType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				input := p.Args["input"].(map[string]interface{})
				newQuote := Quote{
					PostContent: input["postContent"].(string),
					Author:      input["author"].(string),
					Title:       input["title"].(string),
				}

				// Write to Firestore
				app := p.Context.Value("firebaseApp").(*firebase.App)
				client, err := app.Firestore(context.Background())
				if err != nil {
					return nil, err
				}
				defer client.Close()

				_, _, err = client.Collection("posts").Add(context.Background(), newQuote)
				if err != nil {
					return nil, err
				}

				return newQuote, nil
			},
		},
	},
})

// Define the GraphQL schema
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})
