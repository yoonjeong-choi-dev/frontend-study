const sanitizeHtml = require('sanitize-html');

const sanitizeOption = {
  allowedTags: [
    'h1',
    'h2',
    'b',
    'i',
    'u',
    's',
    'p',
    'ul',
    'ol',
    'li',
    'blockquote',
    'a',
    'img',
  ],
  allowedAttributes: {
    a: ['href', 'name', 'target'],
    img: ['src'],
    li: ['class'],
  },
  allowedSchemes: ['data', 'http'],
};

exports.removeHTMLAndShorten = (content) => {
  const filtered = sanitizeHtml(content, {
    allowedTags: [],
  });
  return filtered.length < 200 ? filtered : `${ filtered.slice(0, 200) }...`;
};

exports.removeBadTags = (content) => {
  console.log(content);
  console.log(sanitizeHtml(content, sanitizeOption));
  return sanitizeHtml(content, sanitizeOption);
};

