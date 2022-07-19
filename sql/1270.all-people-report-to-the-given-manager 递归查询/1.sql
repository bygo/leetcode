# Link: https://leetcode.cn/problems/all-people-report-to-the-given-manager

SELECT `e1`.`employee_id`
FROM `employees`          `e1`
         JOIN `employees` `e2` ON `e1`.`manager_id` = `e2`.`employee_id`
         JOIN `employees` `e3` ON `e2`.`manager_id` = `e3`.`employee_id`
WHERE `e1`.`employee_id` != 1
  AND `e3`.`manager_id` = 1
