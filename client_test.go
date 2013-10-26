package evernote

import (
	"fmt"
	"github.com/corps/evernote/edamtypes"
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

	content := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<!DOCTYPE en-note SYSTEM "http://xml.evernote.com/pub/enml2.dtd">
	<en-note><h1>This note is soooo gooood</h1></en-note>`
	title := "Yo momma"
	note, err := noteStore.CreateNote(&accessToken, &edamtypes.Note{Title: &title, Content: &content, NotebookGuid: notebook.Guid})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created a note: ", *note.Guid)
}
