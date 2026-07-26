// 外部の検索エンジン(Elasticsearch 互換 SaaS)への HTTP クライアント。
const BASE = process.env.SEARCH_ENGINE_URL ?? 'https://search-engine.example.com';

export const engine = {
  async query(q: string): Promise<unknown> {
    const res = await fetch(`${BASE}/indexes/products/search?q=${encodeURIComponent(q)}`);
    return res.json();
  },

  async index(id: string): Promise<unknown> {
    const res = await fetch(`${BASE}/indexes/products/documents`, {
      method: 'POST',
      body: JSON.stringify({ id }),
    });
    return res.json();
  },
};
