# Title: Project Employees I
# Link: https://leetcode-cn.com/problems/project-employees-i

SELECT `project_id`, ROUND(AVG(`experience_years`), 2) `average_years`
FROM `project`                 AS `p`
         INNER JOIN `employee` AS `e`
                    ON `p`.`employee_id` = `e`.`employee_id`
GROUP BY `project_id`;