package chain_of_responsibility

import (
	"encoding/json"
	"testing"
)

func TestCORPattern(t *testing.T) {
	book := Book{"Go design patterns", "Ismayil Malik"}
	bytes, _ := json.Marshal(book)
	bookDb := &BookDb{
		make(map[string]Book),
		nil,
	}
	logger := &Logger{
		next: bookDb,
	}
	chain := &Deserializer{
		next: logger,
	}

	t.Run(`It will be unmurshalled on deserializer handler then will be logged
		  on logger handler and at the end will be presisted by third handler`, func(t *testing.T) {

		err := chain.Handle(bytes)
		if err != nil {
			t.Fatalf("Something went wrong: %s", err.Error())
		}

		storedBook := bookDb.store[book.Title]
		if storedBook != book {
			t.Errorf("Expected %b but got %b", book, storedBook)
		}
	})

}
