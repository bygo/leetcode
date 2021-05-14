# Link: https://leetcode-cn.com/problems/median-employee-salary

SELECT `id`, `Company`, `Salary`
FROM (SELECT `id`,
             `Company`,
             `Salary`,
             ROW_NUMBER() OVER (PARTITION BY `Company` ORDER BY `Salary`) AS `rank`,
             COUNT(*) OVER (PARTITION BY `Company`)                       AS `count`
      FROM `Employee`) `t`
WHERE `RANK` BETWEEN `count` / 2 AND `count` / 2 + 1