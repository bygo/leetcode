# Link: https://leetcode-cn.com/problems/reported-posts

SELECT `extra` "report_reason", count(DISTINCT `post_id`) "report_count"
FROM `Actions`
WHERE `action_date` = '2019-07-04'
  AND `action` = 'report'
  AND `extra` IS NOT NULL
GROUP BY `extra`