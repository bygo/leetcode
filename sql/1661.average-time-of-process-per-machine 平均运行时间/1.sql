# Link: https://leetcode-cn.com/problems/average-time-of-process-per-machine

SELECT `machine_id`,
       round(sum(IF(`activity_type` = 'end', `timestamp`, -`timestamp`)) / count(*) * 2, 3) `processing_time`
FROM `activity`
GROUP BY `machine_id`