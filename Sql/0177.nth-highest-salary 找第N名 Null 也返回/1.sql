# Link: https://leetcode.cn/problems/nth-highest-salary

CREATE FUNCTION getNthHighestSalary(`n` INT) RETURNS INT
BEGIN
    SET `n` := `n` - 1;
    RETURN (
        SELECT (
                   SELECT DISTINCT `salary`
                   FROM `employee`
                   ORDER BY `salary` DESC
                   LIMIT `n`,1) `nthhighestsalary`
    );
END