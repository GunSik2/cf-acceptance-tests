package assets

type Assets struct {
	Dora                     string
	HelloWorld               string
	Node                     string
	NodeWithProcfile         string
	NodeWithWebsocket        string
	Java                     string
	Golang                   string
	Python                   string
	LoggregatorLoadGenerator string
	ServiceBroker            string
	AsyncServiceBroker       string
	Php                      string
	SecurityGroupBuildpack   string
	Fuse                     string
	RubySimple               string
	DoraZip                  string
}

func NewAssets() Assets {
	return Assets{
		Dora:             "../assets/dora",
		HelloWorld:       "../assets/hello-world",
		Node:             "../assets/node",
		NodeWithProcfile: "../assets/node-with-procfile",
		Java:             "../assets/java",
		Golang:           "../assets/golang",
		Python:           "../assets/python",
		LoggregatorLoadGenerator: "../assets/loggregator-load-generator",
		ServiceBroker:            "../assets/service_broker",
		AsyncServiceBroker:       "../assets/service_broker",
		Php:                      "../assets/php",
		SecurityGroupBuildpack: "../assets/security_group_buildpack.zip",
		Fuse:       "../assets/fuse-mount",
		RubySimple: "../assets/ruby_simple",
		DoraZip:    "../assets/dora.zip",
	}
}
