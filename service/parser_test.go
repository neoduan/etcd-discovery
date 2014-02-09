package service

import (
	"github.com/coreos/go-etcd/etcd"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	sampleNode = &etcd.Node{
		Key: "/services/test/example.org",
		Value: `
		{
			"Name": "example.org",
			"User": "user",
			"Password": "password",
			"Port": "port"
		}
		`,
	}
	sampleNodes = etcd.Nodes{*sampleNode, *sampleNode}
)

var (
	sampleResult = &Host{
		Name:     "example.org",
		User:     "user",
		Password: "password",
		Port:     "port",
	}
)

func TestBuildHostsFromNodes(t *testing.T) {
	hosts := buildHostsFromNodes(sampleNodes)
	Convey("Given a sample response with 2 nodes, we got 2 hosts", t, func() {
		So(len(hosts), ShouldEqual, 2)
		So(hosts[0], ShouldResemble, sampleResult)
		So(hosts[1], ShouldResemble, sampleResult)
	})
}

func TestBuildHostFromNode(t *testing.T) {
	host := buildHostFromNode(sampleNode)
	Convey("Given a sample response, we got a filled Host", t, func() {
		So(host, ShouldResemble, sampleResult)
	})
}