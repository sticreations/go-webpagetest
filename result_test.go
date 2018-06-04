package webpagetest

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingResultWithPlrAsString(t *testing.T) {
	var response, err = ioutil.ReadFile("./testdata/TestResultPlrAsString.json")
	assert.Nil(t, err)
	_, err = parseResultResponse(response)
	if err != nil {
		fmt.Errorf("ERROR: %v", err)
	}
	assert.Nil(t, err)
}
