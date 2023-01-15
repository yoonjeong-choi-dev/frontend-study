import { IResolvers } from '@graphql-tools/utils';
import { v4 } from 'uuid';

import { GqlContext } from './GqlContext';
import { CREATE_TODO_EVENT, Todo, User } from '../data/types';
import { todos, users } from '../data/local-storage';


const resolvers: IResolvers = {
  Query: {
    getUser: async (
      obj: any,
      args: {
        id: string;
      },
      ctx: GqlContext,
      info: any
    ): Promise<User | undefined> => {
      return users.find(user => user.id === args.id);
    },
    getUsers: async (
      obj: any,
      args: null,
      ctx: GqlContext,
      info: any
    ): Promise<Array<User>> => {
      return users;
    },
    getTodos: async (
      obj: any,
      args: null,
      ctx: GqlContext,
      info: any
    ): Promise<Array<Todo>> => {
      return todos;
    }
  },
  Mutation: {
    addUser: async (
      parent: any,
      args: {
        username: string;
        email?: string,
      },
      ctx: GqlContext,
      info: any
    ): Promise<User> => {
      users.push({
        id: v4(),
        ...args,
      });
      return users[users.length - 1];
    },
    addTodo: async (
      parent: any,
      args: {
        title: string;
        description?: string,
      },
      { pubsub }: GqlContext,
      info: any
    ): Promise<Todo> => {
      const todo = { id: v4(), ...args };
      todos.push(todo);
      pubsub.publish(CREATE_TODO_EVENT, { todo });
      return todos[todos.length - 1];
    }
  },
  Subscription: {
    createTodo: {
      subscribe: (parent, args: null, { pubsub }: GqlContext) =>
        pubsub.asyncIterator(CREATE_TODO_EVENT)
    }
  }
};

export default resolvers;