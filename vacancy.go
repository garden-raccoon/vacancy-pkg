package orders

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/orders-admin-package/models"
	proto "github.com/garden-raccoon/orders-admin-package/protocols/orders-admin-package"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

// Debug on/off
var Debug = true

// IOrderAdminPkgAPI is
type IOrderAdminPkgAPI interface {
	GetOrders() ([]*models.OrderAdmin, error)
	OrderByName(name string) (*models.OrderAdmin, error)
	CreateOrders(s []*models.OrderAdmin) error
	DeleteOrder(uuid uuid.UUID) error
	CreateMeals(s []*models.MealsDb) error
	GetMeals() ([]*models.MealsDb, error)

	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	proto.OrderAdminServiceClient
}

// New create new Battles Api instance
func New(addr string) (IOrderAdminPkgAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create Battles Api:  %w", err)
	}

	api.OrderAdminServiceClient = proto.NewOrderAdminServiceClient(api.ClientConn)
	return api, nil
}
func (api *Api) DeleteOrder(uuid uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	req := &proto.OrderDeleteReq{
		Uuid: uuid.Bytes(),
	}
	_, err := api.OrderAdminServiceClient.DeleteOrder(ctx, req)
	if err != nil {
		return fmt.Errorf("DeleteOrder api request: %w", err)
	}
	return nil
}

func (api *Api) GetOrders() ([]*models.OrderAdmin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.OrdersAdmin
	empty := new(proto.OrderAdminEmpty)
	resp, err := api.OrderAdminServiceClient.GetOrders(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	orders := models.OrdersFromProto(resp)

	return orders, nil
}
func (api *Api) GetMeals() ([]*models.MealsDb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.MealsDb
	empty := new(proto.MealsDbEmpty)
	resp, err := api.OrderAdminServiceClient.GetMeals(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	meals := models.MealsFromProto(resp)

	return meals, nil
}
func (api *Api) CreateOrders(s []*models.OrderAdmin) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	orders := models.OrdersToProto(s)

	_, err = api.OrderAdminServiceClient.CreateOrders(ctx, orders)
	if err != nil {
		return fmt.Errorf("create orders api request: %w", err)
	}
	return nil
}

func (api *Api) CreateMeals(s []*models.MealsDb) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	meals := models.MealsToProto(s)

	_, err = api.OrderAdminServiceClient.CreateMeals(ctx, meals)
	if err != nil {
		return fmt.Errorf("create meals api request: %w", err)
	}
	return nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	return
}

// OrderByName is
func (api *Api) OrderByName(name string) (*models.OrderAdmin, error) {
	getter := &proto.OrderAdminGetter{
		Getter: &proto.OrderAdminGetter_Name{
			Name: name,
		},
	}
	return api.getOrder(getter)
}

// getOrder is
func (api *Api) getOrder(getter *proto.OrderAdminGetter) (*models.OrderAdmin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	resp, err := api.OrderAdminServiceClient.OrderByName(ctx, getter)
	if err != nil {
		return nil, fmt.Errorf("get battle api request: %w", err)
	}

	return models.OrderAdminFromProto(resp), nil
}
