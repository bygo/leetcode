# Link: https://leetcode.cn/problems/find-total-time-spent-by-each-employee

SELECT `event_day` `day`, `emp_id`, SUM(`out_time` - `in_time`) `total_time`
FROM `employees`
GROUP BY `day`, `emp_id`