# Link: https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers


SELECT AVG(`Number`) `Median`
FROM (SELECT *,
             SUM(`Frequency`) OVER (ORDER BY `Number` ASC)  `n1`,
             SUM(`Frequency`) OVER (ORDER BY `Number` DESC) `n2`
      FROM `Numbers`) `t`
WHERE `n1` BETWEEN `n2` - `Frequency` AND `n2` + `Frequency`;
