library(dplyr)
library(gapminder)

# Exercise 1. 갭마인더 데이터
# Exericse 1-a. 2007년 1인당 GDP
gapminder %>%
  filter(year == 2007) %>%
  select(country, gdpPercap)

# Exericse 1-b. 2007년 대륙별 평균 수명의 평균 및 중앙값
gapminder %>%
  filter(year == 2007) %>%
  group_by(continent) %>%
  summarize(
    n_size = n(),
    mean_life = mean(lifeExp),
    medain_life = median(lifeExp)
  ) %>%
  select(continent, n_size, mean_life, medain_life)


# Exercise 2. Load an external file
# https://www.kaggle.com/datasets/arnabchaki/data-science-salaries-2023
salaries <- read.csv('tmp/ds_salaries.csv')

# Exercise 3. Summarize the data from Exercise 2
salaries %>%
  filter(employment_type == "FT" & salary_currency == "USD") %>%
  group_by(experience_level) %>%
  summarize(
    n_size = n(),
    mean_salary = mean(salary),
    medain_salary = median(salary)
  ) %>%
  select(experience_level, mean_salary, medain_salary, n_size)


# Exercise 3. IDMB Data
# Exercise 3-a. Load Data
idmb <- read.csv('tmp/imdb_top_1000.csv')
glimpse(idmb)

# Exercise 3-b. 연도별 리뷰받은 영화의 개수
idmb %>%
  group_by(Released_Year) %>%
  summarize(num_movies = n()) %>%
  select(Released_Year, num_movies)

# Exercise 3-b. 연도별 리뷰 평균
idmb %>%
  group_by(Released_Year) %>%
  summarize(
    mean_rating = mean(IMDB_Rating),
    median_rating = median(IMDB_Rating)
  ) %>%
  select(Released_Year, mean_rating, median_rating)