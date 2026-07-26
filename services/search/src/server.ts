import Fastify from 'fastify';

import { indexProduct, searchProducts } from './handlers.js';

// search サービスの REST 表面(Fastify)。catalog から呼ばれる。
export function buildServer() {
  const app = Fastify({ logger: true });

  app.get('/search/products', async (req) => {
    const q = (req.query as { q?: string }).q ?? '';
    return searchProducts(q);
  });

  app.post('/search/index', async (req) => {
    return indexProduct(String(req.body));
  });

  app.get('/health', async () => ({ status: 'ok' }));

  return app;
}

buildServer().listen({ port: 8087 });
