// Code generated by github.com/vanstinator/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/vanstinator/genqlient/graphql"
	"github.com/vanstinator/genqlient/internal/testutil"
)

// CustomMarshalSliceResponse is returned by CustomMarshalSlice on success.
type CustomMarshalSliceResponse struct {
	AcceptsListOfListOfListsOfDates bool `json:"acceptsListOfListOfListsOfDates"`
	WithPointer                     bool `json:"withPointer"`
}

// GetAcceptsListOfListOfListsOfDates returns CustomMarshalSliceResponse.AcceptsListOfListOfListsOfDates, and is useful for accessing the field via an interface.
func (v *CustomMarshalSliceResponse) GetAcceptsListOfListOfListsOfDates() bool {
	return v.AcceptsListOfListOfListsOfDates
}

// GetWithPointer returns CustomMarshalSliceResponse.WithPointer, and is useful for accessing the field via an interface.
func (v *CustomMarshalSliceResponse) GetWithPointer() bool { return v.WithPointer }

// __CustomMarshalSliceInput is used internally by genqlient
type __CustomMarshalSliceInput struct {
	Datesss  [][][]time.Time  `json:"-"`
	Datesssp [][][]*time.Time `json:"-"`
}

// GetDatesss returns __CustomMarshalSliceInput.Datesss, and is useful for accessing the field via an interface.
func (v *__CustomMarshalSliceInput) GetDatesss() [][][]time.Time { return v.Datesss }

// GetDatesssp returns __CustomMarshalSliceInput.Datesssp, and is useful for accessing the field via an interface.
func (v *__CustomMarshalSliceInput) GetDatesssp() [][][]*time.Time { return v.Datesssp }

func (v *__CustomMarshalSliceInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*__CustomMarshalSliceInput
		Datesss  [][][]json.RawMessage `json:"datesss"`
		Datesssp [][][]json.RawMessage `json:"datesssp"`
		graphql.NoUnmarshalJSON
	}
	firstPass.__CustomMarshalSliceInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Datesss
		src := firstPass.Datesss
		*dst = make(
			[][][]time.Time,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]time.Time,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]time.Time,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					if len(src) != 0 && string(src) != "null" {
						err = testutil.UnmarshalDate(
							src, dst)
						if err != nil {
							return fmt.Errorf(
								"unable to unmarshal __CustomMarshalSliceInput.Datesss: %w", err)
						}
					}
				}
			}
		}
	}

	{
		dst := &v.Datesssp
		src := firstPass.Datesssp
		*dst = make(
			[][][]*time.Time,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]*time.Time,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]*time.Time,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					if len(src) != 0 && string(src) != "null" {
						*dst = new(time.Time)
						err = testutil.UnmarshalDate(
							src, *dst)
						if err != nil {
							return fmt.Errorf(
								"unable to unmarshal __CustomMarshalSliceInput.Datesssp: %w", err)
						}
					}
				}
			}
		}
	}
	return nil
}

type __premarshal__CustomMarshalSliceInput struct {
	Datesss [][][]json.RawMessage `json:"datesss"`

	Datesssp [][][]json.RawMessage `json:"datesssp"`
}

func (v *__CustomMarshalSliceInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *__CustomMarshalSliceInput) __premarshalJSON() (*__premarshal__CustomMarshalSliceInput, error) {
	var retval __premarshal__CustomMarshalSliceInput

	{

		dst := &retval.Datesss
		src := v.Datesss
		*dst = make(
			[][][]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]json.RawMessage,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]json.RawMessage,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					var err error
					*dst, err = testutil.MarshalDate(
						&src)
					if err != nil {
						return nil, fmt.Errorf(
							"unable to marshal __CustomMarshalSliceInput.Datesss: %w", err)
					}
				}
			}
		}
	}
	{

		dst := &retval.Datesssp
		src := v.Datesssp
		*dst = make(
			[][][]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]json.RawMessage,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]json.RawMessage,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					if src != nil {
						var err error
						*dst, err = testutil.MarshalDate(
							src)
						if err != nil {
							return nil, fmt.Errorf(
								"unable to marshal __CustomMarshalSliceInput.Datesssp: %w", err)
						}
					}
				}
			}
		}
	}
	return &retval, nil
}

// The query or mutation executed by CustomMarshalSlice.
const CustomMarshalSlice_Operation = `
query CustomMarshalSlice ($datesss: [[[Date!]!]!]!, $datesssp: [[[Date!]!]!]!) {
	acceptsListOfListOfListsOfDates(datesss: $datesss)
	withPointer: acceptsListOfListOfListsOfDates(datesss: $datesssp)
}
`

func CustomMarshalSlice(
	client_ graphql.Client,
	datesss [][][]time.Time,
	datesssp [][][]*time.Time,
) (*CustomMarshalSliceResponse, error) {
	req_ := &graphql.Request{
		OpName: "CustomMarshalSlice",
		Query:  CustomMarshalSlice_Operation,
		Variables: &__CustomMarshalSliceInput{
			Datesss:  datesss,
			Datesssp: datesssp,
		},
	}
	var err_ error

	var data_ CustomMarshalSliceResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

