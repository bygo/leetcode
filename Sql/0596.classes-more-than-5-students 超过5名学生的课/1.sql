# Link: https://leetcode.cn/problems/classes-more-than-5-students

SELECT `class`
FROM `courses`
GROUP BY `class`
HAVING COUNT(DISTINCT `student`) >= 5