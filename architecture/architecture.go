package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	// "github.com/blushft/go-diagrams/nodes/alibabacloud"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/aws"

	// "github.com/blushft/go-diagrams/nodes/azure"
	// "github.com/blushft/go-diagrams/nodes/elastic"
	// "github.com/blushft/go-diagrams/nodes/firebase"
	"github.com/blushft/go-diagrams/nodes/gcp"
	// "github.com/blushft/go-diagrams/nodes/generic"
	// "github.com/blushft/go-diagrams/nodes/k8s"
	// "github.com/blushft/go-diagrams/nodes/oci"
	// "github.com/blushft/go-diagrams/nodes/openstack"
	// "github.com/blushft/go-diagrams/nodes/outscale"
	"github.com/blushft/go-diagrams/nodes/programming"
	// "github.com/blushft/go-diagrams/nodes/saas"
)

// Font Struct
type Font struct {
	Name  	string
	Size  	int
	Color 	string
}

// GroupOptions Struct
type GroupOptions struct {
	Label           string
	LabelJustify    string
	Direction       string
	PenColor        string
	BackgroundColor string
	Shape           string
	Style           string
	Font            Font
	Attributes      map[string]string
}

func main() {
	// Configure a new diagram
	d, err := diagram.New(
		diagram.Filename("architecture"),
		diagram.Label("App"),
		diagram.Direction("TB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Frontend
	ng := programming.Framework.Angular(diagram.NodeLabel("Chapter Website"))
	dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))

	// Backend APIs
	halfDome := programming.Language.Nodejs(diagram.NodeLabel("Half Dome"))
	tuolumne := programming.Framework.Laravel(diagram.NodeLabel("Tuolumne"))
	tenaya := programming.Language.Go(diagram.NodeLabel("Tenaya"))
	mariposa := gcp.Devtools.MavenAppEnginePlugin(diagram.NodeLabel("Mariposa"))
	
	// DevOps
	apiGateway := aws.Network.ApiGateway(diagram.NodeLabel("Api Gateway"))

	// defaultOpts := diagram.DefaultGroupOptions()
	// apiOpts := diagram.GroupOption(defaultOpts)

	api := diagram.NewGroup("API").Label("API Layer")
	api.NewGroup("Half Dome").Label("Service 1").Add(
		halfDome,
		apps.Database.Postgresql(diagram.NodeLabel("Replica")), 
		apps.Container.Docker(diagram.NodeLabel("Docker")),
	)
	api.NewGroup("Tuolumne").Label("Service 2").Add(
		tuolumne,
		apps.Database.Postgresql(diagram.NodeLabel("Replica")), 
		apps.Container.Docker(diagram.NodeLabel("Docker")),
	)
	api.NewGroup("Tenaya").Label("Service 3").Add(
		tenaya,
		apps.Database.Postgresql(diagram.NodeLabel("Replica")), 
		apps.Container.Docker(diagram.NodeLabel("Docker")),
	)
	api.NewGroup("Mariposa").Label("Service 4").Add(
		mariposa,
		apps.Database.Postgresql(diagram.NodeLabel("Replica")), 
		apps.Container.Docker(diagram.NodeLabel("Docker")),
	)
	api.ConnectAllFrom(apiGateway.ID(), diagram.Forward())

	d.Connect(dns, ng, diagram.Forward())
	d.Connect(ng, apiGateway, diagram.Forward())
	d.Connect(apiGateway, halfDome, diagram.Forward())
	d.Connect(apiGateway, tuolumne, diagram.Forward())
	d.Connect(apiGateway, tenaya, diagram.Forward())
	d.Connect(apiGateway, mariposa, diagram.Forward())
	d.Group(api)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
