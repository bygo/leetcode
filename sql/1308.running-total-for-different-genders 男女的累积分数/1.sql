# Link: https://leetcode-cn.com/problems/running-total-for-different-genders

SELECT `gender`,
       `day`,
       SUM(`score_points`) OVER (PARTITION BY `gender` ORDER BY `day`) `total`
FROM `scores`;