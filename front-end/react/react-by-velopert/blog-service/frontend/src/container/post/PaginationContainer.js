import { useParams, useSearchParams } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { POST_LIST_ACTION_TYPE } from '../../redux/post/posts';
import Pagination from '../../component/post/list/Pagination';

const PaginationContainer = () => {
  const { username } = useParams();
  const [searchParams] = useSearchParams();

  const tag = searchParams.get('tag');
  const page = parseInt(searchParams.get('page'), 10) || 1;

  const { posts, lastPage, loading } = useSelector((state) => ({
    posts: state.posts.posts,
    lastPage: state.posts.lastPage,
    loading: state.loading[POST_LIST_ACTION_TYPE],
  }));

  // const currentPage = useMemo(()=> {
  //   if(page<1 || loading) return 1;
  //   return Math.min(page, lastPage);
  // },[page, lastPage,loading]);

  if (loading || !posts) return null;

  return (
    <Pagination
      page={ page }
      lastPage={ lastPage }
      tag={ tag }
      username={ username }
    />
  );
};

export default PaginationContainer;