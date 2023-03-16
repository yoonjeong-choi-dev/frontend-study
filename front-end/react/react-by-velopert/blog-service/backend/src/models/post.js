const mongoose = require('mongoose');

const { Schema, model } = mongoose;

// Schema
const PostSchema = new Schema({
  title: String,
  body: String,
  tags: [String],
  user: {
    _id: mongoose.Types.ObjectId,
    name: String,
  },
}, { timestamps: true });

// Model
const Post = model('Post', PostSchema);

module.exports = Post;