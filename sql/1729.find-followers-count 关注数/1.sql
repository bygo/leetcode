# Link: https://leetcode-cn.com/problems/find-followers-count

SELECT `user_id`, count(`follower_id`) `followers_count`
FROM `followers`
GROUP BY `user_id`
ORDER BY `user_id`

