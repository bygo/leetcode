# Link: https://leetcode-cn.com/problems/project-employees-iii

SELECT `project_id`, `employee_id`
FROM (SELECT `project_id`,
             `p`.`employee_id`,
             RANK() OVER (PARTITION BY `project_id` ORDER BY `experience_years` DESC) AS `r`
      FROM `project` `p`
               INNER JOIN `employee` `e`
                          ON `P`.`employee_id` = `e`.`employee_id`) `a`
WHERE `r` = 1;