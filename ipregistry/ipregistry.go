package ipregistry

import (
	"github.com/evalphobia/ipregistry-go/client"
	"github.com/evalphobia/ipregistry-go/config"
	"github.com/evalphobia/ipregistry-go/log"
)

// IPRegistry is service struct for Ipregistry API.
type IPRegistry struct {
	client *client.Client
	logger log.Logger
}

// New creates IPRegistry from Config data.
func New(conf config.Config) (*IPRegistry, error) {
	cli, err := conf.Client()
	if err != nil {
		return nil, err
	}

	return &IPRegistry{
		client: cli,
		logger: log.DefaultLogger,
	}, nil
}

// SetLogger sets internal API logger.
func (s *IPRegistry) SetLogger(logger log.Logger) {
	s.logger = logger
}

// SingleIP executes Single IP Lookup API.
func (s *IPRegistry) SingleIP(ipaddr string) (IPResponse, error) {
	resp := IPResponse{}
	err := s.client.CallGET("/"+ipaddr, nil, &resp)
	return resp, err
}

// BatchIP executes Batch IP Lookup API.
func (s *IPRegistry) BatchIP(ipList []string) (IPListResponse, error) {
	resp := IPListResponse{}
	err := s.client.CallPOST("/", ipList, &resp)
	return resp, err
}

// UserAgent executes UserAgent Parse API.
func (s *IPRegistry) UserAgent(userAgentList []string) (UserAgentResponse, error) {
	resp := UserAgentResponse{}
	err := s.client.CallPOST("/user_agent", userAgentList, &resp)
	return resp, err
}
