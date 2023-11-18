// Code generated by github.com/vanstinator/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/vanstinator/genqlient/graphql"
	"github.com/vanstinator/genqlient/internal/testutil"
)

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type UserQueryInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         testutil.ID      `json:"id"`
	Role       Role             `json:"role"`
	Names      []string         `json:"names"`
	HasPokemon testutil.Pokemon `json:"hasPokemon"`
	Birthdate  time.Time        `json:"-"`
}

// GetEmail returns UserQueryInput.Email, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetEmail() string { return v.Email }

// GetName returns UserQueryInput.Name, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetName() string { return v.Name }

// GetId returns UserQueryInput.Id, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetId() testutil.ID { return v.Id }

// GetRole returns UserQueryInput.Role, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetRole() Role { return v.Role }

// GetNames returns UserQueryInput.Names, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetNames() []string { return v.Names }

// GetHasPokemon returns UserQueryInput.HasPokemon, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetHasPokemon() testutil.Pokemon { return v.HasPokemon }

// GetBirthdate returns UserQueryInput.Birthdate, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetBirthdate() time.Time { return v.Birthdate }

func (v *UserQueryInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*UserQueryInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoUnmarshalJSON
	}
	firstPass.UserQueryInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Birthdate
		src := firstPass.Birthdate
		if len(src) != 0 && string(src) != "null" {
			err = testutil.UnmarshalDate(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}
	return nil
}

type __premarshalUserQueryInput struct {
	Email string `json:"email"`

	Name string `json:"name"`

	Id testutil.ID `json:"id"`

	Role Role `json:"role"`

	Names []string `json:"names"`

	HasPokemon testutil.Pokemon `json:"hasPokemon"`

	Birthdate json.RawMessage `json:"birthdate"`
}

func (v *UserQueryInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *UserQueryInput) __premarshalJSON() (*__premarshalUserQueryInput, error) {
	var retval __premarshalUserQueryInput

	retval.Email = v.Email
	retval.Name = v.Name
	retval.Id = v.Id
	retval.Role = v.Role
	retval.Names = v.Names
	retval.HasPokemon = v.HasPokemon
	{

		dst := &retval.Birthdate
		src := v.Birthdate
		var err error
		*dst, err = testutil.MarshalDate(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal UserQueryInput.Birthdate: %w", err)
		}
	}
	return &retval, nil
}

// __unexportedInput is used internally by genqlient
type __unexportedInput struct {
	Query UserQueryInput `json:"query"`
}

// GetQuery returns __unexportedInput.Query, and is useful for accessing the field via an interface.
func (v *__unexportedInput) GetQuery() UserQueryInput { return v.Query }

// unexportedResponse is returned by unexported on success.
type unexportedResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User unexportedUser `json:"user"`
}

// GetUser returns unexportedResponse.User, and is useful for accessing the field via an interface.
func (v *unexportedResponse) GetUser() unexportedUser { return v.User }

// unexportedUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type unexportedUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns unexportedUser.Id, and is useful for accessing the field via an interface.
func (v *unexportedUser) GetId() testutil.ID { return v.Id }

// The query or mutation executed by unexported.
const unexported_Operation = `
query unexported ($query: UserQueryInput) {
	user(query: $query) {
		id
	}
}
`

func unexported(
	client_ graphql.Client,
	query UserQueryInput,
) (*unexportedResponse, error) {
	req_ := &graphql.Request{
		OpName: "unexported",
		Query:  unexported_Operation,
		Variables: &__unexportedInput{
			Query: query,
		},
	}
	var err_ error

	var data_ unexportedResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

