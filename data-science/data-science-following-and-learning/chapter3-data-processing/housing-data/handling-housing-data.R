library(dplyr)

# boston: 14 columns, 506 data
boston <- read.table('tmp/housing.data')

# 데이터 컬럼들에 변수 할당
names(boston) <- c('crim', 'zn', 'indux', 'chas', 'nox', 'rm', 'age', 'dis', 'rad', 'tax', 'ptratio', 'black', 'lstat', 'mdev')

# 데이터 컬럼에 대한 각 로우의 데이터 출력
glimpse(boston)

# 산점도행렬(두 컬럼 간 관계) 그래프
plot(boston)

# 기본적인 통계량 출력
summary(boston)