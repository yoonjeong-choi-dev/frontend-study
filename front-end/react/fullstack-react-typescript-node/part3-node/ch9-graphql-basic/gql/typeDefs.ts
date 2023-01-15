import { gql } from 'apollo-server-express';

const typeDefs = gql`
  type User {
    id: ID!
    username: String!
    email: String
  }
  
  type Todo {
    id: ID!
    title: String!
    description: String
  }
  
  type Query {
    getUser(id: ID): User
    getUsers: [User!]
    getTodos: [Todo!]
  }
  
  type Mutation {
      addUser(username: String!, email: String): User
      addTodo(title: String!, description: String): Todo
  }
  
  type Subscription {
      createTodo: Todo
  }
`;

export default typeDefs;