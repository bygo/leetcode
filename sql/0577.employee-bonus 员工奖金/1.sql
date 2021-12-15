# Link: https://leetcode-cn.com/problems/employee-bonus

SELECT `name`, `bonus`
FROM `employee`
         LEFT JOIN `bonus` ON `employee`.`empid` = `bonus`.`empid`
WHERE `bonus`.`bonus` < 1000
   OR `bonus` IS NULL