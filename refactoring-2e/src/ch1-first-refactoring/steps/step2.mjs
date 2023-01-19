export default function (invoice, plays) {
  // 반복문 쪼개기 : 청구서 결과 계산하는 부분
  let result = `청구 내역(customer: ${invoice.customer})\n`;
  for (let performance of invoice.performances) {
    // Update result
    result += `\t${getPlay(performance).name}: ${usd(getAmountFor(performance))} (${performance.audience} seats)\n`
  }

  result += `Total Amount: ${usd(getTotalAmount())}\n`;
  result += `Reward Points: ${getTotalVolumeCredits()} points\n`;
  return result;

  /*
  * 아래는 기능 분리 및 로컬 변수 제거를 위한 함수들
  * */
  function usd(number) {
    return new Intl.NumberFormat('es-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2,
    }).format(number / 100);
  }

  function getPlay(performance) {
    return plays[performance.playID];
  }

  function getAmountFor(performance) {
    let result = 0;
    switch (getPlay(performance).type) {
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
        throw new Error(`Unknown Type: ${getPlay(performance).type}`);
    }
    return result;
  }

  function getVolumeCredit(performance) {
    let result = 0;
    result += Math.max(performance.audience - 30, 0);
    if (getPlay(performance).type === 'comedy') {
      result += Math.floor(performance.audience / 5);
    }

    return result;
  }

  // 반복문 쪼개기 : 공연료 및 청구서 결과 계산하는 부분 => 함수 추출
  function getTotalAmount() {
    let result = 0;
    for (let performance of invoice.performances) {
      result += getAmountFor(performance);
    }
    return result;
  }

  // 반복문 쪼개기 : 보너스 포인트 계산하는 부분 => 함수 추출
  function getTotalVolumeCredits() {
    let result = 0;
    for (let performance of invoice.performances) {
      // calculate bonus point
      result += getVolumeCredit(performance);
    }
    return result;
  }
}