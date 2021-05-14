# Link: https://leetcode-cn.com/problems/department-top-three-salaries

SELECT `Department`, `Employee`, `Salary`
FROM (SELECT `d`.`Name`                                                                       `Department`,
             `e1`.`Name`                                                                      `Employee`,
             `e1`.`Salary`                                                                    `Salary`,
             dense_rank() OVER (PARTITION BY `e1`.`DepartmentId` ORDER BY `e1`.`Salary` DESC) `r`
      FROM `Employee` `e1`
               JOIN `Department` `d`
                    ON
                        `e1`.`DepartmentId` = `d`.`Id`) `t`
WHERE `r` <= 3;

#

SELECT `d`.`Name`    `Department`,
       `e1`.`Name`   `Employee`,
       `e1`.`Salary` `Salary`
FROM `Employee` `e1`
         JOIN `Department` `d`
              ON
                      `e1`.`DepartmentId` = `d`.`Id`
                      AND (SELECT count(DISTINCT `e2`.`Salary`)
                           FROM `Employee` AS `e2`
                           WHERE `e1`.`Salary` < `e2`.`Salary`
                             AND `e1`.`DepartmentId` = `e2`.`DepartmentId`) < 3
ORDER BY `Salary` DESC;