# Link: https://leetcode-cn.com/problems/patients-with-a-condition

# Write your MySQL query statement below

    SELECT *
    FROM `patients`
    WHERE `conditions` LIKE 'DIAB1%'
       OR `conditions` LIKE '% DIAB1%'