# Link: https://leetcode.cn/problems/students-with-invalid-departments

SELECT `id`, `name`
FROM `students`
WHERE `department_id` NOT IN (SELECT `id` FROM `departments`)