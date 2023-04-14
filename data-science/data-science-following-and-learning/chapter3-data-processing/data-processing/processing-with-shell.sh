#!/bin/bash

# download data
curl https://archive.ics.uci.edu/ml/machine-learning-databases/adult/adult.data > adult.data

echo "Display the first 10 rows"
head adult.data

echo ""
echo "Display the last 10 rows"
tail adult.data

echo ""
echo "Save the first 5 row to another data"
head -5 adult.data > adult.data.small
cat adult.data.small

echo ""
echo "Change comma to tab and save to another data"
tr "," "\t" < adult.data.small > adult.data.small.tab
cat adult.data.small.tab

echo ""
echo "Display the number of row:"
wc -l adult.data

echo ""
echo "Calculate the frequency distribution of the 'work class' column(second column)"
# cut -d ',' -f 2: 콤바 구분자를 기준으로 두 번째 열 추출
# sort | uniq -c | sort -nr: 알파벳 정렬 -> 각 행이 몇 번 등장했는지 카운팅 -> 횟수로 정렬
cut -d ',' -f 2 < adult.data | sort | uniq -c | sort -nr

# clean up
rm adult.data adult.data.small adult.data.small.tab