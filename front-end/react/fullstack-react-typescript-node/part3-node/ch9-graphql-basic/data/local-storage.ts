import { v4 } from 'uuid';

import { User, Todo } from './types';

export const users: Array<User> = [
  {
    id: v4(),
    username: 'Yoonjeong Choi',
    email: 'yjchoi7166@gmail.com',
  },
  {
    id: v4(),
    username: 'Naver Yoonjeong',
    email: 'yjchoi7166@naver.com',
  }
];

export const todos: Array<Todo> = [
  {
    id: v4(),
    title: 'First Todo',
    description: 'Work out',
  },
  {
    id: v4(),
    title: 'Second Todo',
    description: 'Master CSS',
  },
  {
    id: v4(),
    title: 'Third Todo',
    description: 'React Programming',
  },
  {
    id: v4(),
    title: 'Last Todo',
  },
];