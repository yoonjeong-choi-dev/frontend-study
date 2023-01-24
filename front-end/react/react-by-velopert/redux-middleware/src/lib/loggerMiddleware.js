// redux-logger 로 대체 가능
const loggerMiddleware = store => next => action => {
  console.group(action && action.type);
  console.log('prev state', store.getState());
  console.log('action', action);

  // 다음 미들웨어 or 리듀서에 전달
  next(action);

  // 전달된 결과
  console.log('cur state', store.getState());
  console.groupEnd();
}

export default loggerMiddleware;