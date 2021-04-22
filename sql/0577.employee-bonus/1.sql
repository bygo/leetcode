# Title: Employee Bonus
# Link: https://leetcode-cn.com/problems/employee-bonus

SELECT `name`, `bonus`
FROM `Employee`
         LEFT JOIN `Bonus` ON `Employee`.`empid` = `Bonus`.`empid`
WHERE `bonus`.`bonus` < 1000
   OR `bonus` IS NULL