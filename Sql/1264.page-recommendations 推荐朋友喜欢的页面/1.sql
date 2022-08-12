# Link: https://leetcode.cn/problems/page-recommendations

SELECT DISTINCT `page_id` `recommended_page`
FROM `likes`
WHERE `user_id` IN (
    SELECT (
               CASE
                   WHEN `user1_id` = 1 THEN `user2_id`
                   WHEN `user2_id` = 1 THEN `user1_id`
                   END
               ) `user_id`
    FROM `friendship`
    WHERE `user1_id` = 1
       OR `user2_id` = 1
)
  AND `page_id` NOT IN (
    SELECT `page_id`
    FROM `likes`
    WHERE `user_id` = 1
)