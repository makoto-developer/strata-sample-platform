import { gql, request } from '../lib/gql.js';

const PRODUCTS_QUERY = gql`
  query Products($keyword: String) {
    products(keyword: $keyword) {
      id
      name
      price
    }
  }
`;

export async function searchProducts(keyword: string): Promise<unknown> {
  return request(PRODUCTS_QUERY, { keyword });
}

/** 商品詳細は BFF の REST を使う(GraphQL を経由しない導線)。 */
export async function fetchProduct(id: string): Promise<unknown> {
  const res = await fetch(`/api/products/${id}`);
  return res.json();
}
