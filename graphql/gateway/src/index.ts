import { ApolloGateway, IntrospectAndCompose } from '@apollo/gateway';
import { ApolloServer } from '@apollo/server';

// サブグラフ(users / orders / catalog)を束ねるスーパーグラフ。
const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: 'users', url: 'http://graphql-users:4001/graphql' },
      { name: 'orders', url: 'http://graphql-orders:4002/graphql' },
      { name: 'catalog', url: 'http://graphql-catalog:4003/graphql' },
    ],
  }),
});

export const server = new ApolloServer({ gateway });
