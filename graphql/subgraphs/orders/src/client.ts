// BFF(gateway)の REST エンドポイントを叩く。
const BASE = process.env.GATEWAY_URL ?? 'http://gateway:8080';

export const orderClient = {
  async get(id: string): Promise<unknown> {
    const res = await fetch(`${BASE}/api/orders/${id}`);
    return res.json();
  },
  async listForUser(userId: string): Promise<unknown> {
    const res = await fetch(`${BASE}/api/users/${userId}/orders`);
    return res.json();
  },
  async create(userId: string, sku: string, quantity: number): Promise<unknown> {
    const res = await fetch(`${BASE}/api/orders`, {
      method: 'POST',
      body: JSON.stringify({ userId, sku, quantity }),
    });
    return res.json();
  },
  async cancel(id: string): Promise<unknown> {
    const res = await fetch(`${BASE}/api/orders/${id}/cancel`, { method: 'POST' });
    return res.json();
  },
};
