# Link: https://leetcode-cn.com/problems/find-the-start-and-end-number-of-continuous-ranges

SELECT MIN(`log_id`) `start_id`,
       MAX(`log_id`) `end_id`
FROM (
         SELECT DISTINCT `log_id`,
                         `log_id` - ROW_NUMBER() OVER ( ORDER BY `log_id` ASC ) `r`
         FROM `logs`
     ) `t`
GROUP BY `r`
ORDER BY `start_id`
