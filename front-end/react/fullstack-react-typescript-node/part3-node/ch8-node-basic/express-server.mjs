import express from 'express';
import bodyParser from 'body-parser';

const port = 7166;
const app = express();
app.use(bodyParser.json());

// Define Router
const router = express.Router();
router.get('/', (req, res, next) => {
  res.send('Hello Express!');
})
router.post('/', (req, res, next) => {
  console.log(req.body)
  res.send(`Post Test : Your name is ${req.body.name}`);
})

// Define api
const users = [
  {
    id:1,
    username: 'yj',
  },
  {
    id:2,
    username: 'yoonjeong',
  },
  {
    id:3,
    username: 'yoonjeong choi',
  },
];
router.get('/api/v1/users', (req, res, next) => {
  res.send(users);
})
router.get('/api/v1/user', (req, res, next) => {
  const id = Number(req.query.id);
  console.log(id);
  const user = users.find((item) => item.id === id);
  if(user) {
    res.send(`User ${id} name is '${user.username}'`);
  } else {
    res.send(`There is no user with id ${id}`);
  }
});
router.post('/api/v1/user', (req, res, next) => {
  const id = req.body.id;
  const username = req.body.username;
  users.push({id, username});
  res.send(users);
})

// entry logger
app.use((req, res, next) => {
  console.log(`First Middleware : ${req.url}`);
  next();
});

app.use(router);

// Error Handle

app.listen({port}, () => {
  console.log(`Express Node Server has loaded : ${port}`);
});
