package sql_funcs

import "testing"

func TestGetArticleThumbnail(t *testing.T) {
	result, err := GetArticleThumbnail(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("%+v", result)
}
