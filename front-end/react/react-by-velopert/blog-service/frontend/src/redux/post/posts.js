import createRequestThunk, { createRequestActionTypes } from '../../lib/createRequestThunk';
import * as postsAPI from '../../lib/api/post';
import { handleActions } from 'redux-actions';

export const POST_LIST_ACTION_TYPE = 'posts/LIST_POSTS';
const [LIST_POSTS, LIST_POSTS_SUCCESS, LIST_POSTS_FAILURE] = createRequestActionTypes(POST_LIST_ACTION_TYPE);

export const listPosts = createRequestThunk(LIST_POSTS, postsAPI.list);

const initialState = {
  lastPage: 1,
  posts: null,
  error: null,
};

const posts = handleActions({
  [LIST_POSTS_SUCCESS]: (state, { payload: posts, meta: response }) => ({
    ...state,
    posts,
    lastPage: parseInt(response.headers['last-page'], 10),
    error: null,
  }),
  [LIST_POSTS_FAILURE]: (state, { payload: error }) => ({
    ...state,
    post: null,
    error,
  }),
}, initialState);

export default posts;