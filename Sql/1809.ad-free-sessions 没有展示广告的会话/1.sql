# Link: https://leetcode-cn.com/problems/ad-free-sessions

SELECT `session_id`
FROM `playback`
WHERE `session_id` NOT IN (
    SELECT DISTINCT `session_id`
    FROM `playback`     `p`
             JOIN `ads` `a` ON `p`.`customer_id` = `a`.`customer_id`
    WHERE `timestamp` BETWEEN `start_time` AND `end_time`)