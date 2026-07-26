import { userClient } from './client.js';

// users サブグラフのリゾルバ。gRPC の UserService に委譲する。
export const resolvers = {
  Query: {
    me: async (_p: unknown, _a: unknown, ctx: { userId: string }) => userClient.getUser(ctx.userId),
    user: async (_p: unknown, args: { id: string }) => userClient.getUser(args.id),
    users: async (_p: unknown, args: { limit?: number }) => userClient.listUsers(args.limit ?? 20),
  },
  Mutation: {
    updateProfile: async (_p: unknown, args: { id: string; name: string }) =>
      userClient.updateProfile(args.id, args.name),
    deactivateUser: async (_p: unknown, args: { id: string }) => userClient.deactivateUser(args.id),
  },
  User: {
    __resolveReference: async (ref: { id: string }) => userClient.getUser(ref.id),
  },
};
