// Documentation of the articles API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:10000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package api

import "example-go-api/repository"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body string
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body string
}

// A list of articles
// swagger:response articlesResponse
type articlesResponseWrapper struct {
	// All current articles
	// in: body
	Body []repository.Article
}

// Data structure representing a single article
// swagger:response articleResponse
type articleResponseWrapper struct {
	// Newly created article
	// in: body
	Body repository.Article
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}
