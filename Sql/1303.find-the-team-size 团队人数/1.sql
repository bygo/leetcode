# Link: https://leetcode-cn.com/problems/find-the-team-size

SELECT `employee_id`, COUNT(*) OVER (PARTITION BY `team_id`) `team_size`
FROM `employee`
