# Link: https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers


SELECT AVG(`number`) `median`
FROM (SELECT *,
             SUM(`frequency`) OVER (ORDER BY `number` ASC)  `n1`,
             SUM(`frequency`) OVER (ORDER BY `number` DESC) `n2`
      FROM `numbers`) `t`
WHERE `n1` BETWEEN `n2` - `frequency` AND `n2` + `frequency`;
