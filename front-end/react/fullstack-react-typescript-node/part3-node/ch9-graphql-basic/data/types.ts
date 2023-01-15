export interface User {
  id: string;
  username: string;
  email?: string;
}

export interface Todo {
  id: string;
  title: string;
  description?: string;
}

export const CREATE_TODO_EVENT = 'SUBS_NEW_TODO';