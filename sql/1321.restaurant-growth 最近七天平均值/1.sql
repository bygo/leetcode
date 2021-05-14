# Link: https://leetcode-cn.com/problems/restaurant-growth

# Write your MySQL query statement below
SELECT `visited_on`,
       SUM(`amount`) OVER (ORDER BY `visited_on` ROWS 6 PRECEDING )               `amount`,
       ROUND(SUM(`amount`) OVER (ORDER BY `visited_on` ROWS 6 PRECEDING ) / 7, 2) `average_amount`
FROM (
         SELECT `visited_on`,
                SUM(`amount`) `amount`
         FROM `customer`
         GROUP BY `visited_on`) `a`
ORDER BY `visited_on`
LIMIT 6, 18446744073709551615

