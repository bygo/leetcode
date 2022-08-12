# Link: https://leetcode.cn/problems/queries-quality-and-percentage

SELECT `query_name`,
       ROUND(AVG(`rating` / `position`), 2)                   `quality`,
       ROUND(SUM(IF(`rating` < 3, 1, 0)) * 100 / COUNT(*), 2) `poor_query_percentage`
FROM `queries`
GROUP BY `query_name`