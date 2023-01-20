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
    const result = Object.assign({}, performance);
    result.play = getPlay(performance);
    result.amount = getAmountFor(result);
    result.volumeCredit = getVolumeCredit(result);
    return result;
  }

  // 데이터 정제에 사용하였던 중첩 함수들 : refinePerformance 에서 사용
  function getPlay(performance) {
    return plays[performance.playID];
  }

  function getAmountFor(performance) {
    let result = 0;
    switch (performance.play.type) {
      case 'tragedy':
        result = 40000;

        // bonus amount
        if (performance.audience > 30) {
          result += 1000 * (performance.audience - 30);
        }
        break;
      case  'comedy':
        result = 30000;

        // bonus amount
        result += 300 * performance.audience;
        if (performance.audience > 20) {
          result += 10000 + 500 * (performance.audience - 20);
        }
        break;
      default:
        throw new Error(`Unknown Type: ${performance.play.type}`);
    }
    return result;
  }

  function getVolumeCredit(performance) {
    let result = 0;
    result += Math.max(performance.audience - 30, 0);
    if (performance.play.type === 'comedy') {
      result += Math.floor(performance.audience / 5);
    }

    return result;
  }

// 반복문 쪼개기 : 공연료 및 청구서 결과 계산하는 부분 => 함수 추출
  function getTotalAmount(data) {
    return data.performances.reduce((acc, performance) => acc + performance.amount, 0);
    // let result = 0;
    // for (let performance of data.performances) {
    //   result += performance.amount;
    // }
    // return result;
  }

// 반복문 쪼개기 : 보너스 포인트 계산하는 부분 => 함수 추출
  function getTotalVolumeCredits(data) {
    return data.performances.reduce((acc, performance) => acc + performance.volumeCredit, 0);

    // let result = 0;
    // for (let performance of data.performances) {
    //   // calculate bonus point
    //   result += performance.volumeCredit;
    // }
    // return result;
  }
}

