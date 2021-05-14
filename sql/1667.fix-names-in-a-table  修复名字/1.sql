# Link: https://leetcode-cn.com/problems/fix-names-in-a-table

SELECT `user_id`,
       concat(upper(left(`name`, 1)), lower(substr(`name`, 2))) `name`
FROM `users`
ORDER BY `user_id`