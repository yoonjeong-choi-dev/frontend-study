import http from 'http';

import express from 'express';
import { ApolloServer } from 'apollo-server-express';
import { ApolloServerPluginDrainHttpServer } from 'apollo-server-core';
import { makeExecutableSchema } from '@graphql-tools/schema';

import { WebSocketServer } from 'ws';
import { useServer } from 'graphql-ws/lib/use/ws';
import { PubSub } from 'graphql-subscriptions';

import typeDefs from './gql/typeDefs';
import resolvers from './gql/resolvers';

const port = 7166;

const app = express();
const pubsub = new PubSub();
const httpServer = http.createServer(app);

const schema = makeExecutableSchema({
  typeDefs, resolvers
});

const webSocketServer = new WebSocketServer({
  server: httpServer,
  path: '/graphql',
});

const serverCleanup = useServer(
  {
    schema,
    context: () => {
      return { pubsub };
    },
  },
  webSocketServer
);

const apolloServer = new ApolloServer({
  schema,
  context: ({ req, res }: any) => {
    return { req, res, pubsub };
  },
  csrfPrevention: true,
  cache: "bounded",
  plugins: [
    ApolloServerPluginDrainHttpServer({ httpServer }),
    {
      async serverWillStart() {
        return {
          async drainServer() {
            await serverCleanup.dispose();
          },
        };
      },
    },
  ],
})


apolloServer.start().then(() => {
  apolloServer.applyMiddleware({ app, cors: true });
  httpServer.listen({ port }, () => {
    console.log(`GraphQL server is ready with port ${ port }`);
  })
})