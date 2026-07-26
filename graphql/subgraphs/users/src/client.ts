// user サービスの REST エンドポイントを呼ぶ(gRPC スタブ未生成のため REST 経由)。
const BASE = process.env.USER_HTTP_URL ?? 'http://user:8081';

export const userClient = {
  async getUser(id: string): Promise<unknown> {
    const res = await fetch(`${BASE}/api/v1/users/${id}`);
    return res.json();
  },
  async listUsers(limit: number): Promise<unknown> {
    const res = await fetch(`${BASE}/api/v1/users?limit=${limit}`);
    return res.json();
  },
  async updateProfile(id: string, name: string): Promise<unknown> {
    const res = await fetch(`${BASE}/api/v1/users/${id}/profile`, {
      method: 'PATCH',
      body: JSON.stringify({ name }),
    });
    return res.json();
  },
  async deactivateUser(id: string): Promise<boolean> {
    const res = await fetch(`${BASE}/api/v1/users/${id}`, { method: 'DELETE' });
    return res.ok;
  },
};
