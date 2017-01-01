package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Park", func() {
	var (
		park Park
	)

	BeforeEach(func() {
		park = Park{
			Name:        "Boyd Lake",
			Street:      "3720 North County Road",
			City:        "Loveland",
			Zip:         "80538",
			Email:       "boyd.lake@state.co.us",
			Description: "Colorful sailboats skimming blue water.",
			URL:         "http://cpw.state.co.us/placestogo/parks/BoydLake",
		}
	})

	Describe("Database Transactions", func() {
		Context("InsertPark", func() {
			It("should insert a park into the database", func() {
			})
		})
	})
})
