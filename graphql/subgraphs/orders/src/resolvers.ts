import { orderClient } from './client.js';

// orders サブグラフのリゾルバ。BFF(gateway)の REST 経由で order サービスを呼ぶ。
export const resolvers = {
  Query: {
    order: async (_p: unknown, args: { id: string }) => orderClient.get(args.id),
    orders: async (_p: unknown, args: { userId: string }) => orderClient.listForUser(args.userId),
  },
  Mutation: {
    placeOrder: async (_p: unknown, args: { userId: string; sku: string; quantity: number }) =>
      orderClient.create(args.userId, args.sku, args.quantity),
    cancelOrder: async (_p: unknown, args: { id: string }) => orderClient.cancel(args.id),
  },
  User: {
    orders: async (user: { id: string }) => orderClient.listForUser(user.id),
  },
};
