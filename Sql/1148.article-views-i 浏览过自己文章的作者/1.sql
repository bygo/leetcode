# Link: https://leetcode-cn.com/problems/article-views-i

SELECT DISTINCT `author_id` `id`
FROM `views`
WHERE `viewer_id` = `author_id`
ORDER BY `author_id`