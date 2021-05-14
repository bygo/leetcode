# Link: https://leetcode-cn.com/problems/find-total-time-spent-by-each-employee

SELECT `event_day` `day`, `emp_id`, sum(`out_time` - `in_time`) `total_time`
FROM `employees`
GROUP BY `day`, `emp_id`