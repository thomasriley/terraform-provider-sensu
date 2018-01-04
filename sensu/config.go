package sensu

import (
  "fmt"
  "log"
)

// Config defines the configuration options for the Sensu client
type Config struct {
  // The Sensu API auth token
  Token string

  // Backend host url
  Backend_Host string

  // Backend host port
  Backend_Port string

  // Backend host protocol
  Backend_Protocol string
}

const invalidCreds = `
No valid credentials found for Sensu provider.
`

// Client returns a new Sensu client
func (c *Config) Client() (*sensu.Client, error) {
  if c.Token == "" {
    return nil, fmt.Errorf(invalidCreds)
  }

  log.Printf("[INFO] Sensu client configured")

  return c, nil
}