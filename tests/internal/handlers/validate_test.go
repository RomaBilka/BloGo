package handlers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name    string
		request createValidateateUserRequest
		ok      bool
		err     error
	}{
		{
			name: "Ok",
			request: createValidateateUserRequest{
				Name:  "tesr",
				Email: "test@tt.com",
				Phone: "1234567890",
			},
			ok: true,
		},
		{
			name: "Bad request, short user name",
			request: createValidateateUserRequest{
				Name:  "",
				Email: "test@tt.com",
				Phone: "1234567890",
			},
			ok: false,
			err: fmt.Errorf("%s", "Bad request, short user name"),
		},
		{
			name: "Bad request, wrong user email",
			request: createValidateateUserRequest{
				Name:  "test",
				Email: "test@com",
				Phone: "1234567890",
			},
			ok: false,
			err: fmt.Errorf("%s", "Bad request, wrong user email"),
		},
		{
			name: "Bad request, short user phone",
			request: createValidateateUserRequest{
				Name:  "test",
				Email: "test@tt.com",
				Phone: "123",
			},
			ok: false,
			err: fmt.Errorf("%s", "Bad request, short user phone"),
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ok, err := testCase.request.validate()

			if !ok {
				assert.Equal(t, err, testCase.err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
