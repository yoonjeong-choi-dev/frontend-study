import { combineReducers } from 'redux';
import loading from './loading';
import auth from './auth/auth';
import user from './auth/user';
import write from './post/write';
import post from './post/post';
import posts from './post/posts';

const rootReducer = combineReducers({
  loading,
  auth,
  user,
  write,
  post,
  posts,
});

export default rootReducer;