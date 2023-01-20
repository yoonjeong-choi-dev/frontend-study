// 공연에 대한 전체 금액 및 적립 포인트를 계산하는 추상 클래스
class PerformanceCalculator {
  constructor(performance, play) {
    this.performance = performance;
    this.play = play;
  }

  // getAmountFor(performance)
  get amount() {
    throw new Error('This property is calculated in sub classes');
  }

  // getVolumeCredit(performance)
  get volumeCredit() {
    return Math.max(this.performance.audience - 30, 0);
  }
}

class TragedyCalculator extends PerformanceCalculator {
  // getAmountFor(performance)
  get amount() {
    let result = 40000;
    if (this.performance.audience > 30) {
      result += 1000 * (this.performance.audience - 30);
    }

    return result;
  }
}

class ComedyCalculator extends PerformanceCalculator {
  // getAmountFor(performance)
  get amount() {
    let result = 30000;

    // bonus amount
    result += 300 * this.performance.audience;
    if (this.performance.audience > 20) {
      result += 10000 + 500 * (this.performance.audience - 20);
    }

    return result;
  }

  // getVolumeCredit(performance)
  get volumeCredit() {
    return super.volumeCredit + Math.floor(this.performance.audience / 5);;
  }
}

// Factory Method Pattern
function createPerformanceCalculator(performance, play) {
  switch (play.type) {
    case 'tragedy':
      return new TragedyCalculator(performance, play);
    case  'comedy':
      return new ComedyCalculator(performance, play);
    default:
      throw new Error(`Unknown Type: ${play.type}`);
  }
}

// 출력에 필요하도록 데이터 정제
export default function createStatementData(invoice, plays) {
  const result = {
    customer: invoice.customer,
    performances: invoice.performances,
  };
  result.customer = invoice.customer;
  result.performances = invoice.performances.map(refinePerformance);
  result.totalAmount = getTotalAmount(result);
  result.totalVolumeCredits = getTotalVolumeCredits(result);
  return result;


  function refinePerformance(performance) {
    const calculator = createPerformanceCalculator(performance, getPlay(performance));
    const result = Object.assign({}, performance);
    result.play = calculator.play;
    result.amount = calculator.amount;
    result.volumeCredit = calculator.volumeCredit;
    return result;
  }

  // 데이터 정제에 사용하였던 중첩 함수들 : refinePerformance 에서 사용
  function getPlay(performance) {
    return plays[performance.playID];
  }

  function getTotalAmount(data) {
    return data.performances.reduce((acc, performance) => acc + performance.amount, 0);
  }

  function getTotalVolumeCredits(data) {
    return data.performances.reduce((acc, performance) => acc + performance.volumeCredit, 0);
  }
}

