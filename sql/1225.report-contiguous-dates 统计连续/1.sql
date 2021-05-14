# Link: https://leetcode-cn.com/problems/report-contiguous-dates


SELECT `type` `period_state`, min(`date`) `start_date`, max(`date`) `end_date`
FROM (
         SELECT `type`, `date`, subdate(`date`, row_number() OVER (PARTITION BY `type` ORDER BY `date`)) `diff`
         FROM (
                  SELECT 'failed' `type`, `fail_date` `date`
                  FROM `failed`
                  UNION ALL
                  SELECT 'succeeded' `type`, `success_date` `date`
                  FROM `succeeded`
              ) `t1`
     ) `t2`
WHERE `date` BETWEEN '2019-01-01' AND '2019-12-31'
GROUP BY `type`, `diff`
ORDER BY `start_date`