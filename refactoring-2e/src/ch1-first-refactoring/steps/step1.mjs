export default function (invoice, plays) {
  let totalAmount = 0;
  let volumeCredits = 0;
  let result = `청구 내역(customer: ${invoice.customer})\n`;

  const format = new Intl.NumberFormat('es-US', {
    style: 'currency',
    currency: 'USD',
    minimumFractionDigits: 2,
  }).format;

  for (let performance of invoice.performances) {
    const play = plays[performance.playID];
    let currentAmount = getAmountFor(performance, play);

    // calculate bonus point
    volumeCredits += Math.max(performance.audience - 30, 0);
    if (play.type === 'comedy') {
      volumeCredits += Math.floor(performance.audience / 5);
    }

    // Update result
    result += `\t${play.name}: ${format(currentAmount / 100)} (${performance.audience} seats)\n`
    totalAmount += currentAmount;
  }

  result += `Total Amount: ${format(totalAmount / 100)}\n`;
  result += `Reward Points: ${volumeCredits} points\n`;
  return result;

  function getAmountFor(performance, play) {
    let result = 0;
    switch (play.type) {
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
        throw new Error(`Unknown Type: ${play.type}`);
    }
    return result;
  }
}