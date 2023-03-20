import createRequestThunk, { createRequestActionTypes } from '../../lib/createRequestThunk';
import { createAction, handleActions } from 'redux-actions';
import * as postsAPI from '../../lib/api/post';

export const POST_ACTION_TYPE = 'post/READ_POST';

const UNLOAD_POST = 'post/UNLOAD_POST';
const [READ_POST, READ_POST_SUCCESS, READ_POST_FAILURE] = createRequestActionTypes(POST_ACTION_TYPE);

export const unloadPost = createAction(UNLOAD_POST);
export const readPost = createRequestThunk(READ_POST, postsAPI.read);

const initialState = {
  post: null,
  error: null,
};

const post = handleActions({
  [UNLOAD_POST]: () => initialState,
  [READ_POST_SUCCESS]: (state, { payload: post }) => ({
    ...state,
    post,
    error: null,
  }),
  [READ_POST_FAILURE]: (state, { payload: error }) => ({
    ...state,
    post: null,
    error,
  }),
}, initialState);

export default post;