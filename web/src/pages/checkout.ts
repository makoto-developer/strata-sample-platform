import { gql, request } from '../lib/gql.js';

// チェックアウト: 注文を作り、進捗は BFF の REST で追う
const PLACE_ORDER = gql`
  mutation PlaceOrder($userId: ID!, $sku: String!, $quantity: Int!) {
    placeOrder(userId: $userId, sku: $sku, quantity: $quantity) {
      id
      status
    }
  }
`;

export async function placeOrder(userId: string, sku: string, quantity: number): Promise<unknown> {
  return request(PLACE_ORDER, { userId, sku, quantity });
}

/** 注文詳細は BFF の REST から取る。 */
export async function fetchOrder(id: string): Promise<unknown> {
  const res = await fetch(`/api/orders/${id}`);
  return res.json();
}

/** キャンセルも BFF 経由。 */
export async function cancelOrder(id: string): Promise<unknown> {
  const res = await fetch(`/api/orders/${id}/cancel`, { method: 'POST' });
  return res.json();
}
