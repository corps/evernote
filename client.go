package evernote

import (
	"github.com/corps/evernote/edamnotestore"
	"github.com/corps/evernote/edamuserstore"
	"github.com/samuel/go-thrift/thrift"
	"net/rpc"
)

type EvernoteUserStoreUrl string

const (
	SANDBOX EvernoteUserStoreUrl = "https://sandbox.evernote.com/edam/user"
)

type EvernoteClient struct {
	AccessToken  string
	UserStoreUrl string
	NoteStoreUrl string
}

func NewEvernoteClient(userStoreUrl EvernoteUserStoreUrl, accessToken string) *EvernoteClient {
	return &EvernoteClient{
		AccessToken:  accessToken,
		UserStoreUrl: string(userStoreUrl),
	}
}

func newRpcClient(url string) *rpc.Client {
	return thrift.NewClient(NewTHttpClientTransport(url), thrift.NewBinaryProtocol(false, true), false)
}

// May make a network call to the user store in order to obtain the url for the note store client.
// IF so, the result is set on the the client (thus future calls should not make a network call.)
func (client *EvernoteClient) FetchNoteStore() (*edamnotestore.NoteStoreClient, error) {
	if len(client.NoteStoreUrl) == 0 {
		userStore := client.GetUserStore()
		url, err := userStore.GetNoteStoreUrl(&client.AccessToken)
		if err != nil {
			return nil, err
		}

		client.NoteStoreUrl = *url
	}

	return client.GetNoteStore(), nil
}

// Only works if a NoteStoreUrl has been provided.  Otherwise, use FetchNoteStore
func (client *EvernoteClient) GetNoteStore() *edamnotestore.NoteStoreClient {
	return &edamnotestore.NoteStoreClient{Client: newRpcClient(client.NoteStoreUrl)}
}

func (client *EvernoteClient) GetUserStore() *edamuserstore.UserStoreClient {
	return &edamuserstore.UserStoreClient{Client: newRpcClient(client.UserStoreUrl)}
}
