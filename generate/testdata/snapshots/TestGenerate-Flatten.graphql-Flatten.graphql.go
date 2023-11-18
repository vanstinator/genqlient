// Code generated by github.com/vanstinator/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/vanstinator/genqlient/graphql"
	"github.com/vanstinator/genqlient/internal/testutil"
)

// ChildVideoFields includes the GraphQL fields of Video requested by the fragment ChildVideoFields.
type ChildVideoFields struct {
	// ID is documented in the Content interface.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns ChildVideoFields.Id, and is useful for accessing the field via an interface.
func (v *ChildVideoFields) GetId() testutil.ID { return v.Id }

// GetName returns ChildVideoFields.Name, and is useful for accessing the field via an interface.
func (v *ChildVideoFields) GetName() string { return v.Name }

// ContentFields includes the GraphQL fields of Content requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
//
// ContentFields is implemented by the following types:
// ContentFieldsArticle
// ContentFieldsVideo
// ContentFieldsTopic
type ContentFields interface {
	implementsGraphQLInterfaceContentFields()
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
	// GetUrl returns the interface-field "url" from its implementation.
	GetUrl() string
}

func (v *ContentFieldsArticle) implementsGraphQLInterfaceContentFields() {}
func (v *ContentFieldsVideo) implementsGraphQLInterfaceContentFields()   {}
func (v *ContentFieldsTopic) implementsGraphQLInterfaceContentFields()   {}

func __unmarshalContentFields(b []byte, v *ContentFields) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(ContentFieldsArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ContentFieldsVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ContentFieldsTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for ContentFields: "%v"`, tn.TypeName)
	}
}

func __marshalContentFields(v *ContentFields) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *ContentFieldsArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*ContentFieldsArticle
		}{typename, v}
		return json.Marshal(result)
	case *ContentFieldsVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*ContentFieldsVideo
		}{typename, v}
		return json.Marshal(result)
	case *ContentFieldsTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*ContentFieldsTopic
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for ContentFields: "%T"`, v)
	}
}

// ContentFields includes the GraphQL fields of Article requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ContentFieldsArticle struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// GetName returns ContentFieldsArticle.Name, and is useful for accessing the field via an interface.
func (v *ContentFieldsArticle) GetName() string { return v.Name }

// GetUrl returns ContentFieldsArticle.Url, and is useful for accessing the field via an interface.
func (v *ContentFieldsArticle) GetUrl() string { return v.Url }

// ContentFields includes the GraphQL fields of Topic requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ContentFieldsTopic struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// GetName returns ContentFieldsTopic.Name, and is useful for accessing the field via an interface.
func (v *ContentFieldsTopic) GetName() string { return v.Name }

// GetUrl returns ContentFieldsTopic.Url, and is useful for accessing the field via an interface.
func (v *ContentFieldsTopic) GetUrl() string { return v.Url }

// ContentFields includes the GraphQL fields of Video requested by the fragment ContentFields.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ContentFieldsVideo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// GetName returns ContentFieldsVideo.Name, and is useful for accessing the field via an interface.
func (v *ContentFieldsVideo) GetName() string { return v.Name }

// GetUrl returns ContentFieldsVideo.Url, and is useful for accessing the field via an interface.
func (v *ContentFieldsVideo) GetUrl() string { return v.Url }

// InnerQueryFragment includes the GraphQL fields of Query requested by the fragment InnerQueryFragment.
// The GraphQL type's documentation follows.
//
// Query's description is probably ignored by almost all callers.
type InnerQueryFragment struct {
	RandomVideo VideoFields        `json:"randomVideo"`
	RandomItem  ContentFields      `json:"-"`
	OtherVideo  ContentFieldsVideo `json:"otherVideo"`
}

// GetRandomVideo returns InnerQueryFragment.RandomVideo, and is useful for accessing the field via an interface.
func (v *InnerQueryFragment) GetRandomVideo() VideoFields { return v.RandomVideo }

// GetRandomItem returns InnerQueryFragment.RandomItem, and is useful for accessing the field via an interface.
func (v *InnerQueryFragment) GetRandomItem() ContentFields { return v.RandomItem }

// GetOtherVideo returns InnerQueryFragment.OtherVideo, and is useful for accessing the field via an interface.
func (v *InnerQueryFragment) GetOtherVideo() ContentFieldsVideo { return v.OtherVideo }

func (v *InnerQueryFragment) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InnerQueryFragment
		RandomItem json.RawMessage `json:"randomItem"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InnerQueryFragment = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomItem
		src := firstPass.RandomItem
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalContentFields(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal InnerQueryFragment.RandomItem: %w", err)
			}
		}
	}
	return nil
}

type __premarshalInnerQueryFragment struct {
	RandomVideo VideoFields `json:"randomVideo"`

	RandomItem json.RawMessage `json:"randomItem"`

	OtherVideo ContentFieldsVideo `json:"otherVideo"`
}

func (v *InnerQueryFragment) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InnerQueryFragment) __premarshalJSON() (*__premarshalInnerQueryFragment, error) {
	var retval __premarshalInnerQueryFragment

	retval.RandomVideo = v.RandomVideo
	{

		dst := &retval.RandomItem
		src := v.RandomItem
		var err error
		*dst, err = __marshalContentFields(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal InnerQueryFragment.RandomItem: %w", err)
		}
	}
	retval.OtherVideo = v.OtherVideo
	return &retval, nil
}

// VideoFields includes the GraphQL fields of Video requested by the fragment VideoFields.
type VideoFields struct {
	// ID is documented in the Content interface.
	Id     testutil.ID            `json:"id"`
	Parent VideoFieldsParentTopic `json:"parent"`
}

// GetId returns VideoFields.Id, and is useful for accessing the field via an interface.
func (v *VideoFields) GetId() testutil.ID { return v.Id }

// GetParent returns VideoFields.Parent, and is useful for accessing the field via an interface.
func (v *VideoFields) GetParent() VideoFieldsParentTopic { return v.Parent }

// VideoFieldsParentTopic includes the requested fields of the GraphQL type Topic.
type VideoFieldsParentTopic struct {
	VideoChildren []ChildVideoFields `json:"videoChildren"`
}

// GetVideoChildren returns VideoFieldsParentTopic.VideoChildren, and is useful for accessing the field via an interface.
func (v *VideoFieldsParentTopic) GetVideoChildren() []ChildVideoFields { return v.VideoChildren }

// The query or mutation executed by ComplexNamedFragments.
const ComplexNamedFragments_Operation = `
query ComplexNamedFragments {
	... QueryFragment
}
fragment QueryFragment on Query {
	... InnerQueryFragment
}
fragment InnerQueryFragment on Query {
	randomVideo {
		... VideoFields
	}
	randomItem {
		__typename
		... ContentFields
	}
	otherVideo: randomVideo {
		... ContentFields
	}
}
fragment VideoFields on Video {
	id
	parent {
		videoChildren {
			... ChildVideoFields
		}
	}
}
fragment ContentFields on Content {
	name
	url
}
fragment ChildVideoFields on Video {
	id
	name
}
`

func ComplexNamedFragments(
	client_ graphql.Client,
) (*InnerQueryFragment, error) {
	req_ := &graphql.Request{
		OpName: "ComplexNamedFragments",
		Query:  ComplexNamedFragments_Operation,
	}
	var err_ error

	var data_ InnerQueryFragment
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

