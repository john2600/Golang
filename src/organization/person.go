package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Employee struct {
	Name
}

type Name struct {
	First string
	last  string
}

type Person struct {
	Name
	First          string
	last           string
	twitterHandler TwitterHandler
	Citizen
	Conflict
}

type Handler struct {
	handle string
	name   string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.First, n.last)
}

func (h Handler) randomFun() {
}

type TwitterHandler string

type Identificable interface {
	ID() string
}

type Citizen interface {
	Identificable
	Country() []string
}

type Conflict interface {
	ID() string
}

type socialSecurityNumber string

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Citizen() string {
	return "United State of America"
}

func NewSocialSecurityNumber(value string) Identificable {
	return socialSecurityNumber(value)
}

type europeanUnionIdentifier struct {
	id      string
	country []string
}

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: []string{country},
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: []string{country},
		}
	case europeanUnionIdentifier:
		return v
	case Person:
		euID, _ := v.Citizen.(europeanUnionIdentifier)
		return euID
	default:
		panic("using invalid type to initialize EU Identifier")
	}

}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() []string {
	return eui.country
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{First: firstName,
			last: lastName,
		},
		Citizen: citizen,
	}
}
func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")

	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

func (p *Person) ID() string {
	//return p.Citizen.ID()
	return p.Citizen.ID()
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}
func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
