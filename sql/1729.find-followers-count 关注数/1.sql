# Link: https://leetcode.cn/problems/find-followers-count

SELECT `user_id`, COUNT(`follower_id`) `followers_count`
FROM `followers`
GROUP BY `user_id`
ORDER BY `user_id`

