package shared

import (
	bill "github.com/l0n3w0lf93/finance-share-plugin/proto/bill"
	product "github.com/l0n3w0lf93/finance-share-plugin/proto/product"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of product that talks over RPC.
type GRPCClient struct {
	productclient product.ProductsClient
	billClient    bill.BillsClient
}

func (m *GRPCClient) CreateProducts(key string, code string) error {
	_, err := m.productclient.CreateProducts(context.Background(), &product.CreateProductsRequest{
		Name: key,
		Code: code,
	})
	return err
}
func (m *GRPCClient) DeleteProducts(key string, name string) error {
	_, err := m.productclient.DeleteProducts(context.Background(), &product.DeleteProductsRequest{})
	return err
}
func (m *GRPCClient) UpdateProducts(key string, name string) error {
	_, err := m.productclient.UpdateProducts(context.Background(), &product.UpdateProductsRequest{})
	return err
}
func (m *GRPCClient) GetProducts(key string, name string) error {
	_, err := m.productclient.GetProducts(context.Background(), &product.GetProductsRequest{})
	return err
}
func (m *GRPCClient) ListProducts(key string, name string) error {
	_, err := m.productclient.ListProducts(context.Background(), &product.ListProductsRequest{})
	return err
}
func (m *GRPCClient) CreateBills(userId string, productId string, price float32, amount int64, buyAt int64) error {
	_, err := m.billClient.CreateBills(context.Background(), &bill.CreateBillsRequest{ProductId: productId, UserId: userId, Amount: amount, Price: price, BuyAt: buyAt})
	return err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	ProductImpl Product
	BillImpl    Bill
	bill.UnimplementedBillsServer
	product.UnimplementedProductsServer
}

func (m *GRPCServer) CreateProducts(ctx context.Context, req *product.CreateProductsRequest) (*product.CreateProductsReply, error) {
	return &product.CreateProductsReply{}, m.ProductImpl.CreateProducts(req.Code, req.Name)
}
func (m *GRPCServer) UpdateProducts(ctx context.Context, req *product.UpdateProductsRequest) (*product.UpdateProductsReply, error) {
	return &product.UpdateProductsReply{}, m.ProductImpl.UpdateProducts("key", "value")
}
func (m *GRPCServer) DeleteProducts(ctx context.Context, req *product.DeleteProductsRequest) (*product.DeleteProductsReply, error) {
	return &product.DeleteProductsReply{}, m.ProductImpl.DeleteProducts("key", "value")
}
func (m *GRPCServer) GetProducts(ctx context.Context, req *product.GetProductsRequest) (*product.GetProductsReply, error) {
	return &product.GetProductsReply{}, m.ProductImpl.GetProducts("key", "value")
}
func (m *GRPCServer) ListProducts(ctx context.Context, req *product.ListProductsRequest) (*product.ListProductsReply, error) {
	return &product.ListProductsReply{}, m.ProductImpl.ListProducts("key", "value")
}
func (m *GRPCServer) CreateBills(ctx context.Context, req *bill.CreateBillsRequest) (*bill.CreateBillsReply, error) {
	return &bill.CreateBillsReply{}, m.BillImpl.CreateBills(req.ProductId, req.UserId, req.Amount, req.Price, req.BuyAt)
}
func (m *GRPCServer) mustEmbedUnimplementedProductsServer() {}
func (m *GRPCServer) mustEmbedUnimplementedBillsServer()    {}
