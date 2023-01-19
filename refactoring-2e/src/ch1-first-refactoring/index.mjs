import { plays, invoices } from './loadData.mjs';

import original from './steps/step0.mjs';

//import statement from './steps/step1.mjs';
//import statement from './steps/step2.mjs';
//import statement, { htmlStatement } from './steps/step3/step3.mjs';
import { htmlStatement as originalHTML } from './steps/step3/step3.mjs';
import statement, { htmlStatement } from './steps/step4/step4.mjs';

console.log('Chapter 1');

console.log('Original');
const originResult = original(invoices[0], plays);
console.log(originResult);

const originHTMLResult = originalHTML(invoices[0], plays);
console.log(originHTMLResult);

const refactoredResult = statement(invoices[0], plays);
console.log('Is Good for plain text? ', originResult === refactoredResult);
//console.log(refactoredResult);

const refactoredHtmlResult = htmlStatement(invoices[0], plays);
console.log('Is Good for HTML? ', originHTMLResult === refactoredHtmlResult);