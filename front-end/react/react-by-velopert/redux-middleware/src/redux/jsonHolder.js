import * as api from '../lib/api';
import { handleActions } from 'redux-actions';
import createRequestThunk from './lib/createRequestThunk';

export const GET_POST = 'jsonHolder/GET_POST';
const GET_POST_SUCCESS = 'jsonHolder/GET_POST_SUCCESS';

export const GET_USERS = 'jsonHolder/GET_USERS';
const GET_USERS_SUCCESS = 'jsonHolder/GET_USERS_SUCCESS';

// redux-thunk 미들웨어가 처리할 thunk 함수들 정의
// redux-thunk 미들웨어에 적용할 액션 생성 함수
export const getPost = createRequestThunk(GET_POST, api.getPost);
export const getUsers = createRequestThunk(GET_USERS, api.getUsers);

// 초기 상태 및 리듀서 정의
const initialState = {
  post: null,
  users: null,
};

const jsonHolder = handleActions({
  [GET_POST_SUCCESS]: (state, action) => ({
    ...state,
    post: action.payload
  }),
  [GET_USERS_SUCCESS]: (state, action) => ({
    ...state,
    users: action.payload
  }),
}, initialState);

export default jsonHolder;