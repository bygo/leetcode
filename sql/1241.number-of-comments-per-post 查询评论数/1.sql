# Link: https://leetcode-cn.com/problems/number-of-comments-per-post

SELECT `post_id`, COUNT(`sub_id`) `number_of_comments`
FROM (
         SELECT DISTINCT `post`.`sub_id` `post_id`, `sub`.`sub_id` `sub_id`
         FROM `submissions`               `post`
                  LEFT JOIN `submissions` `sub`
                            ON `post`.`sub_id` = `sub`.`parent_id`
         WHERE `post`.`parent_id` IS NULL
     ) `t`
GROUP BY `post_id`
ORDER BY `post_id`
