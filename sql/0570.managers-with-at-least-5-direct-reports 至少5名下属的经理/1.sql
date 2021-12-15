# Link: https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports

SELECT `name`
FROM `employee`                                                                              `t1`
         JOIN (SELECT `managerid` FROM `employee` GROUP BY `managerid` HAVING COUNT(*) >= 5) `t2`
              ON `t1`.`id` = `t2`.`managerid`;

#

SELECT `name`
FROM `employee`
WHERE `id` IN (SELECT DISTINCT `managerid`
               FROM (SELECT `managerid`,
                            COUNT(`managerid`) OVER (PARTITION BY `managerid`) `c`
                     FROM `employee`
                     ORDER BY `c` DESC
                    ) `t`
               WHERE `c` >= 5)