package logic

import (
	"testing"

	"github.com/kiritocyanpine/go-tiny-url/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func cleanupTinyURLInstance() {
	tinyUrlInstance = nil
}

func TestUrlAddFunction(t *testing.T) {
	mockDB := mocks.NewPersistant(t)

	mockDB.On("Set", mock.Anything, mock.Anything).Return(nil).Once()

	tinyurl := CreateTinyUrl(mockDB)

	tinyurl.AddNewUrlQuery("a")
	cleanupTinyURLInstance()
}

func TestGetHasedValue(t *testing.T) {

	testData := []struct {
		testName       string
		url            string
		expectedLength int
	}{
		{
			testName:       "Url length less than 10 characters",
			url:            "go.org",
			expectedLength: 12,
		},
		{
			testName:       "Url length equal to 10 characters",
			url:            "golang.com",
			expectedLength: 12,
		},
		{
			testName:       "Url length more than 10 characters",
			url:            "https://github.com/KiritoCyanPine?tab=repositories",
			expectedLength: 12,
		},
		{
			testName:       "Url length 0",
			url:            "",
			expectedLength: 12,
		},
	}

	for _, test := range testData {
		mockDB := mocks.NewPersistant(t)
		tinyurl := CreateTinyUrl(mockDB)
		mockDB.On("Set", mock.Anything, mock.Anything).Return(nil)

		result, err := tinyurl.AddNewUrlQuery(test.url)
		assert.NoError(t, err)
		assert.Equal(t, len(result), test.expectedLength, "%s :hash length should always be equal", test.testName)

		cleanupTinyURLInstance()
	}
}
