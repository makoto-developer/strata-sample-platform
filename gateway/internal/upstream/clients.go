package upstream

import (
	catalogv1 "github.com/makoto-developer/strata-sample-platform/proto/catalog/v1"
	orderv1 "github.com/makoto-developer/strata-sample-platform/proto/order/v1"
	userv1 "github.com/makoto-developer/strata-sample-platform/proto/user/v1"
)

// Clients は BFF が束ねる下流 gRPC クライアント。
type Clients struct {
	Users    userv1.UserServiceClient
	Orders   orderv1.OrderServiceClient
	Catalog  catalogv1.CatalogServiceClient
}

func Dial() *Clients { return &Clients{} }
