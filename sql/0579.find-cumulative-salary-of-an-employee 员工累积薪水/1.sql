# Link: https://leetcode.cn/problems/find-cumulative-salary-of-an-employee


SELECT `e1`.`id`,
       `e1`.`month`,
       (IFNULL(`e1`.`salary`, 0) + IFNULL(`e2`.`salary`, 0) + IFNULL(`e3`.`salary`, 0)) `salary`
FROM (SELECT `id`,
             MAX(`month`) `month`
      FROM `employee`
      GROUP BY `id`
      HAVING COUNT(*) > 1) AS `maxmonth`
         LEFT JOIN
     `employee`               `e1` ON `maxmonth`.`id` = `e1`.`id`
         AND `maxmonth`.`month` > `e1`.`month`
         LEFT JOIN
     `employee`               `e2` ON `e2`.`id` = `e1`.`id`
         AND `e2`.`month` = `e1`.`month` - 1
         LEFT JOIN
     `employee`               `e3` ON `e3`.`id` = `e1`.`id`
         AND `e3`.`month` = `e1`.`month` - 2
ORDER BY `id`, `month` DESC;

#

SELECT `id`, `month`, SUM(`salary`) `salary`
FROM (
         SELECT `e1`.`id`, `e1`.`month`, `e1`.`salary`
         FROM (SELECT `id`, `month`, `salary`, MAX(`month`) OVER (PARTITION BY `id`,`month`) `maxmonth`
               FROM `employee`)   `e1`
                  JOIN `employee` `e2`
         WHERE `e1`.`month` != `e1`.`maxmonth`
           AND `e1`.`id` = `e2`.`id`
           AND `e2`.`month` BETWEEN `e1`.`month` - 2 AND `e1`.`month`
         ORDER BY `id`
     ) `t`
GROUP BY `id`, `month`
ORDER BY `id`, `month` DESC;
