# Link: https://leetcode-cn.com/problems/percentage-of-users-attended-a-contest


SELECT `contest_id`,
       round(count(DISTINCT `r`.`user_id`) / count(DISTINCT `u`.`user_id`) * 100, 2) `percentage`
FROM `register`       `r`
         JOIN `users` `u`
GROUP BY `contest_id`
ORDER BY `percentage` DESC, `contest_id`