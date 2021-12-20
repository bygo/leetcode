# Link: https://leetcode-cn.com/problems/median-employee-salary

SELECT `id`, `company`, `salary`
FROM (SELECT `id`,
             `company`,
             `salary`,
             ROW_NUMBER() OVER (PARTITION BY `company` ORDER BY `salary`) `rank`,
             COUNT(*) OVER (PARTITION BY `company`)                       `count`
      FROM `employee`) `t`
WHERE `rank` BETWEEN `count` / 2 AND `count` / 2 + 1