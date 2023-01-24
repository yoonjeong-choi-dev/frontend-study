const JsonHolder = ({title, loadingPost, loadingUsers, post, users}) => {
  return (
    <div>
      <h2>{title}</h2>
      <section>
        <h3>Post - 1</h3>
        {loadingPost && <h2>Loading...</h2>}
        {!loadingPost && post && (
          <div>
            <h3>Title</h3>
            <div>{post.title}</div>
            <h4>Content</h4>
            <div>{post.body}</div>
          </div>
        )}
      </section>
      <section>
        <h3>Users list</h3>
        {loadingUsers && <h2>Loading...</h2>}
        {!loadingUsers && users && (
          <ul>
            {users.map(user => (
              <li key={user.id}>
                {user.username} ({user.email})
              </li>
            ))}
          </ul>
        )}
      </section>
    </div>
  );
}

export default JsonHolder;