package ovs

import (
	"errors"
	"fmt"
)

var (
	ErrNoAddr   = errors.New("no address")
	ErrNotFound = errors.New("not found")
)

// Client is the ovn client
type Client struct {
	OvnNbAddress                  string
	OvnTimeout                    int
	OvnSbAddress                  string
	OVNIcNBAddress                string
	ClusterRouter                 string
	ClusterTcpLoadBalancer        string
	ClusterUdpLoadBalancer        string
	ClusterTcpSessionLoadBalancer string
	ClusterUdpSessionLoadBalancer string
	NodeSwitch                    string
	NodeSwitchCIDR                string
}

const (
	OvnNbCtl    = "ovn-nbctl"
	OvnSbCtl    = "ovn-sbctl"
	OVNIcNbCtl  = "ovn-ic-nbctl"
	MayExist    = "--may-exist"
	IfExists    = "--if-exists"
	Policy      = "--policy"
	PolicyDstIP = "dst-ip"
	PolicySrcIP = "src-ip"
)

// NewClient init an ovn client
func NewClient(ovnNbHost string, ovnNbPort int, ovnNbTimeout int, ovnSbHost string, ovnSbPort int, clusterRouter, clusterTcpLoadBalancer, clusterUdpLoadBalancer, clusterTcpSessionLoadBalancer, clusterUdpSessionLoadBalancer, nodeSwitch, nodeSwitchCIDR string) *Client {
	return &Client{
		OvnNbAddress:                  fmt.Sprintf("tcp:%s:%d", ovnNbHost, ovnNbPort),
		OvnSbAddress:                  fmt.Sprintf("tcp:%s:%d", ovnSbHost, ovnSbPort),
		OvnTimeout:                    ovnNbTimeout,
		ClusterRouter:                 clusterRouter,
		ClusterTcpLoadBalancer:        clusterTcpLoadBalancer,
		ClusterUdpLoadBalancer:        clusterUdpLoadBalancer,
		ClusterTcpSessionLoadBalancer: clusterTcpSessionLoadBalancer,
		ClusterUdpSessionLoadBalancer: clusterUdpSessionLoadBalancer,
		NodeSwitch:                    nodeSwitch,
		NodeSwitchCIDR:                nodeSwitchCIDR,
	}
}
