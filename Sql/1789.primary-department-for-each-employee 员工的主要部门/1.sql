# Link: https://leetcode-cn.com/problems/primary-department-for-each-employee

SELECT DISTINCT `employee_id`,
                FIRST_VALUE(`department_id`) OVER (PARTITION BY `employee_id` ORDER BY `primary_flag`) `department_id`
FROM `employee`