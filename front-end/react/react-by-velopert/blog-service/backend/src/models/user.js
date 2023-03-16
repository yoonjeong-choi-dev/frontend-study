const mongoose = require('mongoose');
const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');

const { Schema, model } = mongoose;

const UserSchema = new Schema({
  name: String,
  password: String,
}, { timestamps: true });

UserSchema.methods.setPassword = async function(password) {
  const hashed = await bcrypt.hash(password, 10);
  this.password = hashed;
};

UserSchema.methods.checkPassword = async function(password) {
  const ret = await bcrypt.compare(password, this.password);
  return ret;
};

UserSchema.methods.serialize = function() {
  const data = this.toJSON();
  delete data.password;
  return data;
};

UserSchema.methods.generateToken = function() {
  const token = jwt.sign(
    {
      _id: this.id,
      name: this.name,
    },
    process.env.JWT_SECRET_KEY,
    {
      expiresIn: '7d',
    });
  return token;
};

UserSchema.statics.findByName = function(name) {
  return this.findOne({ name });
};

const User = model('User', UserSchema);
module.exports = User;