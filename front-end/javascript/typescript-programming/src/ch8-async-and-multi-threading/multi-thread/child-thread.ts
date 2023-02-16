namespace Ch8MultiThread {
  // 부모 프로세스로부터 이벤트 리스닝
  process.on('message', data => console.log('Parent process sent a data:', data) );

  // 부모 프로세스로 데이터 전달
  // @ts-ignore
  process.send({ type: 'ack', data: [1, 2, 3] });
}