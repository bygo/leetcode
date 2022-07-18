# Link: https://leetcode.cn/problems/second-highest-salary

SELECT (
           SELECT DISTINCT `salary`
           FROM `employee`
           ORDER BY `salary` DESC
           LIMIT 1,1) `secondhighestsalary`