package resolver

import (
	"context"

	catalogv1 "github.com/makoto-developer/strata-sample-platform/proto/catalog/v1"
)

// Resolver は gqlgen のルートリゾルバ。gRPC の CatalogService に委譲する。
type Resolver struct {
	Catalog catalogv1.CatalogServiceClient
}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (r *Resolver) Query() *queryResolver { return &queryResolver{r} }

func (r *Resolver) Mutation() *mutationResolver { return &mutationResolver{r} }

// Product は 1 件の商品を返す。
func (r *queryResolver) Product(ctx context.Context, id string) (*catalogv1.GetProductReply, error) {
	return r.Catalog.GetProduct(ctx, &catalogv1.GetProductRequest{Id: id})
}

// Products は商品を検索して返す。
func (r *queryResolver) Products(ctx context.Context, keyword *string) (*catalogv1.SearchProductsReply, error) {
	q := ""
	if keyword != nil {
		q = *keyword
	}
	return r.Catalog.SearchProducts(ctx, &catalogv1.SearchProductsRequest{Id: q})
}

// UpsertProduct は商品を登録・更新する。
func (r *mutationResolver) UpsertProduct(ctx context.Context, id string, name string, price int) (*catalogv1.UpsertProductReply, error) {
	return r.Catalog.UpsertProduct(ctx, &catalogv1.UpsertProductRequest{Id: id})
}
