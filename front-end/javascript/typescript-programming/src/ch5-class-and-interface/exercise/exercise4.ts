namespace Ch5Exercise4 {
  type Method = 'get' | 'post';

  interface Request {
    url: string;
    method: Method;
    data?: object;
  }

  class RequestBuilder {
    url?: string;
    method?: Method;
    data?: object;

    setUrl(url: string): this & Pick<Request, 'url'> {
      return Object.assign(this, { url });
    }

    setMethod(method: Method): this & Pick<Request, 'method'> {
      return Object.assign(this, { method });
    }

    setData(data: object): this & Pick<Request, 'data'> {
      return Object.assign(this, {data});
    }

    // 현재 객체가 Request 타입보다 구체적인 경우에만 호출 가능
    // i.e 적어도 url 및 method 가 설정된 상태에서만 호출
    send(this: Request) {
      return `[${ this.method?.toUpperCase() }] ${ this.url } ->\n ${ JSON.stringify(this.data, null, 2) }`;
    }
  }

  const myData = {
    name: 'YJ',
    age: 31,
    detail: {
      email: 'yjchoi7166@gmail.com',
      company: 'moloco',
    }
  }

  // TS2684: The 'this' context of type 'RequestBuilder' is not assignable to method's 'this' of type 'Request'.
  //new RequestBuilder().send();
  //new RequestBuilder().setUrl('test').send();
  //new RequestBuilder().setMethod('get').send();
  console.log(new RequestBuilder().setUrl('http://localhost:8080').setMethod('get').send());
  console.log(new RequestBuilder().setData(myData).setUrl('http://localhost:8080').setMethod('get').send());
}