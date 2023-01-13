//import React, { useEffect, useState } from 'react';
import React from 'react';
import styled from 'styled-components';
import NewsItem from './NewsItem';
import axios from 'axios';
import usePromise from '../hooks/usePromise';

const SCListContainer = styled.div`
  box-sizing: border-box;
  width: 768px;

  padding-bottom: 3rem;
  margin: 0 auto;
  margin-top: 2rem;

  @media screen and (max-width: 768px) {
    width: 100%;
    padding: 0 1rem;
  }
`;

const NewsList = ({category}) => {
  // const [articles, setArticles] = useState(null);
  // const [loading, setLoading] = useState(false);
  // useEffect(() => {
  //   console.log(category);
  //   const fetchData = async () => {
  //     setLoading(true);
  //     try {
  //       const params = {
  //         apiKey: process.env.REACT_APP_NEW_API_KEY,
  //         country: 'kr',
  //       };
  //       if(category !== 'all') {
  //         params['category'] = category;
  //       }
  //
  //       const rep = await axios.get(
  //         'https://newsapi.org/v2/top-headlines', {params}
  //       );
  //
  //       setArticles(rep.data.articles);
  //     } catch (e) {
  //       console.error(e);
  //     }
  //     setLoading(false);
  //   }
  //   fetchData();
  // }, [category]);

  const [loading, response, error] = usePromise(() => {
    const params = {
      apiKey: process.env.REACT_APP_NEW_API_KEY,
      country: 'kr',
    };
    if (category !== 'all') {
      params['category'] = category;
    }

    return axios.get(
      'https://newsapi.org/v2/top-headlines', {params}
    );
  }, [category])

  if (loading) {
    return <SCListContainer>Loading...</SCListContainer>
  }

  if (!response) {
    return null;
  }

  if (error) {
    return <SCListContainer>Error Occur!</SCListContainer>
  }

  const {articles} = response.data;
  return (
    <SCListContainer>
      {articles && articles.map(article => (
        <NewsItem key={article.url} article={article}/>
      ))}
    </SCListContainer>
  );
};

export default NewsList;