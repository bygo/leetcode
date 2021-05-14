# Link: https://leetcode-cn.com/problems/biggest-window-between-visits


SELECT `user_id`, max(`diff`) `biggest_window`
FROM (
         SELECT `user_id`,
                datediff(
                        lead(`visit_date`, 1, '2021-01-01') OVER (PARTITION BY `user_id` ORDER BY `visit_date`),
                        `visit_date`
                    ) `diff`
         FROM `uservisits`
     ) `tmp`
GROUP BY `user_id`

