library(dplyr)

# tbl_df 함수를 이용하여 데이터 프레임 래핑
# => 데이터를 현재 화면에 맞는 사이즈로 줄여서 출력
iris_df <- tbl_df(iris)
iris_df

class(iris_df)

# 데이터프레임의 각 열에 대한 행의 값들을 출력
glimpse(iris_df)

# x %>% f(y) <-> x maps to f(x,y)k
# same as head(iris)
iris %>% head
head(iris)

# same as head(iris, 10)
iris %>% head(10)
head(iris, 10)

# filter(df, filter condition)
library(gapminder)
filter(gapminder, country=='Korea, Rep.')
filter(gapminder, year==2007)
gapminder %>% filter(country=='Korea, Rep.' & year==2007)

# arrange(df, column1, column2...) where column# : 정렬 기준
arrange(gapminder, year, country) %>% head(10)

# select(df, col1, col2...) where col# : 선택할 열
gapminder %>% select(pop, gdpPercap)

# mutate(df, f1(col...), f2(col...)...) where f# : 추가/변형시킬 열에 대한 함수
gapminder %>% mutate(
  total_gdp = pop * gdpPercap,
  le_gdp_ratio = lifeExp / gdpPercap,
  lgrk = le_gdp_ratio * 100
)


# summarize(df, f1(col..), f2(col..)) where f# : 벡터를 입력받아 하나의 통계량(스칼라) 반환 함수
gapminder %>% summarize(
  n_obs = n(),
  n_countries = n_distinct(country),
  n_yeers = n_distinct(year),
  median_gdp_per_cap = median(gdpPercap),
  max_gdp_per_cap = max(gdpPercap)
)

# sample_n(# of sample), sample_frac(ratio of sample): 랜덤 샘플링
# default replace=FALSE: 비복원 추출
sample_n(gapminder, 10)
gapminder %>% sample_frac(0.01)

# distint(df, col): col에 대한 고유한 행 값들 출력
distinct(gapminder, country)
gapminder %>% distinct(year)

# group_by: 데이터들의 그룹핑
# => 이후 통계량 및 랜덤 샘플링은 각 그룹에 대해서 연산
gapminder %>%
  filter(year == 2007) %>%
  group_by(continent) %>%
  summarize(
    n_countries = n_distinct(country),
    medain_life = median(lifeExp)
  )

# inner, left, right, outer join with dplyr
(df1 <- data_frame(x = c(1,2), y=2:1))
(df2 <- data_frame(x = c(1,3), a=10, b="a"))

df1 %>% inner_join(df2)
df1 %>% left_join(df2)
df1 %>% right_join(df2)
df1 %>% full_join(df2)