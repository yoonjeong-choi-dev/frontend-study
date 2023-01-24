import * as api from '../lib/api';
import { handleActions } from 'redux-actions';

const GET_POST = 'jsonHolder/GET_POST';
const GET_POST_SUCCESS = 'jsonHolder/GET_POST_SUCCESS';
const GET_POST_FAILURE = 'jsonHolder/GET_POST_FAILURE';

const GET_USERS = 'jsonHolder/GET_USERS';
const GET_USERS_SUCCESS = 'jsonHolder/GET_USERS_SUCCESS';
const GET_USERS_FAILURE = 'jsonHolder/GET_USERS_FAILURE';

// redux-thunk 미들웨어가 처리할 thunk 함수들 정의
// 액션 생성 함수와 비슷한 역할
export const getPost = id => async dispatch => {
  // 요청을 시작함을 알림
  dispatch({type: GET_POST});

  // API Call
  try {
    const response = await api.getPost(id);
    dispatch({
      type: GET_POST_SUCCESS,
      payload: response.data,
    });
  } catch (e) {
    dispatch({
      type: GET_POST_FAILURE,
      payload: e,
      error: true,
    });
    throw e;
  }
};

export const getUsers = () => async dispatch => {
  dispatch({type: GET_USERS});

  try {
    const response = await api.getUsers();
    dispatch({
      type: GET_USERS_SUCCESS,
      payload: response.data,
    });
  } catch (e) {
    dispatch({
      type: GET_USERS_FAILURE,
      payload: e,
      error: true,
    });
    throw e;
  }
};

// 초기 상태 및 리듀서 정의
const initialState = {
  loading: {
    GET_POST: false,
    GET_USERS: false,
  },
  post: null,
  users: null,
};

const jsonHolder = handleActions({
  [GET_POST]: state => ({
    ...state,
    loading: {
      ...state.loading,
      GET_POST: true,
    }
  }),
  [GET_POST_SUCCESS]: (state, action) => ({
    ...state,
    loading: {
      ...state.loading,
      GET_POST: false,
    },
    post: action.payload
  }),
  [GET_POST_FAILURE]: (state, action) => ({
    ...state,
    loading: {
      ...state.loading,
      GET_POST: false,
    },
  }),
  [GET_USERS]: state => ({
    ...state,
    loading: {
      ...state.loading,
      GET_USERS: true,
    }
  }),
  [GET_USERS_SUCCESS]: (state, action) => ({
    ...state,
    loading: {
      ...state.loading,
      GET_USERS: false,
    },
    users: action.payload
  }),
  [GET_USERS_FAILURE]: (state, action) => ({
    ...state,
    loading: {
      ...state.loading,
      GET_USERS: false,
    },
  }),
}, initialState);

export default jsonHolder;