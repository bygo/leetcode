# Link: https://leetcode-cn.com/problems/hopper-company-queries-i

WITH RECURSIVE `t1`(`month`) AS (
    SELECT 1
    UNION ALL
    SELECT `month` + 1
    FROM `t1`
    WHERE `month` + 1 <= 12
)

SELECT `t1`.`month`,
       sum(ifnull(`c`, 0)) OVER (ORDER BY `t1`.`month`) `active_drivers`,
       ifnull(`active_rides`, 0)                        `accepted_rides`
FROM `t1`
         LEFT JOIN
     (
         SELECT if(year(`join_date`) < '2020', 1, month(`join_date`)) `month`, count(`driver_id`) `c`
         FROM `drivers`
         WHERE year(`join_date`) <= '2020'
         GROUP BY `month`
     ) `t2` ON `t1`.`month` = `t2`.`month`
         LEFT JOIN
     (
         SELECT month(`r`.`requested_at`) `month`, count(`driver_id`) `active_rides`
         FROM `acceptedrides`       `a`
                  LEFT JOIN `rides` `r` ON `a`.`ride_id` = `r`.`ride_id`
             AND year(`r`.`requested_at`) >= '2020' AND year(`r`.`requested_at`) < '2021'
         GROUP BY `month`
     ) `t3` ON `t1`.`month` = `t3`.`month`