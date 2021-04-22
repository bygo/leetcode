# Title: Classes More Than 5 Students
# Link: https://leetcode-cn.com/problems/classes-more-than-5-students

SELECT `class`
FROM `courses`
GROUP BY `class`
HAVING COUNT(DISTINCT `student`) >= 5