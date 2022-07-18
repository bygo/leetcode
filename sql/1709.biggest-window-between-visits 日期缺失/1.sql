# Link: https://leetcode.cn/problems/biggest-window-between-visits


SELECT `user_id`, MAX(`diff`) `biggest_window`
FROM (
         SELECT `user_id`,
                DATEDIFF(
                        LEAD(`visit_date`, 1, '2021-01-01') OVER (PARTITION BY `user_id` ORDER BY `visit_date`),
                        `visit_date`
                    ) `diff`
         FROM `uservisits`
     ) `tmp`
GROUP BY `user_id`

