# Link: https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee


SELECT `E1`.`id`,
       `E1`.`month`,
       (IFNULL(`E1`.`salary`, 0) + IFNULL(`E2`.`salary`, 0) + IFNULL(`E3`.`salary`, 0)) AS `Salary`
FROM (SELECT `id`,
             MAX(`month`) AS `month`
      FROM `Employee`
      GROUP BY `id`
      HAVING COUNT(*) > 1) AS `maxmonth`
         LEFT JOIN
     `Employee` `E1` ON `maxmonth`.`id` = `E1`.`id`
         AND `maxmonth`.`month` > `E1`.`month`
         LEFT JOIN
     `Employee` `E2` ON `E2`.`id` = `E1`.`id`
         AND `E2`.`month` = `E1`.`month` - 1
         LEFT JOIN
     `Employee` `E3` ON `E3`.`id` = `E1`.`id`
         AND `E3`.`month` = `E1`.`month` - 2
ORDER BY `id`, `month` DESC;

#

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
