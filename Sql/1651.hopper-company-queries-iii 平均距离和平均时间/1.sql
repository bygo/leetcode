# Link: https://leetcode-cn.com/problems/hopper-company-queries-iii

WITH RECURSIVE `t1`(`month`) AS (
    SELECT 1
    UNION ALL
    SELECT `month` + 1
    FROM `t1`
    WHERE `month` + 1 <= 12
)

SELECT `month`, `average_ride_distance`, `average_ride_duration`
FROM (SELECT `month`,
             ROUND(AVG(IFNULL(`ride_distance`, 0)) OVER (ORDER BY `month` ROWS BETWEEN CURRENT ROW AND 2 FOLLOWING ),
                   2)                             `average_ride_distance`,
             ROUND(AVG(IFNULL(`ride_duration`, 0)) OVER (ORDER BY `month` ROWS BETWEEN CURRENT ROW AND 2 FOLLOWING ),
                   2)                             `average_ride_duration`,
             ROW_NUMBER() OVER (ORDER BY `month`) `r`
      FROM `t1`
               LEFT JOIN
           (SELECT DISTINCT MONTH(`requested_at`)                                                        `month`,
                            SUM(`ride_distance`)
                                OVER (PARTITION BY MONTH(`requested_at`) ORDER BY MONTH(`requested_at`)) `ride_distance`,
                            SUM(`ride_duration`)
                                OVER (PARTITION BY MONTH(`requested_at`) ORDER BY MONTH(`requested_at`)) `ride_duration`
            FROM `acceptedrides`
                     JOIN `rides` USING (`ride_id`)
            WHERE YEAR(`requested_at`) = 2020) `a`
           USING (`month`)) `b`
WHERE `b`.`r` <= 10
