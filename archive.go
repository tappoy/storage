// This package provides an interface for archiving to cloud services.
//
// Supported protocols:
//   - OpenStack  https://docs.openstack.org/2024.1/api/
//   - AWS        https://docs.aws.amazon.com/AmazonS3/latest/API/Welcome.html
//
// Tested services:
//   - ConoHa     https://doc.conoha.jp/api-vps3/
//   - Sakura     https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html
//
// May be supported:
//   - CloudFlare https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
//
// Others:
//   - Google     https://cloud.google.com/storage/docs/json_api/v1
package archive

import (
	"io"
	"time"
)

type Object struct {
	Name         string    `json:"name"`
	Hash         string    `json:"hash"`
	Bytes        int64     `json:"bytes"`
	LastModified time.Time `json:"last_modified"`
}

type Client interface {
	// List retrieves a object list in the container.
	List(prefix string) ([]Object, error)

	// Put creates an object.
	Put(object string, body io.Reader) error

	// Delete deletes an object.
	Delete(object string) error

	// Head retrieves an object metadata.
	Head(object string) (Object, error)
}