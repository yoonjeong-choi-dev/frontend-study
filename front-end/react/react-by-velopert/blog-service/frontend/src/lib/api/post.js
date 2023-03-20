import client from './client';

const BASE_URI = '/api/posts';

export const write = ({ title, body, tags }) =>
  client.post(BASE_URI, { title, body, tags });

export const read = (id) => client.get(`${ BASE_URI }/${ id }`);

export const list = ({ page, username, tag }) => {
  return client.get(BASE_URI, {
    params: {
      page,
      username,
      tag,
    },
  });
};

export const update = ({ id, title, body, tags }) => {
  return client.patch(`${ BASE_URI }/${ id }`, {
    title, body, tags,
  });
};

export const remove = (id) => client.delete(`${ BASE_URI }/${ id }`);