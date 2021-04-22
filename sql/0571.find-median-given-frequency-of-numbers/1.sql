# Title: Find Median Given Frequency of Numbers
# Link: https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers


SELECT AVG(`Number`) AS `Median`
FROM (SELECT *,
             SUM(`Frequency`) OVER (ORDER BY `Number` ASC)  AS `n1`,
             SUM(`Frequency`) OVER (ORDER BY `Number` DESC) AS `n2`
      FROM `Numbers`) AS `t`
WHERE `n1` BETWEEN `n2` - `Frequency` AND `n2` + `Frequency`;
