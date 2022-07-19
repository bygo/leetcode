# Link: https://leetcode.cn/problems/report-contiguous-dates


SELECT `type` `period_state`, MIN(`date`) `start_date`, MAX(`date`) `end_date`
FROM (
         SELECT `type`, `date`, SUBDATE(`date`, ROW_NUMBER() OVER (PARTITION BY `type` ORDER BY `date`)) `diff`
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