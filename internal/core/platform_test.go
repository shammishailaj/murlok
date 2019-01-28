package core

import (
	"encoding/json"
	"testing"

	"github.com/pkg/errors"
)

type input struct {
	Name string
}

type badInput struct {
	Name string
	Func func()
}

type output struct {
	Greeting string
}

func TestPlatform(t *testing.T) {
	tests := []struct {
		scenario string
		method   string
		in       interface{}
		expected output
		err      bool
	}{
		{
			scenario: "call succeed",
			method:   "test.Greet",
			in:       input{Name: "Maxence"},
			expected: output{Greeting: "Hello, Maxence"},
		},
		{
			scenario: "call with input containing a func returns an error",
			method:   "test.Greet",
			in:       badInput{Name: "Maxence"},
			err:      true,
		},
		{
			scenario: "call unsupported method returns an error",
			method:   "test.Bye",
			in:       input{Name: "Maxence"},
			err:      true,
		},
		{
			scenario: "call a method that errors returns an error",
			method:   "test.Error",
			in:       input{Name: "Maxence"},
			err:      true,
		},
		{
			scenario: "call a method without with empty output succeed",
			method:   "test.EmptyOutput",
			in:       input{Name: "Maxence"},
		},
	}

	platform := Platform{}

	platform.Handler = func(rawcall string) error {
		call := platformCall{}
		if err := json.Unmarshal([]byte(rawcall), &call); err != nil {
			return err
		}

		name := call.In.(map[string]interface{})["Name"].(string)

		out, err := json.Marshal(output{
			Greeting: "Hello, " + name,
		})
		if err != nil {
			return err
		}

		switch call.Method {
		case "test.Greet":
			platform.Return(call.ReturnID, string(out), "")
			return nil

		case "test.Error":
			platform.Return(call.ReturnID, "", "error!")
			return nil

		case "test.EmptyOutput":
			platform.Return(call.ReturnID, "", "")
			return nil

		default:
			return errors.Errorf("%s: unknown platform method", call.Method)
		}
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			out := output{}

			err := platform.Call(test.method, &out, test.in)
			if test.err {
				if err == nil {
					t.Fatal("error is nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}

			if out != test.expected {
				t.Errorf("expected output: %+v", test.expected)
				t.Errorf("returned output: %+v", out)
			}
		})
	}
}

func TestPlatformReturnPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	platform := Platform{}
	platform.Return("test", "", "")
	t.Error("no panic")
}
