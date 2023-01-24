import { createAction, handleActions } from 'redux-actions';
import { takeLatest } from 'redux-saga/effects';

import * as api from '../lib/api';
import createRequestSaga from './lib/createRequestSaga';

export const GET_POST = 'jsonHolderForSaga/GET_POST';
const GET_POST_SUCCESS = 'jsonHolderForSaga/GET_POST_SUCCESS';

export const GET_USERS = 'jsonHolderForSaga/GET_USERS';
const GET_USERS_SUCCESS = 'jsonHolderForSaga/GET_USERS_SUCCESS';


// Redux Sage 에서 사용하기 위한 액션 생성 함수
export const getPost = createAction(GET_POST, id => id);
export const getUsers = createAction(GET_USERS);

export const getPostSaga = createRequestSaga(GET_POST, api.getPost);
export const getUsersSaga = createRequestSaga(GET_USERS, api.getUsers);

// 초기 상태 및 리듀서 정의
const initialState = {
  post: null,
  users: null,
};

export function* jsonHolderSaga() {
  yield takeLatest(GET_POST, getPostSaga);
  yield takeLatest(GET_USERS, getUsersSaga);
}

const jsonHolderForSaga = handleActions({
  [GET_POST_SUCCESS]: (state, action) => ({
    ...state,
    post: action.payload
  }),
  [GET_USERS_SUCCESS]: (state, action) => ({
    ...state,
    users: action.payload
  }),
}, initialState);

export default jsonHolderForSaga;