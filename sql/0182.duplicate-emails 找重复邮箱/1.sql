# Link: https://leetcode-cn.com/problems/duplicate-emails

SELECT `Email`
FROM `Person`
GROUP BY `Email`
HAVING 1 < count(`Email`)