package main

import (
	"example/organization"
	"fmt"
)

func main() {
	//var p organization.Identificable =organization.Person{}
	//p :=organization.NewPerson("John","Wilson", organization.NewSocialSecurityNumber("123-45-6789"))
	p := organization.NewPerson("John", "Wilson", organization.NewEuropeanUnionIdentifier("adfad-123432", "Poland"))
	err := p.SetTwitterHandler("@sample")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	/*
		fmt.Println(p.TwitterHandler())
		fmt.Println(p.TwitterHandler().RedirectUrl())
		p.First = "colling"
		fmt.Println(p.First)
		fmt.Println(p.Name.First)
		fmt.Println(p.FullName())
		fmt.Println(p.ID())
		fmt.Println(p.Country())
	*/
	name1 := Name{First: "", Last: ""}

	//name2 := OtherName{First:"James", Last:"Wilson2"}

	//ssn :=organization.NewSocialSecurityNumber("123-45-6789")
	//eu :=organization.NewEuropeanUnionIdentifier("12345","france")
	//eu2 :=organization.NewEuropeanUnionIdentifier("12345","france")
	portfolio := map[Name][]organization.Person{}
	portfolio[name1] = []organization.Person{p}

	//if name1 == (Name{}) {
	//	fmt.Println("we match")
	//}
	fmt.Println(p.ID())
	fmt.Println(p.Country())

}

type Name struct {
	First string
	Last  string
	//Middle [] string
}

type OtherName struct {
	First string
	Last  string
	//Middle [] string
}
