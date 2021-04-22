# Title: Find Cumulative Salary of an Employee
# Link: https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee


SELECT `id`, `Month`, SUM(`Salary`) AS `Salary`
FROM (
         SELECT `e1`.`id`, `e1`.`month`, `e1`.`salary`
         FROM (SELECT `id`, `month`, `salary`, MAX(`month`) OVER (PARTITION BY `id`,`month`) AS `maxmonth`
               FROM `Employee`) `e1`
                  JOIN `Employee` `e2`
         WHERE `e1`.`month` != `e1`.`maxmonth`
           AND `e1`.`id` = `e2`.`id`
           AND `e2`.`month` BETWEEN `e1`.`month` - 2 AND `e1`.`month`
         ORDER BY `id`
     ) `t`
GROUP BY `id`, `month`
ORDER BY `id`, `month` DESC;
