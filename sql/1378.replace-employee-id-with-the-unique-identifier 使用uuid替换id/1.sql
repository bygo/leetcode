# Link: https://leetcode.cn/problems/replace-employee-id-with-the-unique-identifier

SELECT `unique_id`, `name`
FROM `employeeuni`
         RIGHT JOIN `employees` ON `employees`.`id` = `employeeuni`.`id`