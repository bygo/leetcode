# Link: https://leetcode-cn.com/problems/replace-employee-id-with-the-unique-identifier

SELECT `unique_id`, `name`
FROM `employeeuni`
         RIGHT JOIN `employees` ON `employees`.`id` = `employeeuni`.`id`