# Link: https://leetcode.cn/problems/calculate-salaries


SELECT `s`.`company_id`,
       `s`.`employee_id`,
       `s`.`employee_name`,
       (
           CASE
               WHEN `maxsalary` < 1000 THEN `salary`
               WHEN `maxsalary` >= 1000 AND `maxsalary` < 10000 THEN ROUND(`salary` - `salary` * 0.24)
               WHEN `maxsalary` >= 10000 THEN ROUND(`salary` - `salary` * 0.49)
               END
           ) `salary`
FROM `salaries`      `s`
         LEFT JOIN (
                       SELECT `company_id`, MAX(`salary`) `maxsalary`
                       FROM `salaries`
                       GROUP BY `company_id`
                   ) `m`
                   ON `m`.`company_id` = `s`.`company_id`
