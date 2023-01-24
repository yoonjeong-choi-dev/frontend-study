import axios from 'axios';

const GET_USERS_PENDING = 'users/GET_USERS_PENDING';
const GET_USERS_SUCCESS = 'users/GET_USERS_SUCCESS';
const GET_USERS_FAILURE = 'users/GET_USERS_FAILURE';

const GET_USER_PENDING = 'users/GET_USER';
const GET_USER_SUCCESS = 'users/GET_USER_SUCCESS';
const GET_USER_FAILURE = 'users/GET_USER_FAILURE';

// Action Creator
const getUsersPending = () => ({type: GET_USERS_PENDING});
const getUsersSuccess = payload => ({type: GET_USERS_SUCCESS, payload});
const getUsersFailure = payload => ({
  type: GET_USERS_FAILURE,
  error: true,
  payload,
});

const getUserPending = () => ({type: GET_USER_PENDING});
const getUserSuccess = payload => ({type: GET_USER_SUCCESS, payload});
const getUserFailure = payload => ({
  type: GET_USER_FAILURE,
  error: true,
  payload,
});

export const getUsers = () => async dispatch => {
  try {
    dispatch(getUsersPending());
    const res = await axios.get('https://jsonplaceholder.typicode.com/users');
    dispatch(getUsersSuccess(res));
  } catch (e) {
    dispatch(getUsersFailure(e));
    throw e;
  }
};

export const getUser = (id) => async dispatch => {
  try {
    dispatch(getUserPending());
    const res = await axios.get(`https://jsonplaceholder.typicode.com/users/${id}`);
    dispatch(getUserSuccess(res));
  } catch (e) {
    dispatch(getUserFailure(e));
    throw e;
  }
};

// 초기 상태 및 리듀서 정의
const initialState = {
  users: null,
  user: null,
  loading: {
    users: false,
    user: false
  },
  error: {
    users: null,
    user: null
  }
};

function users(state = initialState, action) {
  switch (action.type) {
    case GET_USERS_PENDING:
      return {
        ...state,
        loading: { ...state.loading, users: true },
        error: { ...state.error, users: null }
      };
    case GET_USERS_SUCCESS:
      return {
        ...state,
        loading: { ...state.loading, users: false },
        users: action.payload.data
      };
    case GET_USERS_FAILURE:
      return {
        ...state,
        loading: { ...state.loading, users: false },
        error: { ...state.error, users: action.payload }
      };
    case GET_USER_PENDING:
      return {
        ...state,
        loading: { ...state.loading, user: true },
        error: { ...state.error, user: null }
      };
    case GET_USER_SUCCESS:
      return {
        ...state,
        loading: { ...state.loading, user: false },
        user: action.payload.data
      };
    case GET_USER_FAILURE:
      return {
        ...state,
        loading: { ...state.loading, user: false },
        error: { ...state.error, user: action.payload }
      };
    default:
      return state;
  }
}

export default users;