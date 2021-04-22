# Title: Department Top Three Salaries
# Link: https://leetcode-cn.com/problems/department-top-three-salaries

SELECT `Department`.`Name` AS `Department`,
       `e1`.`Name`         AS `Employee`,
       `e1`.`Salary`       AS `Salary`
FROM `Employee` AS `e1`
         JOIN `Department`
              ON
                      `e1`.`DepartmentId` = `Department`.`Id`
                      AND (SELECT count(DISTINCT `e2`.`Salary`)
                           FROM `Employee` AS `e2`
                           WHERE `e1`.`Salary` < `e2`.`Salary`
                             AND `e1`.`DepartmentId` = `e2`.`DepartmentId`) < 3
ORDER BY `Salary` DESC;
