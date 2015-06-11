package aws

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

// DefaultChainCredentials is a Credentials which will find the first available
// credentials Value from the list of Providers.
//
// This should be used in the default case. Once the type of credentials are
// known switching to the specific Credentials will be more efficient.
var DefaultChainCredentials = credentials.NewChainCredentials(
	[]credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
		&credentials.EC2RoleProvider{ExpiryWindow: 5 * time.Minute},
	})

// The default number of retries for a service. The value of -1 indicates that
// the service specific retry default will be used.
const DefaultRetries = -1

// DefaultConfig is the default all service configuration will be based off of.
// By default, all clients use this structure for initialization options unless
// a custom configuration object is passed in.
//
// You may modify this global structure to change all default configuration
// in the SDK. Note that configuration options are copied by value, so any
// modifications must happen before constructing a client.
var DefaultConfig = &Config{
	Credentials:             DefaultChainCredentials,
	Endpoint:                String(""),
	Region:                  String(os.Getenv("AWS_REGION")),
	DisableSSL:              Bool(false),
	HTTPClient:              http.DefaultClient,
	LogHTTPBody:             Bool(false),
	LogLevel:                Int(0),
	Logger:                  os.Stdout,
	MaxRetries:              Int(DefaultRetries),
	DisableParamValidation:  Bool(false),
	DisableComputeChecksums: Bool(false),
	S3ForcePathStyle:        Bool(false),
}

// A Config provides service configuration for service clients. By default,
// all clients will use the {DefaultConfig} structure.
type Config struct {
	// The credentials object to use when signing requests. Defaults to
	// {DefaultChainCredentials}.
	Credentials *credentials.Credentials

	// An optional endpoint URL (hostname only or fully qualified URI)
	// that overrides the default generated endpoint for a client. Set this
	// to `""` to use the default generated endpoint.
	//
	// @note You must still provide a `Region` value when specifying an
	//   endpoint for a client.
	Endpoint *string

	// The region to send requests to. This parameter is required and must
	// be configured globally or on a per-client basis unless otherwise
	// noted. A full list of regions is found in the "Regions and Endpoints"
	// document.
	//
	// @see http://docs.aws.amazon.com/general/latest/gr/rande.html
	//   AWS Regions and Endpoints
	Region *string

	// Set this to `true` to disable SSL when sending requests. Defaults
	// to `false`.
	DisableSSL *bool

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	HTTPClient *http.Client

	// Set this to `true` to also log the body of the HTTP requests made by the
	// client.
	//
	// @note `LogLevel` must be set to a non-zero value in order to activate
	//   body logging.
	LogHTTPBody *bool

	// An integer value representing the logging level. The default log level
	// is zero (0), which represents no logging. Set to a non-zero value to
	// perform logging.
	LogLevel *int

	// The logger writer interface to write logging messages to. Defaults to
	// standard out.
	Logger io.Writer

	// The maximum number of times that a request will be retried for failures.
	// Defaults to -1, which defers the max retry setting to the service specific
	// configuration.
	MaxRetries *int

	// Disables semantic parameter validation, which validates input for missing
	// required fields and/or other semantic request input errors.
	DisableParamValidation *bool

	// Disables the computation of request and response checksums, e.g.,
	// CRC32 checksums in Amazon DynamoDB.
	DisableComputeChecksums *bool

	// Set this to `true` to force the request to use path-style addressing,
	// i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will
	// use virtual hosted bucket addressing when possible
	// (`http://BUCKET.s3.amazonaws.com/KEY`).
	//
	// @note This configuration option is specific to the Amazon S3 service.
	// @see http://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
	//   Amazon S3: Virtual Hosting of Buckets
	S3ForcePathStyle *bool
}

// Copy will return a shallow copy of the Config object.
func (c Config) Copy() Config {
	dst := c
	return dst
}

// Merge merges the newcfg attribute values into this Config. Each attribute
// will be merged into this config if the newcfg attribute's value is non-zero.
// Due to this, newcfg attributes with zero values cannot be merged in. For
// example bool attributes cannot be cleared using Merge, and must be explicitly
// set on the Config structure.
func (c Config) Merge(newcfg *Config) *Config {
	cfg := c
	if newcfg == nil {
		return &cfg
	}

	if newcfg.Credentials != nil {
		cfg.Credentials = newcfg.Credentials
	}

	if newcfg.Endpoint != nil {
		cfg.Endpoint = newcfg.Endpoint
	}

	if newcfg.Region != nil {
		cfg.Region = newcfg.Region
	}

	if newcfg.DisableSSL != nil {
		cfg.DisableSSL = newcfg.DisableSSL
	}

	if newcfg.HTTPClient != nil {
		cfg.HTTPClient = newcfg.HTTPClient
	}

	if newcfg.LogHTTPBody != nil {
		cfg.LogHTTPBody = newcfg.LogHTTPBody
	}

	if newcfg.LogLevel != nil {
		cfg.LogLevel = newcfg.LogLevel
	}

	if newcfg.Logger != nil {
		cfg.Logger = newcfg.Logger
	}

	if newcfg.MaxRetries != nil {
		cfg.MaxRetries = newcfg.MaxRetries
	}

	if newcfg.DisableParamValidation != nil {
		cfg.DisableParamValidation = newcfg.DisableParamValidation
	}

	if newcfg.DisableComputeChecksums != nil {
		cfg.DisableComputeChecksums = newcfg.DisableComputeChecksums
	}

	if newcfg.S3ForcePathStyle != nil {
		cfg.S3ForcePathStyle = newcfg.S3ForcePathStyle
	}

	return &cfg
}
