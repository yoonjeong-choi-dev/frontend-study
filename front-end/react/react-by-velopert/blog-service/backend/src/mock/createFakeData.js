const Post = require('../models/post');

module.exports = function() {
  const date = new Date().toLocaleDateString();
  const posts = [...Array(40).keys()].map((i) => ({
    title: `Post-${ i }-${ date }`,
    tags: ['Mock', 'Data', 'Test'],
    body: 'Simple Network Management Protocol (SNMP) is an Internet Standard protocol for collecting and organizing information about managed devices on IP networks and for modifying that information to change device behaviour. Devices that typically support SNMP include cable modems, routers, switches, servers, workstations, printers, and more. SNMP is widely used in network management for network monitoring. SNMP exposes management data in the form of variables on the managed systems organized in a management information base (MIB) which describe the system status and configuration. These variables can then be remotely queried (and, in some circumstances, manipulated) by managing applications.',
  }));

  return Post.insertMany(posts);
};
