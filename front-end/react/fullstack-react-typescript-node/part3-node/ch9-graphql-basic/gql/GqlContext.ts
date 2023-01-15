import { Request, Response } from 'express';
import { PubSub } from 'graphql-subscriptions';

export interface GqlContext {
  req: Request;
  rep: Response;
  pubsub: PubSub;
}