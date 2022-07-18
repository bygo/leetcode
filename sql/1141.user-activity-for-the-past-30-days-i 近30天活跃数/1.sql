# Link: https://leetcode.cn/problems/user-activity-for-the-past-30-days-i


SELECT IFNULL(
               ROUND(
                       COUNT(DISTINCT `session_id`) / COUNT(DISTINCT `user_id`),
                       2),
               0) `average_sessions_per_user`
FROM `activity`
WHERE '2019-06-27' < `activity_date`