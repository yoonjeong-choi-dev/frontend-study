import createStatementData from './createStatementData.mjs';

export default function (invoice, plays) {
  return renderPlainText(createStatementData(invoice, plays));
}

// 첫번째 단계에서 처리한 데이터를 이용하여 문자열 출력
// => invoice 및 plays 매개 변수 불필요
function renderPlainText(data) {
  let result = `청구 내역(customer: ${data.customer})\n`;
  for (let performance of data.performances) {
    // Update result
    result += `\t${performance.play.name}: ${usd(performance.amount)} (${performance.audience} seats)\n`
  }

  result += `Total Amount: ${usd(data.totalAmount)}\n`;
  result += `Reward Points: ${data.totalVolumeCredits} points\n`;
  return result;
}


// New feature: render HTML
export function htmlStatement(invoice, plays) {
  return renderHTML(createStatementData(invoice, plays));
}

function renderHTML(data) {
  let result = `<h1>청구 내역(customer: ${data.customer})</h1>`

  result += '<table>\n'
  result += '<tr><th>Performance</th><th>Seats</th><th>Price</th></tr>'

  for (let performance of data.performances) {
    result +=
      ` <tr>
          <td>${performance.play.name}</td>
          <td>${performance.audience}</td>
          <td>${performance.amount}</td>
        </tr>
      `;
  }
  result += '</table>\n'

  result += `<p>Total Amount: <em>${usd(data.totalAmount)}</em></p>\n`;
  result += `<p>Reward Points: <em>${data.totalVolumeCredits}</em> points</p>\n`;
  return result;
}


function usd(number) {
  return new Intl.NumberFormat('es-US', {
    style: 'currency',
    currency: 'USD',
    minimumFractionDigits: 2,
  }).format(number / 100);
}