package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/alibabacloud"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/blushft/go-diagrams/nodes/azure"
	"github.com/blushft/go-diagrams/nodes/elastic"
	"github.com/blushft/go-diagrams/nodes/firebase"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/blushft/go-diagrams/nodes/generic"
	"github.com/blushft/go-diagrams/nodes/k8s"
	"github.com/blushft/go-diagrams/nodes/oci"
	"github.com/blushft/go-diagrams/nodes/openstack"
	"github.com/blushft/go-diagrams/nodes/outscale"
	"github.com/blushft/go-diagrams/nodes/programming"
	"github.com/blushft/go-diagrams/nodes/saas"
)

func getDNS(name string) *diagram.Node {
	switch name {
	case "alibabacloud":
		return alibabacloud.Web.Dns(diagram.NodeLabel("DNS"))
	case "gcp":
		return gcp.Network.Dns(diagram.NodeLabel("DNS"))
	case "azure":
		return azure.Network.DnsZones(diagram.NodeLabel("DNS"))
	case "oci":
		return oci.Connectivity.Dns(diagram.NodeLabel("DNS"))
	}

	return gcp.Network.Dns(diagram.NodeLabel("DNS"))
}

func main() {
	// Configure a new diagram
	d, err := diagram.New(
		diagram.Filename("architecture"), 
		diagram.Label("App"), 
		diagram.Direction("LR"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ng := programming.Framework.Angular(diagram.NodeLabel("Chapter Website"))
	dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))
	aws_general := aws.General.General(diagram.NodeLabel("general"))
	kibana := elastic.Elasticsearch.Kibana(diagram.NodeLabel("Kibana"))
	firebase := firebase.Base.Firebase(diagram.NodeLabel("Firebase"))
	blank := generic.Blank.Blank(diagram.NodeLabel("Blank"))
	helm := k8s.Ecosystem.Helm(diagram.NodeLabel("Helm"))
	ansible := openstack.Deployment.Ansible(diagram.NodeLabel("Ansible"))
	net := outscale.Network.Net(diagram.NodeLabel("Net"))
	slack := saas.Chat.Slack(diagram.NodeLabel("Slack"))
	docker :=  apps.Container.Docker(diagram.NodeLabel("Docker"))

	dc := diagram.NewGroup("Test").Add(
		ng, dns, aws_general, kibana, firebase, blank, helm, ansible, net, slack, docker,
	)

	d.Group(dc)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
