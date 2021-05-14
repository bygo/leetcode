# Link: https://leetcode-cn.com/problems/new-users-daily-count


SELECT `activity_date` `login_date`, count(*) `user_count`
FROM (
         SELECT `user_id`, min(`activity_date`) `activity_date`
         FROM `traffic`
         WHERE `activity` = 'login'
         GROUP BY `user_id`) `t`
WHERE '2019-03-31' < `activity_date`
GROUP BY `login_date