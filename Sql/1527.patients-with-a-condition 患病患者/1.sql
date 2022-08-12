# Link: https://leetcode.cn/problems/patients-with-a-condition

# Write your MySQL query statement below

SELECT *
FROM `patients`
WHERE `conditions` LIKE 'DIAB1%'
   OR `conditions` LIKE '% DIAB1%'