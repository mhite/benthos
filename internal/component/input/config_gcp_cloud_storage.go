package input

// GCPGCSPubSubConfig contains configuration for hooking up the GCS input with a Pub/Sub subscription.
type GCPGCSPubSubConfig struct {
	Project      string `json:"project" yaml:"project"`
	Subscription string `json:"subscription" yaml:"subscription"`
}

// NewGCPGCSPubSubConfig creates a new GCPGCSPubSubConfig with default values.
func NewGCPGCSPubSubConfig() GCPGCSPubSubConfig {
	return GCPGCSPubSubConfig{
		Project:      "",
		Subscription: "",
	}
}

// GCPCloudStorageConfig contains configuration fields for the Google Cloud
// Storage input type.
type GCPCloudStorageConfig struct {
	Bucket        string             `json:"bucket" yaml:"bucket"`
	Prefix        string             `json:"prefix" yaml:"prefix"`
	Codec         string             `json:"codec" yaml:"codec"`
	DeleteObjects bool               `json:"delete_objects" yaml:"delete_objects"`
	PubSub        GCPGCSPubSubConfig `json:"pubsub" yaml:"pubsub"`
}

// NewGCPCloudStorageConfig creates a new GCPCloudStorageConfig with default
// values.
func NewGCPCloudStorageConfig() GCPCloudStorageConfig {
	return GCPCloudStorageConfig{
		Codec: "all-bytes",
	}
}
