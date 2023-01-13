// import React, { useCallback, useState } from 'react';
// import NewsList from './component/NewsList';
// import CategoryHeader from './component/CategoryHeader';
import React from 'react';
import { Route, Routes } from 'react-router-dom';
import NewsPage from './page/NewsPage';


function App() {
  // const [category, setCategory] = useState('all');
  // const onSelect = useCallback((category) => {
  //   setCategory(category)
  // }, []);

  return (
    // <>
    //   <CategoryHeader category={category} onSelect={onSelect}/>
    //   <NewsList category={category}/>
    // </>
    <Routes>
      <Route path={'/'} element={<NewsPage/>}/>
      <Route path={'/:category'} element={<NewsPage/>}/>
    </Routes>
  );
}

export default App;
