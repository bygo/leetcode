# Link: https://leetcode-cn.com/problems/reported-posts-ii

SELECT ROUND(AVG(`proportion`) * 100, 2) `average_daily_percent`
FROM (
         SELECT `a`.`action_date`,
                COUNT(DISTINCT `r`.`post_id`) / COUNT(DISTINCT `a`.`post_id`) `proportion`
         FROM `actions`                `a`
                  LEFT JOIN `removals` `r`
                            ON `a`.`post_id` = `r`.`post_id`
         WHERE `extra` = 'spam'
         GROUP BY `a`.`action_date`
     ) `a`;
