// Package shared contains shared data between the host and plugins.
package shared

import (
	"context"

	"google.golang.org/grpc"

	bill "github.com/l0n3w0lf93/finance-share-plugin/proto/bill"
	product "github.com/l0n3w0lf93/finance-share-plugin/proto/product"

	"github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"product_grpc": &ProductGRPCPlugin{},
}

// KV is the interface that we're exposing as a plugin.
type Product interface {
	CreateProducts(string, string) error
	UpdateProducts(string, string) error
	DeleteProducts(string, string) error
	GetProducts(string, string) error
	ListProducts(string, string) error
}

type Bill interface {
	CreateBills(string, string, int64, float32, int64) error // CreateBill is the RPC method we'll be calling.
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type ProductGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	ProductImpl Product
	BillImpl    Bill
}

func (p *ProductGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	gs := &GRPCServer{
		ProductImpl: p.ProductImpl,
		BillImpl:    p.BillImpl,
	}
	product.RegisterProductsServer(s, gs)
	bill.RegisterBillsServer(s, gs)
	return nil
}

func (p *ProductGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		productclient: product.NewProductsClient(c),
		billClient:    bill.NewBillsClient(c),
	}, nil
}
