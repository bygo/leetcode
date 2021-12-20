# Link: https://leetcode-cn.com/problems/activity-participants


SELECT `activity`
FROM (SELECT `activity`,
             DENSE_RANK() OVER (ORDER BY COUNT(*))      `r1`,
             DENSE_RANK() OVER (ORDER BY COUNT(*) DESC) `r2`
      FROM `friends`
      GROUP BY `activity`) `t`
WHERE `r1` != 1
  AND `r2` != 1