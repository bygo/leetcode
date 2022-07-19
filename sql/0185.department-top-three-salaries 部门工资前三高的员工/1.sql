# Link: https://leetcode.cn/problems/department-top-three-salaries

SELECT `department`, `employee`, `salary`
FROM (SELECT `d`.`name`                                                                       `department`,
             `e1`.`name`                                                                      `employee`,
             `e1`.`salary`                                                                    `salary`,
             DENSE_RANK() OVER (PARTITION BY `e1`.`departmentid` ORDER BY `e1`.`salary` DESC) `r`
      FROM `employee`            `e1`
               JOIN `department` `d`
                    ON
                        `e1`.`departmentid` = `d`.`id`) `t`
WHERE `r` <= 3;

#

SELECT `d`.`name`    `department`,
       `e1`.`name`   `employee`,
       `e1`.`salary` `salary`
FROM `employee`            `e1`
         JOIN `department` `d`
              ON
                      `e1`.`departmentid` = `d`.`id`
                      AND (SELECT COUNT(DISTINCT `e2`.`salary`)
                           FROM `employee` AS `e2`
                           WHERE `e1`.`salary` < `e2`.`salary`
                             AND `e1`.`departmentid` = `e2`.`departmentid`) < 3
ORDER BY `salary` DESC;