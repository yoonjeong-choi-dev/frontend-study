import { createAction, handleActions } from 'redux-actions';
import createRequestThunk, { createRequestActionTypes } from '../../lib/createRequestThunk';
import * as postsAPI from '../../lib/api/post';


const INITIALIZE_FIELD = 'write/INITIALIZE_FIELD';
const CHANGE_FIELD = 'write/CHANGE_FIELD';
const [WRITE_POST, WRITE_POST_SUCCESS, WRITE_POST_FAILURE] = createRequestActionTypes(
  'write/WRITE_POST',
);

export const initializeField = createAction(INITIALIZE_FIELD);
export const changeField = createAction(CHANGE_FIELD, ({ key, value }) => ({
  key, value,
}));
export const writePost = createRequestThunk(WRITE_POST, postsAPI.write);

const initialState = {
  title: '',
  body: '',
  tags: [],
  post: null,
  postError: null,
};

const write = handleActions({
  [CHANGE_FIELD]: (state, { payload: { key, value } }) => ({
    ...state,
    [key]: value,
  }),
  [INITIALIZE_FIELD]: (state) => initialState,
  [WRITE_POST]: (state) => ({
    ...state,
    post: null,
    postError: null,
  }),
  [WRITE_POST_SUCCESS]: (state, { payload: post }) => ({
    ...state,
    postError: null,
    post,
  }),
  [WRITE_POST_FAILURE]: (state, { payload: error }) => ({
    ...state,
    postError: error,
  }),
}, initialState);

export default write;