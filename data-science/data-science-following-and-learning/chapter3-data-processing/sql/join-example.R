library(dplyr)
library(sqldf)

# Create two tables
(df1 <- data_frame(x = c(1,2), y=2:1))
(df2 <- data_frame(x = c(1,3), a=10, b="a"))

# inner join
sqldf(
  "select * from df1
    inner join df2
    on df1.x = df2.x
  "
)

# left join
sqldf(
  "select * from df1
    left join df2
    on df1.x = df2.x
  "
)