package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct{
		input http.Header
		want string	
	}{
		{input: http.Header{"Authorization": []string{"ApiKey @1234"}},		    want: "@1234" },
		{input: http.Header{"Authorization": []string{"ApiKey @77712371723"}},  want: "@77712371723" },
		{input: http.Header{"Authorization": []string{"blubkey @1234"}},		want: "malformed authorization header" },
		{input:  http.Header{"Authorization": []string{"blubkey@77712371723"}}, want: "malformed authorization header" },
		{input:  http.Header{"Authorization": []string{"blubkey @777123 71723"}}, want: "malformed authorization header" },
	} 

for _, tc := range tests {
      got, err := GetAPIKey(tc.input)
      if err != nil {
              got = err.Error()
      }
      if got != tc.want {
              t.Errorf("got: %q, want: %q. That ain't right", got, tc.want)
      }
}
}