# Title: Project Employees II
# Link: https://leetcode-cn.com/problems/project-employees-ii

WITH `tmp` AS (SELECT `project_id`, count(*) AS `c` FROM `project` GROUP BY `project_id`)
SELECT `project_id`
FROM `tmp`
WHERE `c` = (SELECT max(`c`) FROM `tmp`);