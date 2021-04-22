# Link: https://leetcode-cn.com/problems/nth-highest-salary

CREATE FUNCTION getNthHighestSalary(`N` INT) RETURNS INT
BEGIN
    SET `N` := `N` - 1;
    RETURN (
        SELECT (
                   SELECT DISTINCT `Salary`
                   FROM `Employee`
                   ORDER BY `Salary` DESC
                   LIMIT `N`,1) `NthHighestSalary`
    );
END