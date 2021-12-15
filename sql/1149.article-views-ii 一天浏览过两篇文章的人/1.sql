# Link: https://leetcode-cn.com/problems/article-views-ii

SELECT DISTINCT `viewer_id` `id`
FROM `views`
GROUP BY `viewer_id`, `view_date`
HAVING 1 < COUNT(DISTINCT `article_id`)
ORDER BY `id`