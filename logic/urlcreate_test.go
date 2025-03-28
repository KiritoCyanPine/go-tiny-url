package logic

import (
	"testing"

	"github.com/kiritocyanpine/go-tiny-url/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUrlAddFunction(t *testing.T) {
	mockDB := mocks.NewPersistant(t)

	tinyurl := CreateTinyUrl(mockDB)

	tinyurl.AddNewUrlQuery("a")
}

func TestGetHasedValue(t *testing.T) {
	mockDB := mocks.NewPersistant(t)

	mockDB.On("Set", mock.Anything, mock.Anything).Return(nil)
	tinyurl := CreateTinyUrl(mockDB)

	testData := []struct {
		testName       string
		url            string
		expectedLength int
	}{
		{
			testName:       "Url length less than 10 characters",
			url:            "go.org",
			expectedLength: 10,
		},
		{
			testName:       "Url length equal to 10 characters",
			url:            "golang.com",
			expectedLength: 10,
		},
		{
			testName:       "Url length more than 10 characters",
			url:            "https://github.com/KiritoCyanPine?tab=repositories",
			expectedLength: 10,
		},
		{
			testName:       "Url length 0",
			url:            "",
			expectedLength: 10,
		},
	}

	for _, test := range testData {
		result, err := tinyurl.AddNewUrlQuery(test.url)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, len(result), test.expectedLength, "%s :hash length should always be equal", test.testName)
	}

}
