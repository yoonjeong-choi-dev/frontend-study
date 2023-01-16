namespace Ch5BuilderPattern {

  type Method = 'get' | 'post' | null;

  class RequestBuilder {
    private url: string | null = null;
    private method: Method = null;
    private data: object | null = null;

    setUrl(url: string): this {
      this.url = url;
      return this;
    }

    setMethod(method: Method): this {
      this.method = method;
      return this;
    }

    setData(data: object): this {
      this.data = data;
      return this;
    }

    send() {
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
  const req = new RequestBuilder().setUrl('http://localhost:8080').setMethod('get').setData(myData);
  console.log(req.send());
}