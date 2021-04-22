# Title: Consecutive Available Seats
# Link: https://leetcode-cn.com/problems/consecutive-available-seats

SELECT `seat_id`
FROM (SELECT `seat_id`, COUNT(*) OVER (PARTITION BY `r`) AS `c`
      FROM (SELECT `seat_id`,
                   `seat_id` - ROW_NUMBER() OVER (ORDER BY `seat_id`) AS `r`
            FROM `cinema`
            WHERE `free` = 1) `t`) `t2`
WHERE `c` > 1
GROUP BY `seat_id`
