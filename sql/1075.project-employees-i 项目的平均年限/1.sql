# Title: Project Employees I
# Link: https://leetcode-cn.com/problems/project-employees-i

SELECT `project_id`, round(avg(`experience_years`), 2) AS `average_years`
FROM `Project` AS `p`
         INNER JOIN `Employee` AS `e`
                    ON `p`.`employee_id` = `e`.`employee_id`
GROUP BY `project_id`;