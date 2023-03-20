import { createAction, handleActions } from 'redux-actions';
import createRequestThunk, { createRequestActionTypes } from '../../lib/createRequestThunk';
import * as postsAPI from '../../lib/api/post';

const INITIALIZE_FIELD = 'write/INITIALIZE_FIELD';
const CHANGE_FIELD = 'write/CHANGE_FIELD';
const [WRITE_POST, WRITE_POST_SUCCESS, WRITE_POST_FAILURE] = createRequestActionTypes(
  'write/WRITE_POST',
);
const SET_ORIGINAL_POST = 'write/SET_ORIGINAL_POST';
const [UPDATE_POST, UPDATE_POST_SUCCESS, UPDATE_POST_FAILURE] = createRequestActionTypes('write/UPDATE_POST'); // 포스트 수정

export const initializeField = createAction(INITIALIZE_FIELD);
export const changeField = createAction(CHANGE_FIELD, ({ key, value }) => ({
  key, value,
}));
export const writePost = createRequestThunk(WRITE_POST, postsAPI.write);
export const setOriginalPost = createAction(SET_ORIGINAL_POST, post => post);
export const updatePost = createRequestThunk(UPDATE_POST, postsAPI.update);

const initialState = {
  title: '',
  body: '',
  tags: [],
  post: null,
  postError: null,
  originalPostId: null,
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
  [SET_ORIGINAL_POST]: (state, { payload: post }) => ({
    ...state,
    originalPostId: post._id,
    title: post.title,
    body: post.body,
    tags: post.tags,
  }),
  [UPDATE_POST_SUCCESS]: (state, { payload: post }) => ({
    ...state,
    postError: null,
    post,
  }),
  [UPDATE_POST_FAILURE]: (state, { payload: error }) => ({
    ...state,
    postError: error,
  }),
}, initialState);

export default write;