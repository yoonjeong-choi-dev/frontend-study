import React from 'react';
import styled from 'styled-components';

const SCContainer = styled.div`
  display: flex;

  & + & {
    margin-top: 3rem;
  }
`;

const SCThumbnail = styled.div`
  margin-right: 1rem;

  img {
    display: block;
    width: 160px;
    height: 160px;
    object-fit: cover;
  }
`;

const SCContent = styled.div`
  h2 {
    margin: 0;

    a {
      color: black;
      text-decoration: none;
      outline: none;
      cursor: pointer;

      &:hover {
        text-decoration: underline;
      }
    }
  }

  p {
    margin: 0.5rem 0 0 0;
    line-height: 1.5;
    white-space: normal;
  }
`;

const NewsItem = ({article}) => {
  const {title, description, url, urlToImage} = article;
  return (
    <SCContainer>
      {urlToImage && (
        <SCThumbnail>
          <a href={url} target="_blank" rel="noopener noreferrer">
            <img src={urlToImage} alt="thumbnail"/>
          </a>
        </SCThumbnail>
      )}
      <SCContent>
        <h2>
          <a href={url} target="_blank" rel="noopener noreferrer">
            {title}
          </a>
        </h2>
        <p>{description}</p>
      </SCContent>
    </SCContainer>
  )
};

export default NewsItem;