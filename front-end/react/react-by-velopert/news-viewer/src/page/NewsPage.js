import React from 'react';
import {useParams} from 'react-router-dom';

import CategoryHeader from '../component/CategoryHeader';
import NewsList from '../component/NewsList';

const NewsPage = () => {
  const params = useParams();
  const category = params.category ?? 'all';

  return (
    <>
      <CategoryHeader category={category}/>
      <NewsList category={category}/>
    </>
  );
}

export default NewsPage;