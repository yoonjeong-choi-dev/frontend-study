namespace Ch4Exercise3 {
  type Reservation = string;

  type Reserve = {
    (from: Date, to: Date, destination: string): Reservation;
    (from: Date, destination: string): Reservation;
    (destination: string): Reservation;
  }

  const reserve: Reserve = (
    fromOrDest: Date | string,
    toOrDest?: Date | string,
    destination?: string,
  ) => {
    let from: Date, to: Date|null = null;
    let dest: string;
    if(typeof fromOrDest === 'string') {
      // third def
      dest = fromOrDest;
      from = new Date();
    } else if(typeof toOrDest === 'string' ) {
      // second def
      from = fromOrDest;
      dest = toOrDest;
    } else if(toOrDest !== undefined && destination != undefined) {
      // first def
      from = fromOrDest;
      to = toOrDest;
      dest = destination;
    } else {
      return 'Invalid Args';
    }

    return `[${dest}] ${from.toLocaleTimeString()} ~ ${to ? to.toLocaleTimeString() : 'UNKNOWN'}`;
  }

  console.log(reserve("Japan"));
  console.log(reserve(new Date(), new Date(), 'Korea'));
  console.log(reserve(new Date(), 'USA'));
}
