package util

import (
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	Version = "beta"
)

func CreateMongoDBClient(publicKey, privateKey string) (*mongodbatlas.Client, error) {
	// setup a transport to handle digest
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	//Initialize the MongoDB Atlas API Client.
    atlas := mongodbatlas.NewClient(client)
    atlas.UserAgent = "mongodbatlas-cloudformation-resources/" + Version
	return atlas, nil
}
