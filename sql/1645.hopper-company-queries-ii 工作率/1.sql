# Link: https://leetcode-cn.com/problems/hopper-company-queries-ii

WITH RECURSIVE `t1`(`month`) AS (
    SELECT 1
    UNION ALL
    SELECT `month` + 1
    FROM `t1`
    WHERE `month` + 1 <= 12
)

SELECT `t1`.`month`,
       IFNULL(ROUND(IFNULL(`active_rides`, 0) / SUM(IFNULL(`c`, 0)) OVER (ORDER BY `t1`.`month`) * 100, 2),
              0) `working_percentage`
FROM `t1`
         LEFT JOIN
     (
         SELECT IF(YEAR(`join_date`) < '2020', 1, MONTH(`join_date`)) `month`, COUNT(DISTINCT `driver_id`) `c`
         FROM `drivers`
         WHERE YEAR(`join_date`) <= '2020'
         GROUP BY `month`
     ) `t2`
     ON `t1`.`month` = `t2`.`month`
         LEFT JOIN
     (
         SELECT MONTH(`r`.`requested_at`) `month`, COUNT(DISTINCT `driver_id`) `active_rides`
         FROM `acceptedrides`       `a`
                  LEFT JOIN `rides` `r` ON `a`.`ride_id` = `r`.`ride_id`
             AND YEAR(`r`.`requested_at`) >= '2020' AND YEAR(`r`.`requested_at`) < '2021'
         GROUP BY `month`
     ) `t3` ON `t1`.`month` = `t3`.`month`