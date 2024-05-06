package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseBody(t *testing.T) {
	data := `{"Name":"Bob","Age":25}`
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(data))
	var user struct {
		Name string
		Age  int
	}

	err := ParseBody(r, &user)
	if err != nil {
		t.Fatalf("ParseBody failed: %s", err)
	}
	if user.Name != "Bob" || user.Age != 25 {
		t.Errorf("ParseBody parsed %v, expected %v", user, struct {
			Name string
			Age  int
		}{"Bob", 25})
	}
}

// func TestValidateJSON(t *testing.T) {
// 	validJSON := `{"Name":"Charlie","Age":29}`
// 	invalidJSON := `{"Name":,"Age":29}` // deliberately broken JSON

// 	if err := ValidateJSON(validJSON); err != nil {
// 		t.Errorf("ValidateJSON marked valid JSON as invalid")
// 	}

// 	if err := ValidateJSON(invalidJSON); err == nil {
// 		t.Errorf("ValidateJSON did not catch errors in invalid JSON")
// 	}
// }

func TestToJson(t *testing.T) {
	user := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}

	expectedJSON := `{"Name":"Alice","Age":30}`
	jsonString, err := ToJson(user)
	if err != nil {
		t.Fatalf("ToJson failed: %s", err)
	}

	if jsonString != expectedJSON {
		t.Errorf("ToJson was incorrect, got: %s, want: %s.", jsonString, expectedJSON)
	}
}
