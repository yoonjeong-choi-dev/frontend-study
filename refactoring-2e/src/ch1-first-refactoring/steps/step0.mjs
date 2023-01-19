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

    let currentAmount = 0;
    switch (play.type) {
      case 'tragedy':
        currentAmount = 40000;

        // bonus amount
        if (performance.audience > 30) {
          currentAmount += 1000 * (performance.audience - 30);
        }
        break;
      case  'comedy':
        currentAmount = 30000;

        // bonus amount
        currentAmount += 300 * performance.audience;
        if (performance.audience > 20) {
          currentAmount += 10000 + 500 * (performance.audience - 20);
        }
        break;
      default:
        throw new Error(`Unknown Type: ${play.type}`);
    }

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
}