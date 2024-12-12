package vacancy

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/vacancy-pkg/models"

	proto "github.com/garden-raccoon/vacancy-pkg/protocols/vacancy"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

type IVacancyAPI interface {
	CreateVacancy(*models.Vacancy) error
	GetVacancies() ([]*models.Vacancy, error)
	HealthCheck() error

	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	mu      sync.Mutex
	*grpc.ClientConn
	proto.VacancyServiceClient
	grpc_health_v1.HealthClient
}

// New create new Battles Api instance
func New(addr string) (IVacancyAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create Battles Api:  %w", err)
	}

	api.VacancyServiceClient = proto.NewVacancyServiceClient(api.ClientConn)
	return api, nil
}

func (api *Api) CreateVacancy(s *models.Vacancy) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	_, err = api.VacancyServiceClient.CreateVacancy(ctx, s.Proto())
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

	api.ClientConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	return
}
func (api *Api) GetVacancies() ([]*models.Vacancy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.Vacancies
	empty := new(proto.VacancyDbEmpty)
	resp, err := api.VacancyServiceClient.GetVacancies(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	vacs := models.VacanciesFromProto(resp)

	return vacs, nil
}
func (api *Api) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	api.mu.Lock()
	defer api.mu.Unlock()

	resp, err := api.HealthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: "userapi"})
	if err != nil {
		return fmt.Errorf("healthcheck error: %w", err)
	}

	if resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return fmt.Errorf("node is %s", errors.New("service is unhealthy"))
	}
	//api.status = service.StatusHealthy
	return nil
}
