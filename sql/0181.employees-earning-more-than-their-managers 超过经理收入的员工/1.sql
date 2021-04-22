# Link: https://leetcode-cn.com/problems/employees-earning-more-than-their-managers

SELECT `E1`.`Name` AS `Employee`
FROM `Employee` AS `E1`
         LEFT JOIN `Employee` AS `E2` ON `E1`.`ManagerId` = `E2`.`id`
WHERE `E2`.`Salary` < `E1`.`Salary`