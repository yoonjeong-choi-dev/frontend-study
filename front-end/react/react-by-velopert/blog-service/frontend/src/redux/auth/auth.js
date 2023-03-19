import produce from 'immer';
import { createAction, handleActions } from 'redux-actions';

import createRequestThunk, { createRequestActionTypes } from '../../lib/createRequestThunk';
import * as authAPI from '../../lib/api/auth';

const CHANGE_FIELD = 'auth/CHANGE_FIELD';
const INITIALIZE_FORM = 'auth/INITIALIZE_FORM';
const [REGISTER, REGISTER_SUCCESS, REGISTER_FAILURE] = createRequestActionTypes(
  'auth/REGISTER',
);
const [LOGIN, LOGIN_SUCCESS, LOGIN_FAILURE] = createRequestActionTypes(
  'auth/LOGIN',
);

// Action Creator for Form Change
export const changeField = createAction(
  CHANGE_FIELD,
  ({ form, key, value }) => ({
    form, // form name containing the field
    key,  // input name
    value,// input value to change
  }),
);

export const initializeForm = createAction(INITIALIZE_FORM, form => form);

// Action Creator for Auth API
export const register = createRequestThunk(REGISTER, authAPI.register);
export const login = createRequestThunk(LOGIN, authAPI.login);


const initialState = {
  register: {
    name: '',
    password: '',
    passwordConfirm: '',
  },
  login: {
    name: '',
    password: '',
  },
  auth: null,
  authError: null,
};


const auth = handleActions({
  [CHANGE_FIELD]: (state, { payload: { form, key, value } }) =>
    produce(state, (draft) => {
      draft[form][key] = value;
    }),
  [INITIALIZE_FORM]: (state, { payload: { form } }) => ({
    ...state,
    [form]: initialState[form],
  }),
  [REGISTER_SUCCESS]: (state, { payload: auth }) => ({
    ...state,
    authError: null,
    auth,
  }),
  [REGISTER_FAILURE]: (state, { payload: error }) => ({
    ...state,
    authError: error,
  }),
  [LOGIN_SUCCESS]: (state, { payload: auth }) => ({
    ...state,
    authError: null,
    auth,
  }),
  [LOGIN_FAILURE]: (state, { payload: error }) => ({
    ...state,
    authError: error,
  }),
}, initialState);

export default auth;