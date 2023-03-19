import client from './client';

const BASE_URI = '/api/auth';

export const login = ({ name, password }) => {
  return client.post(`${ BASE_URI }/login`, { name, password });
};

export const register = ({ name, password }) => {
  return client.post(`${ BASE_URI }/register`, { name, password });
};

export const check = () => client.get(`${ BASE_URI }/check`);

export const logout = () => client.post(`${BASE_URI}/logout`);