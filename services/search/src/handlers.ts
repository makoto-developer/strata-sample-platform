import { engine } from './engine.js';

/** 全文検索を実行する。 */
export async function searchProducts(query: string): Promise<unknown> {
  return engine.query(query);
}

/** 商品を検索インデックスに登録する。 */
export async function indexProduct(id: string): Promise<unknown> {
  return engine.index(id);
}
