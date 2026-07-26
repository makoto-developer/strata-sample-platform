/** GraphQL ドキュメントを組み立てるタグ関数。 */
export function gql(strings: TemplateStringsArray, ...values: unknown[]): string {
  return strings.reduce((acc, s, i) => acc + s + (i < values.length ? String(values[i]) : ''), '');
}

/** スーパーグラフに操作を投げる。 */
export async function request(document: string, variables: Record<string, unknown> = {}): Promise<unknown> {
  const res = await fetch('/graphql', {
    method: 'POST',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ query: document, variables }),
  });
  return res.json();
}
