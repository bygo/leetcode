# Link: https://leetcode-cn.com/problems/department-highest-salary

SELECT `d`.`name` 'Department',
       `e`.`name` 'Employee',
       `Salary`
FROM `Employee` `e`
         JOIN
     `Department` `d` ON `e`.`DepartmentId` = `d`.`Id`
         AND (`e`.`DepartmentId`, `Salary`) IN
             (SELECT `DepartmentId`,
                     MAX(`Salary`)
              FROM `Employee`
              GROUP BY `DepartmentId`
             );

#

SELECT `Department`,
       `Employee`,
       `Salary`
FROM (SELECT `d`.`name`                                                            'Department',
             `e`.`name`                                                            'Employee',
             `e`.`Salary`,
             RANK() OVER (PARTITION BY `e`.`DepartmentId` ORDER BY `Salary` DESC ) `r`
      FROM `Department` `d`
               JOIN `Employee` `e`
                    ON `d`.`id` = `e`.`DepartmentId`
     ) `t`
WHERE `r` = 1;