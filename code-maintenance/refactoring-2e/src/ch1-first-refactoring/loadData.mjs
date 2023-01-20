import { readFile } from 'fs/promises';

export const plays = JSON.parse(await readFile('./data/plays.json', 'utf8'));
export const invoices = JSON.parse(await readFile('./data/invoices.json', 'utf8'));
