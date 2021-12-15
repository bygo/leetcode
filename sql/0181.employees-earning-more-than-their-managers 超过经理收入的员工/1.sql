# Link: https://leetcode-cn.com/problems/employees-earning-more-than-their-managers

SELECT `e1`.`name` `employee`
FROM `employee`               AS `e1`
         LEFT JOIN `employee` AS `e2` ON `e1`.`managerid` = `e2`.`id`
WHERE `e2`.`salary` < `e1`.`salary`