# Link: https://leetcode.cn/problems/top-travellers

SELECT `name`, IFNULL(SUM(`distance`), 0) `travelled_distance`
FROM `users`
         LEFT JOIN `rides` ON `users`.`id` = `rides`.`user_id`
GROUP BY `name`
ORDER BY `travelled_distance` DESC, `name`;
