// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/fujiwara402/goa-sample-app/design
// --out=$(GOPATH)/src/github.com/fujiwara402/goa-sample-app
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// A bottle of wine (default view)
//
// Identifier: application/vnd.goa.example.bottle+json; view=default
type GoaExampleBottle struct {
	// API href for making requests on the bottle
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique bottle ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the GoaExampleBottle media type instance.
func (mt *GoaExampleBottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}
