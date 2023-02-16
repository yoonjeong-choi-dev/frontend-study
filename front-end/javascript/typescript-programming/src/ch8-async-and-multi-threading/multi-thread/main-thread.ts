import {fork} from 'child_process';

namespace Ch8MultiThread {
  const child = fork('./child-thread.ts');
  child.on('message', data => console.log('Child process sent a data:', data));

  child.send({type: 'syn', data: [1]});

}