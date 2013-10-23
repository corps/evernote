package evernote

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestIntegration(t *testing.T) {
	data, err := ioutil.ReadFile("accessToken")
	if err != nil {
		fmt.Println(err)	
		t.FailNow()
	}

	accessToken := string(data)

	client := NewEvernoteClient(SANDBOX, accessToken)
	
	url, err := client.GetUserStore().GetNoteStoreUrl(&accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("Got note store url: ", *url)

	noteStore, err := client.FetchNoteStore()
	if err != nil {
		panic(err)
	}

	notebook, err := noteStore.GetDefaultNotebook(&accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("Got default notebook's name: ", *notebook.Name)
}
