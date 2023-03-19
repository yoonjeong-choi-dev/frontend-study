import client from './client';

const BASE_URI = '/api/posts';

export const write = ({ title, body, tags }) =>
  client.post(`${ BASE_URI }`, { title, body, tags });