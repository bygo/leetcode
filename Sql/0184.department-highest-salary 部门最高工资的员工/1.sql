# Link: https://leetcode.cn/problems/department-highest-salary

SELECT `d`.`name` 'department',
       `e`.`name` 'employee',
       `salary`
FROM `employee`   `e`
         JOIN
     `department` `d` ON `e`.`departmentid` = `d`.`id`
         AND (`e`.`departmentid`, `salary`) IN
             (SELECT `departmentid`,
                     MAX(`salary`)
              FROM `employee`
              GROUP BY `departmentid`
             );

#

SELECT `department`,
       `employee`,
       `salary`
FROM (SELECT `d`.`name`                                                            'department',
             `e`.`name`                                                            'employee',
             `e`.`salary`,
             RANK() OVER (PARTITION BY `e`.`departmentid` ORDER BY `salary` DESC ) `r`
      FROM `department`        `d`
               JOIN `employee` `e`
                    ON `d`.`id` = `e`.`departmentid`
     ) `t`
WHERE `r` = 1;