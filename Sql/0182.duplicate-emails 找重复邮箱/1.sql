# Link: https://leetcode-cn.com/problems/duplicate-emails

SELECT `email`
FROM `person`
GROUP BY `email`
HAVING 1 < COUNT(`email`)