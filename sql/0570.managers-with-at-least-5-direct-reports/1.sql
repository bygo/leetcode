# Title: Managers with at Least 5 Direct Reports
# Link: https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports

SELECT `Name`
FROM `employee` `t1`
         JOIN (SELECT `ManagerId` FROM `employee` GROUP BY `ManagerId` HAVING COUNT(*) >= 5) `t2`
              ON `t1`.`Id` = `t2`.`ManagerId`