# Link: https://leetcode-cn.com/problems/activity-participants


SELECT `activity`
FROM (SELECT `activity`,
             dense_rank() OVER (ORDER BY count(*))      `r1`,
             dense_rank() OVER (ORDER BY count(*) DESC) `r2`
      FROM `friends`
      GROUP BY `activity`) `t`
WHERE `r1` != 1
  AND `r2` != 1