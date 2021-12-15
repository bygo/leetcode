# Link: https://leetcode-cn.com/problems/average-time-of-process-per-machine

SELECT `machine_id`,
       ROUND(SUM(IF(`activity_type` = 'end', `timestamp`, -`timestamp`)) / COUNT(*) * 2, 3) `processing_time`
FROM `activity`
GROUP BY `machine_id`