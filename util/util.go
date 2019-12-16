package util

import (
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	matlasClient "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func CreateMongoDBClient(publicKey, privateKey string) (*matlasClient.Client, error) {
	// setup a transport to handle digest
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	//Initialize the MongoDB Atlas API Client.
	return matlasClient.NewClient(client), nil
}
